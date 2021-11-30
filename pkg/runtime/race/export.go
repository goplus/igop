// export by github.com/goplus/gossa/cmd/qexp

package race

import (
	q "runtime/race"

	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name:          "race",
		Path:          "runtime/race",
		Deps:          map[string]string{},
		Interfaces:    map[string]reflect.Type{},
		NamedTypes:    map[string]gossa.NamedType{},
		AliasTypes:    map[string]reflect.Type{},
		Vars:          map[string]reflect.Value{},
		Funcs:         map[string]reflect.Value{},
		TypedConsts:   map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}
