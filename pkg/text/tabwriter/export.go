// export by github.com/goplus/gossa/cmd/qexp

package tabwriter

import (
	q "text/tabwriter"

	"go/constant"
	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "tabwriter",
		Path: "text/tabwriter",
		Deps: map[string]string{
			"io":           "io",
			"unicode/utf8": "utf8",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]gossa.NamedType{
			"Writer": {reflect.TypeOf((*q.Writer)(nil)).Elem(), "", "Flush,Init,Write,addLine,append,dump,endEscape,flush,flushNoDefers,format,handlePanic,reset,startEscape,terminateCell,updateWidth,write0,writeLines,writeN,writePadding"},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"NewWriter": reflect.ValueOf(q.NewWriter),
		},
		TypedConsts: map[string]gossa.TypedConst{
			"AlignRight":          {reflect.TypeOf(q.AlignRight), constant.MakeInt64(4)},
			"Debug":               {reflect.TypeOf(q.Debug), constant.MakeInt64(32)},
			"DiscardEmptyColumns": {reflect.TypeOf(q.DiscardEmptyColumns), constant.MakeInt64(8)},
			"FilterHTML":          {reflect.TypeOf(q.FilterHTML), constant.MakeInt64(1)},
			"StripEscape":         {reflect.TypeOf(q.StripEscape), constant.MakeInt64(2)},
			"TabIndent":           {reflect.TypeOf(q.TabIndent), constant.MakeInt64(16)},
		},
		UntypedConsts: map[string]gossa.UntypedConst{
			"Escape": {"untyped rune", constant.MakeInt64(255)},
		},
	})
}
