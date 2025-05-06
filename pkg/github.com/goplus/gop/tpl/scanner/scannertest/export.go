// export by github.com/goplus/igop/cmd/qexp

package scannertest

import (
	q "github.com/goplus/gop/tpl/scanner/scannertest"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "scannertest",
		Path: "github.com/goplus/gop/tpl/scanner/scannertest",
		Deps: map[string]string{
			"fmt":                               "fmt",
			"github.com/goplus/gop/scanner":     "scanner",
			"github.com/goplus/gop/token":       "token",
			"github.com/goplus/gop/tpl/scanner": "scanner",
			"github.com/goplus/gop/tpl/token":   "token",
			"go/scanner":                        "scanner",
			"go/token":                          "token",
			"io":                                "io",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"GoScan":  reflect.ValueOf(q.GoScan),
			"GopScan": reflect.ValueOf(q.GopScan),
			"Scan":    reflect.ValueOf(q.Scan),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
