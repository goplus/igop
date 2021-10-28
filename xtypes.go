package gossa

import (
	"fmt"
	"go/token"
	"go/types"
	"reflect"
	"unsafe"

	"github.com/goplus/reflectx"
	"golang.org/x/tools/go/ssa"
)

var (
	tyEmptyStruct = reflect.TypeOf((*struct{})(nil)).Elem()
)

var basicTypes = [...]reflect.Type{
	types.Bool:          reflect.TypeOf(false),
	types.Int:           reflect.TypeOf(0),
	types.Int8:          reflect.TypeOf(int8(0)),
	types.Int16:         reflect.TypeOf(int16(0)),
	types.Int32:         reflect.TypeOf(int32(0)),
	types.Int64:         reflect.TypeOf(int64(0)),
	types.Uint:          reflect.TypeOf(uint(0)),
	types.Uint8:         reflect.TypeOf(uint8(0)),
	types.Uint16:        reflect.TypeOf(uint16(0)),
	types.Uint32:        reflect.TypeOf(uint32(0)),
	types.Uint64:        reflect.TypeOf(uint64(0)),
	types.Uintptr:       reflect.TypeOf(uintptr(0)),
	types.Float32:       reflect.TypeOf(float32(0)),
	types.Float64:       reflect.TypeOf(float64(0)),
	types.Complex64:     reflect.TypeOf(complex64(0)),
	types.Complex128:    reflect.TypeOf(complex128(0)),
	types.String:        reflect.TypeOf(""),
	types.UnsafePointer: reflect.TypeOf(unsafe.Pointer(nil)),

	types.UntypedBool:    reflect.TypeOf(false),
	types.UntypedInt:     reflect.TypeOf(0),
	types.UntypedRune:    reflect.TypeOf('a'),
	types.UntypedFloat:   reflect.TypeOf(0.1),
	types.UntypedComplex: reflect.TypeOf(0 + 1i),
	types.UntypedString:  reflect.TypeOf(""),
}

type FindMethod interface {
	FindMethod(mtyp reflect.Type, fn *types.Func) func([]reflect.Value) []reflect.Value
}

type TypesRecord struct {
	find FindMethod
}

func NewTypesRecord(find FindMethod) *TypesRecord {
	return &TypesRecord{
		find: find,
	}
}

func (r *TypesRecord) ToType(typ types.Type) reflect.Type {
	if rt := rtyp.Tcache.At(typ); rt != nil {
		return rt.(reflect.Type)
	}
	var rt reflect.Type
	switch t := typ.(type) {
	case *types.Basic:
		kind := t.Kind()
		if kind > types.Invalid && kind < types.UntypedNil {
			rt = basicTypes[kind]
		}
	case *types.Pointer:
		elem := r.ToType(t.Elem())
		rt = reflect.PtrTo(elem)
	case *types.Slice:
		elem := r.ToType(t.Elem())
		rt = reflect.SliceOf(elem)
	case *types.Array:
		elem := r.ToType(t.Elem())
		rt = reflect.ArrayOf(int(t.Len()), elem)
	case *types.Map:
		key := r.ToType(t.Key())
		elem := r.ToType(t.Elem())
		rt = reflect.MapOf(key, elem)
	case *types.Chan:
		elem := r.ToType(t.Elem())
		rt = reflect.ChanOf(toReflectChanDir(t.Dir()), elem)
	case *types.Struct:
		rt = r.toStructType(t)
	case *types.Named:
		rt = r.toNamedType(t)
	case *types.Interface:
		rt = r.toInterfaceType(t)
	case *types.Signature:
		in := r.ToTypeList(t.Params())
		out := r.ToTypeList(t.Results())
		rt = reflect.FuncOf(in, out, t.Variadic())
	default:
		panic("unreachable")
	}
	rtyp.Tcache.Set(typ, rt)
	rtyp.Rcache[rt] = typ
	return rt
}

func (r *TypesRecord) toInterfaceType(t *types.Interface) reflect.Type {
	n := t.NumMethods()
	if n == 0 {
		return tyEmptyInterface
	}
	ms := make([]reflect.Method, n, n)
	for i := 0; i < n; i++ {
		fn := t.Method(i)
		mtyp := r.ToType(fn.Type())
		ms[i] = reflect.Method{
			Name: fn.Name(),
			Type: mtyp,
		}
		if pkg := fn.Pkg(); pkg != nil {
			ms[i].PkgPath = pkg.Path()
		}
	}
	return reflectx.InterfaceOf(nil, ms)
}

func (r *TypesRecord) toNamedType(t *types.Named) reflect.Type {
	name := t.Obj()
	if name.Pkg() == nil {
		if name.Name() == "error" {
			return tyErrorInterface
		}
		return r.ToType(t.Underlying())
	}
	utype := r.ToType(t.Underlying())
	typ := reflectx.NamedTypeOf(name.Pkg().Path(), name.Name(), utype)
	if typ.Kind() != reflect.Interface {
		typ = r.toMethodSet(t, typ)
	}
	return typ
}

func (r *TypesRecord) toStructType(t *types.Struct) reflect.Type {
	n := t.NumFields()
	if n == 0 {
		return tyEmptyStruct
	}
	flds := make([]reflect.StructField, n)
	for i := 0; i < n; i++ {
		flds[i] = r.toStructField(t.Field(i), t.Tag(i))
	}
	typ := reflectx.StructOf(flds)
	return typ
}

func (r *TypesRecord) toStructField(v *types.Var, tag string) reflect.StructField {
	name := v.Name()
	typ := r.ToType(v.Type())
	fld := reflect.StructField{
		Name:      name,
		Type:      typ,
		Tag:       reflect.StructTag(tag),
		Anonymous: v.Anonymous(),
	}
	if !token.IsExported(name) {
		fld.PkgPath = v.Pkg().Path()
	}
	return fld
}

func (r *TypesRecord) ToTypeList(tuple *types.Tuple) []reflect.Type {
	n := tuple.Len()
	if n == 0 {
		return nil
	}
	list := make([]reflect.Type, n, n)
	for i := 0; i < n; i++ {
		list[i] = r.ToType(tuple.At(i).Type())
	}
	return list
}

func isPointer(typ types.Type) bool {
	_, ok := typ.Underlying().(*types.Pointer)
	return ok
}

func (r *TypesRecord) toMethodSet1(t *types.Named, styp reflect.Type) reflect.Type {
	return styp
}

func (r *TypesRecord) toMethodSet(t types.Type, styp reflect.Type) reflect.Type {
	methods := IntuitiveMethodSet(t)
	numMethods := len(methods)
	if numMethods == 0 {
		return styp
	}
	var mcount, pcount int
	for i := 0; i < numMethods; i++ {
		sig := methods[i].Type().(*types.Signature)
		pointer := isPointer(sig.Recv().Type())
		if !pointer {
			mcount++
		}
		pcount++
	}
	typ := reflectx.NewMethodSet(styp, mcount, pcount)
	fn := func() error {
		var ms []reflectx.Method
		for i := 0; i < numMethods; i++ {
			fn := methods[i].Obj().(*types.Func)
			sig := methods[i].Type().(*types.Signature)
			pointer := isPointer(sig.Recv().Type())
			mtyp := r.ToType(sig)
			var mfn func(args []reflect.Value) []reflect.Value
			idx := methods[i].Index()
			if len(idx) > 1 {
				isptr := isPointer(fn.Type().Underlying().(*types.Signature).Recv().Type())
				mfn = func(args []reflect.Value) []reflect.Value {
					this := args[0].FieldByIndex(idx[:len(idx)-1])
					if isptr && this.Kind() != reflect.Ptr {
						this = this.Addr()
					}
					m := this.MethodByName(fn.Name())
					return m.Call(args[1:])
				}
			} else {
				mfn = r.find.FindMethod(mtyp, fn)
			}
			var pkgpath string
			if pkg := fn.Pkg(); pkg != nil {
				pkgpath = pkg.Path()
			}
			ms = append(ms, reflectx.MakeMethod(fn.Name(), pkgpath, pointer, mtyp, mfn))
		}
		return reflectx.SetMethodSet(typ, ms, false)
	}
	err := fn()
	if err != nil {
		panic(fmt.Errorf("SetMethodSet failed: %v", typ))
	}
	return typ
}

func toReflectChanDir(d types.ChanDir) reflect.ChanDir {
	switch d {
	case types.SendRecv:
		return reflect.BothDir
	case types.SendOnly:
		return reflect.SendDir
	case types.RecvOnly:
		return reflect.RecvDir
	}
	return 0
}

func (r *TypesRecord) LoadType(typ types.Type) reflect.Type {
	return r.ToType(typ)
}

func (r *TypesRecord) Load(pkg *ssa.Package) error {
	checked := make(map[types.Type]bool)
	for _, v := range pkg.Members {
		typ := v.Type()
		if checked[typ] {
			continue
		}
		checked[typ] = true
		r.LoadType(typ)
		switch t := v.(type) {
		case *ssa.NamedConst:
			// if typ := t.Type().String(); strings.HasPrefix(typ, "untyped ") {
			// 	e.UntypedConsts = append(e.UntypedConsts, fmt.Sprintf("%q: gossa.UntypedConst{%q, %v}", pkgPath+"."+t.Name(), t.Type().String(), p.constToLit(t.Value.Value)))
			// } else {
			// 	e.TypedConsts = append(e.TypedConsts, fmt.Sprintf("%q : gossa.TypedConst{reflect.TypeOf(%v), %v}", pkgPath+"."+t.Name(), pkgName+"."+t.Name(), p.constToLit(t.Value.Value)))
			// }
		case *ssa.Global:
			// e.Vars = append(e.Vars, fmt.Sprintf("%q : reflect.ValueOf(&%v)", pkgPath+"."+t.Name(), pkgName+"."+t.Name()))
		case *ssa.Function:
			// e.Funcs = append(e.Funcs, fmt.Sprintf("%q : reflect.ValueOf(%v)", pkgPath+"."+t.Name(), pkgName+"."+t.Name()))
		case *ssa.Type:
			// typ := t.Type()
			// if obj, ok := t.Object().(*types.TypeName); ok && obj.IsAlias() {
			// 	name := obj.Name()
			// 	switch typ := obj.Type().(type) {
			// 	case *types.Named:
			// 		e.AliasTypes = append(e.AliasTypes, fmt.Sprintf("%q: reflect.TypeOf((*%v.%v)(nil)).Elem()", pkgPath+"."+name, sname, name))
			// 	case *types.Basic:
			// 		e.AliasTypes = append(e.AliasTypes, fmt.Sprintf("%q: reflect.TypeOf((*%v)(nil)).Elem()", pkgPath+"."+name, typ.Name()))
			// 	default:
			// 		log.Panicln("error parser", typ)
			// 	}
			// 	continue
			// }
			// var typeName string
			// switch t := typ.(type) {
			// case *types.Named:
			// 	typeName = t.Obj().Name()
			// 	e.Types = append(e.Types, fmt.Sprintf("reflect.TypeOf((*%v.%v)(nil))", pkgName, typeName))
			// 	if t.Obj().Pkg() != pkg.Pkg {
			// 		continue
			// 	}
			// case *types.Basic:
			// 	typeName = t.Name()
			// 	e.Types = append(e.Types, fmt.Sprintf("reflect.TypeOf((*%v.%v)(nil))", pkgName, typeName))
			// }
			// methods := IntuitiveMethodSet(typ)
			// for _, method := range methods {
			// 	name := method.Obj().Name()
			// 	if token.IsExported(name) {
			// 		info := fmt.Sprintf("(%v).%v", method.Recv(), name)
			// 		if pkgPath == pkgName {
			// 			e.Methods = append(e.Methods, fmt.Sprintf("%q : reflect.ValueOf(%v)", info, info))
			// 		} else {
			// 			var fn string
			// 			if isPointer(method.Recv()) {
			// 				fn = fmt.Sprintf("(*%v.%v).%v", pkgName, typeName, name)
			// 			} else {
			// 				fn = fmt.Sprintf("(%v.%v).%v", pkgName, typeName, name)
			// 			}
			// 			e.Methods = append(e.Methods, fmt.Sprintf("%q: reflect.ValueOf(%v)", info, fn))
			// 		}
			// 	}
			// }
		default:
			_ = t
			panic("unreachable")
		}
	}
	return nil
}

// golang.org/x/tools/go/types/typeutil.IntuitiveMethodSet
func IntuitiveMethodSet(T types.Type) []*types.Selection {
	isPointerToConcrete := func(T types.Type) bool {
		ptr, ok := T.(*types.Pointer)
		return ok && !types.IsInterface(ptr.Elem())
	}

	var result []*types.Selection
	mset := types.NewMethodSet(T)
	if types.IsInterface(T) || isPointerToConcrete(T) {
		for i, n := 0, mset.Len(); i < n; i++ {
			result = append(result, mset.At(i))
		}
	} else {
		// T is some other concrete type.
		// Report methods of T and *T, preferring those of T.
		pmset := types.NewMethodSet(types.NewPointer(T))
		for i, n := 0, pmset.Len(); i < n; i++ {
			meth := pmset.At(i)
			if m := mset.Lookup(meth.Obj().Pkg(), meth.Obj().Name()); m != nil {
				meth = m
			}
			result = append(result, meth)
		}
	}
	return result
}
