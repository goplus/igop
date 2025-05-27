// export by github.com/goplus/ixgo/cmd/qexp

package xml

import (
	q "github.com/goplus/xgo/tpl/encoding/xml"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "xml",
		Path: "github.com/goplus/xgo/tpl/encoding/xml",
		Deps: map[string]string{
			"encoding/xml": "xml",
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
