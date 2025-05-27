// export by github.com/goplus/ixgo/cmd/qexp

package stringutil

import (
	q "github.com/qiniu/x/stringutil"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "stringutil",
		Path: "github.com/qiniu/x/stringutil",
		Deps: map[string]string{
			"unicode":      "unicode",
			"unicode/utf8": "utf8",
			"unsafe":       "unsafe",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Builder": reflect.TypeOf((*q.Builder)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Capitalize":     reflect.ValueOf(q.Capitalize),
			"Concat":         reflect.ValueOf(q.Concat),
			"Diff":           reflect.ValueOf(q.Diff),
			"NewBuilder":     reflect.ValueOf(q.NewBuilder),
			"NewBuilderSize": reflect.ValueOf(q.NewBuilderSize),
			"String":         reflect.ValueOf(q.String),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
