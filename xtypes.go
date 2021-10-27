package gossa

import (
	"go/types"
	"log"
	"reflect"
	"unsafe"

	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/types/typeutil"
)

var (
	record = NewTypesRecord()
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

type TypesRecord struct {
	Tcache *typeutil.Map
}

func NewTypesRecord() *TypesRecord {
	return &TypesRecord{
		Tcache: &typeutil.Map{},
	}
}

func (r *TypesRecord) ToType(typ types.Type) reflect.Type {
	if rt := r.Tcache.At(typ); rt != nil {
		return rt.(reflect.Type)
	}
	if rt := rtyp.Tcache.At(typ); rt != nil {
		r.Tcache.Set(typ, rt)
		return rt.(reflect.Type)
	}
	log.Printf("-----> %v: %T", typ, typ)
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
		//return toStructType(t, ctx)
		panic("impl")
	case *types.Named:
		panic("impl")
		//return toNamedType(t, ctx)
	case *types.Interface:
		panic("impl")
		//return toInterfaceType(t, ctx)
	case *types.Signature:
		in := r.ToTypeList(t.Params())
		out := r.ToTypeList(t.Results())
		rt = reflect.FuncOf(in, out, t.Variadic())
	default:
		panic("unreachable")
	}
	r.Tcache.Set(typ, rt)
	return rt
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

func (r *TypesRecord) LoadType(typ types.Type) {
	rt := r.ToType(typ)
	log.Println("=======>", typ, rt)
}

func (r *TypesRecord) Load(pkg *ssa.Package) error {
	checked := make(map[types.Type]bool)
	for _, v := range pkg.Members {
		typ := v.Type()
		if checked[typ] {
			continue
		}
		checked[typ] = true
		// log.Println(k, v, v.Type())
		switch t := v.(type) {
		case *ssa.NamedConst:
			// if typ := t.Type().String(); strings.HasPrefix(typ, "untyped ") {
			// 	e.UntypedConsts = append(e.UntypedConsts, fmt.Sprintf("%q: gossa.UntypedConst{%q, %v}", pkgPath+"."+t.Name(), t.Type().String(), p.constToLit(t.Value.Value)))
			// } else {
			// 	e.TypedConsts = append(e.TypedConsts, fmt.Sprintf("%q : gossa.TypedConst{reflect.TypeOf(%v), %v}", pkgPath+"."+t.Name(), pkgName+"."+t.Name(), p.constToLit(t.Value.Value)))
			// }
		case *ssa.Global:
			// e.Vars = append(e.Vars, fmt.Sprintf("%q : reflect.ValueOf(&%v)", pkgPath+"."+t.Name(), pkgName+"."+t.Name()))
			r.LoadType(typ)
		case *ssa.Function:
			r.LoadType(typ)
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

	log.Println(r.Tcache.KeysString())

	return nil
}
