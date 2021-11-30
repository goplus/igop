// export by github.com/goplus/gossa/cmd/qexp

package plan9obj

import (
	q "debug/plan9obj"

	"go/constant"
	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "plan9obj",
		Path: "debug/plan9obj",
		Deps: map[string]string{
			"encoding/binary": "binary",
			"errors":          "errors",
			"fmt":             "fmt",
			"io":              "io",
			"os":              "os",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]gossa.NamedType{
			"File":          {reflect.TypeOf((*q.File)(nil)).Elem(), "", "Close,Section,Symbols"},
			"FileHeader":    {reflect.TypeOf((*q.FileHeader)(nil)).Elem(), "", ""},
			"Section":       {reflect.TypeOf((*q.Section)(nil)).Elem(), "", "Data,Open"},
			"SectionHeader": {reflect.TypeOf((*q.SectionHeader)(nil)).Elem(), "", ""},
			"Sym":           {reflect.TypeOf((*q.Sym)(nil)).Elem(), "", ""},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"NewFile": reflect.ValueOf(q.NewFile),
			"Open":    reflect.ValueOf(q.Open),
		},
		TypedConsts: map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{
			"Magic386":   {"untyped int", constant.MakeInt64(491)},
			"Magic64":    {"untyped int", constant.MakeInt64(32768)},
			"MagicAMD64": {"untyped int", constant.MakeInt64(35479)},
			"MagicARM":   {"untyped int", constant.MakeInt64(1607)},
		},
	})
}
