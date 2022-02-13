// export by github.com/goplus/gossa/cmd/qexp

//+build go1.16,!go1.17

package palette

import (
	q "image/color/palette"

	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "palette",
		Path: "image/color/palette",
		Deps: map[string]string{
			"image/color": "color",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]gossa.NamedType{},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"Plan9":   reflect.ValueOf(&q.Plan9),
			"WebSafe": reflect.ValueOf(&q.WebSafe),
		},
		Funcs:         map[string]reflect.Value{},
		TypedConsts:   map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}
