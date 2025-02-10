// export by github.com/goplus/igop/cmd/qexp

package reflect2

import (
	q "github.com/modern-go/reflect2"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "reflect2",
		Path: "github.com/modern-go/reflect2",
		Deps: map[string]string{
			"reflect": "reflect",
			"runtime": "runtime",
			"sync":    "sync",
			"unsafe":  "unsafe",
		},
		Interfaces: map[string]reflect.Type{
			"API":           reflect.TypeOf((*q.API)(nil)).Elem(),
			"ArrayType":     reflect.TypeOf((*q.ArrayType)(nil)).Elem(),
			"InterfaceType": reflect.TypeOf((*q.InterfaceType)(nil)).Elem(),
			"ListType":      reflect.TypeOf((*q.ListType)(nil)).Elem(),
			"MapIterator":   reflect.TypeOf((*q.MapIterator)(nil)).Elem(),
			"MapType":       reflect.TypeOf((*q.MapType)(nil)).Elem(),
			"PtrType":       reflect.TypeOf((*q.PtrType)(nil)).Elem(),
			"SliceType":     reflect.TypeOf((*q.SliceType)(nil)).Elem(),
			"StructField":   reflect.TypeOf((*q.StructField)(nil)).Elem(),
			"StructType":    reflect.TypeOf((*q.StructType)(nil)).Elem(),
			"Type":          reflect.TypeOf((*q.Type)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"Config":            reflect.TypeOf((*q.Config)(nil)).Elem(),
			"UnsafeArrayType":   reflect.TypeOf((*q.UnsafeArrayType)(nil)).Elem(),
			"UnsafeEFaceType":   reflect.TypeOf((*q.UnsafeEFaceType)(nil)).Elem(),
			"UnsafeIFaceType":   reflect.TypeOf((*q.UnsafeIFaceType)(nil)).Elem(),
			"UnsafeMapIterator": reflect.TypeOf((*q.UnsafeMapIterator)(nil)).Elem(),
			"UnsafeMapType":     reflect.TypeOf((*q.UnsafeMapType)(nil)).Elem(),
			"UnsafePtrType":     reflect.TypeOf((*q.UnsafePtrType)(nil)).Elem(),
			"UnsafeSliceType":   reflect.TypeOf((*q.UnsafeSliceType)(nil)).Elem(),
			"UnsafeStructField": reflect.TypeOf((*q.UnsafeStructField)(nil)).Elem(),
			"UnsafeStructType":  reflect.TypeOf((*q.UnsafeStructType)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"ConfigSafe":   reflect.ValueOf(&q.ConfigSafe),
			"ConfigUnsafe": reflect.ValueOf(&q.ConfigUnsafe),
		},
		Funcs: map[string]reflect.Value{
			"DefaultTypeOfKind": reflect.ValueOf(q.DefaultTypeOfKind),
			"IFaceToEFace":      reflect.ValueOf(q.IFaceToEFace),
			"IsNil":             reflect.ValueOf(q.IsNil),
			"IsNullable":        reflect.ValueOf(q.IsNullable),
			"NoEscape":          reflect.ValueOf(q.NoEscape),
			"PtrOf":             reflect.ValueOf(q.PtrOf),
			"PtrTo":             reflect.ValueOf(q.PtrTo),
			"RTypeOf":           reflect.ValueOf(q.RTypeOf),
			"Type2":             reflect.ValueOf(q.Type2),
			"TypeByName":        reflect.ValueOf(q.TypeByName),
			"TypeByPackageName": reflect.ValueOf(q.TypeByPackageName),
			"TypeOf":            reflect.ValueOf(q.TypeOf),
			"TypeOfPtr":         reflect.ValueOf(q.TypeOfPtr),
			"UnsafeCastString":  reflect.ValueOf(q.UnsafeCastString),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
