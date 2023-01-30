//go:build go1.18
// +build go1.18

package igop

import (
	"go/ast"
	"go/types"
	"reflect"
	"strconv"
	"strings"

	"github.com/goplus/reflectx"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/types/typeutil"
)

const (
	enabledTypeParam = true
)

func hasTypeParam(t types.Type) bool {
	switch t := t.(type) {
	case *types.TypeParam:
		return true
	case *types.Named:
		return t.TypeParams() != nil
	case *types.Signature:
		return t.TypeParams() != nil
	}
	return false
}

type nestedStack struct {
	targs []string
	cache []*typeutil.Map
}

func (s *nestedStack) Push(targs string) *typeutil.Map {
	s.targs = append(s.targs, targs)
	m := &typeutil.Map{}
	s.cache = append(s.cache, m)
	return m
}

func (s *nestedStack) Pop() (targs string, cache *typeutil.Map) {
	n := len(s.targs)
	if n >= 1 {
		targs = s.targs[n-1]
		cache = s.cache[n-1]
	}
	s.targs = s.targs[:n-1]
	s.cache = s.cache[:n-1]
	return
}

func (r *TypesRecord) typeId(typ types.Type, t reflect.Type) string {
	path := t.PkgPath()
	if path == "" {
		return t.Name()
	}
	name := path + "." + t.Name()
	if named, ok := typ.(*types.Named); ok {
		if n := r.nested[named.Origin()]; n != 0 {
			name += "·" + strconv.Itoa(n)
		}
	}
	return name
}

func (r *TypesRecord) EnterInstance(fn *ssa.Function) {
	r.fntargs = ""
	r.fntargs = r.parseFuncTypeArgs(fn)
	r.ncache = r.nstack.Push(r.fntargs)
}

func (r *TypesRecord) LeaveInstance(fn *ssa.Function) {
	r.fntargs, r.ncache = r.nstack.Pop()
}

func (r *TypesRecord) parseFuncTypeArgs(fn *ssa.Function) (targs string) {
	name := fn.Name()
	pos := strings.Index(name, "[")
	if pos < 0 {
		return ""
	}
	v := reflectx.FieldByNameX(reflect.ValueOf(fn).Elem(), "typeargs")
	var args []string
	for i := 0; i < v.Len(); i++ {
		typ := v.Index(i).Interface().(types.Type)
		rt, _ := r.ToType(typ)
		args = append(args, r.typeId(typ, rt))
	}
	return strings.Join(args, ",")
}

func (r *TypesRecord) extractNamed(named *types.Named, totype bool) (pkgpath string, name string, typeargs bool, nested bool) {
	obj := named.Obj()
	if pkg := obj.Pkg(); pkg != nil {
		if pkg.Name() == "main" {
			pkgpath = "main"
		} else {
			pkgpath = pkg.Path()
		}
	}
	if r.fntargs != "" && r.nested[named.Origin()] != 0 {
		nested = true
	}
	name = obj.Name()
	var ids string = r.fntargs
	if args := named.TypeArgs(); args != nil {
		typeargs = true
		var targs []string
		for i := 0; i < args.Len(); i++ {
			if totype {
				typ := args.At(i)
				rt, _ := r.ToType(typ)
				targs = append(targs, r.typeId(typ, rt))
			} else {
				targs = append(targs, args.At(i).String())
			}
		}
		if ids != "" {
			ids += ";"
		}
		ids += strings.Join(targs, ",")
	}
	if ids != "" {
		name += "[" + ids + "]"
	}
	return
}

func (r *TypesRecord) LookupReflect(typ types.Type) (rt reflect.Type, ok bool, nested bool) {
	rt, ok = r.loader.LookupReflect(typ)
	if !ok {
		n := len(r.nstack.cache)
		for i := n; i > 0; i-- {
			if rt := r.nstack.cache[i-1].At(typ); rt != nil {
				return rt.(reflect.Type), true, true
			}
		}
		if rt := r.tcache.At(typ); rt != nil {
			return rt.(reflect.Type), true, false
		}
	}
	return
}

func (sp *sourcePackage) Load() (err error) {
	if sp.Info == nil {
		sp.Info = &types.Info{
			Types:      make(map[ast.Expr]types.TypeAndValue),
			Defs:       make(map[*ast.Ident]types.Object),
			Uses:       make(map[*ast.Ident]types.Object),
			Implicits:  make(map[ast.Node]types.Object),
			Scopes:     make(map[ast.Node]*types.Scope),
			Selections: make(map[*ast.SelectorExpr]*types.Selection),
			Instances:  make(map[*ast.Ident]types.Instance),
		}
		if err := types.NewChecker(sp.Context.conf, sp.Context.FileSet, sp.Package, sp.Info).Files(sp.Files); err != nil {
			return err
		}
	}
	return
}
