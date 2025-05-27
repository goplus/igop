// export by github.com/goplus/ixgo/cmd/qexp

package csv

import (
	q "github.com/goplus/xgo/tpl/encoding/csv"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "csv",
		Path: "github.com/goplus/xgo/tpl/encoding/csv",
		Deps: map[string]string{
			"encoding/csv": "csv",
			"strings":      "strings",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"New": reflect.ValueOf(q.New),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
