// export by github.com/goplus/igop/cmd/qexp

//+build go1.16,!go1.17

package plan9obj

import (
	q "debug/plan9obj"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
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
		NamedTypes: map[string]reflect.Type{
			"File":          reflect.TypeOf((*q.File)(nil)).Elem(),
			"FileHeader":    reflect.TypeOf((*q.FileHeader)(nil)).Elem(),
			"Section":       reflect.TypeOf((*q.Section)(nil)).Elem(),
			"SectionHeader": reflect.TypeOf((*q.SectionHeader)(nil)).Elem(),
			"Sym":           reflect.TypeOf((*q.Sym)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"NewFile": reflect.ValueOf(q.NewFile),
			"Open":    reflect.ValueOf(q.Open),
		},
		TypedConsts: map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{
			"Magic386":   {"untyped int", constant.MakeInt64(int64(q.Magic386))},
			"Magic64":    {"untyped int", constant.MakeInt64(int64(q.Magic64))},
			"MagicAMD64": {"untyped int", constant.MakeInt64(int64(q.MagicAMD64))},
			"MagicARM":   {"untyped int", constant.MakeInt64(int64(q.MagicARM))},
		},
	})
}
