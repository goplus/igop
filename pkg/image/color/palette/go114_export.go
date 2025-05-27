// export by github.com/goplus/ixgo/cmd/qexp

//+build go1.14,!go1.15

package palette

import (
	q "image/color/palette"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "palette",
		Path: "image/color/palette",
		Deps: map[string]string{
			"image/color": "color",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"Plan9":   reflect.ValueOf(&q.Plan9),
			"WebSafe": reflect.ValueOf(&q.WebSafe),
		},
		Funcs:         map[string]reflect.Value{},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
