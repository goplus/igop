// export by github.com/goplus/igop/cmd/qexp

package regexposix

import (
	q "github.com/goplus/gop/tpl/encoding/regexposix"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "regexposix",
		Path: "github.com/goplus/gop/tpl/encoding/regexposix",
		Deps: map[string]string{
			"regexp": "regexp",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"New": reflect.ValueOf(q.New),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
