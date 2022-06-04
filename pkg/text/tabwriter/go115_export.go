// export by github.com/goplus/igop/cmd/qexp

//+build go1.15,!go1.16

package tabwriter

import (
	q "text/tabwriter"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "tabwriter",
		Path: "text/tabwriter",
		Deps: map[string]string{
			"io":           "io",
			"unicode/utf8": "utf8",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]igop.NamedType{
			"Writer": {reflect.TypeOf((*q.Writer)(nil)).Elem(), "", "Flush,Init,Write,addLine,append,dump,endEscape,flush,flushNoDefers,format,handlePanic,reset,startEscape,terminateCell,updateWidth,write0,writeLines,writeN,writePadding"},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"NewWriter": reflect.ValueOf(q.NewWriter),
		},
		TypedConsts: map[string]igop.TypedConst{
			"AlignRight":          {reflect.TypeOf(q.AlignRight), constant.MakeInt64(int64(q.AlignRight))},
			"Debug":               {reflect.TypeOf(q.Debug), constant.MakeInt64(int64(q.Debug))},
			"DiscardEmptyColumns": {reflect.TypeOf(q.DiscardEmptyColumns), constant.MakeInt64(int64(q.DiscardEmptyColumns))},
			"FilterHTML":          {reflect.TypeOf(q.FilterHTML), constant.MakeInt64(int64(q.FilterHTML))},
			"StripEscape":         {reflect.TypeOf(q.StripEscape), constant.MakeInt64(int64(q.StripEscape))},
			"TabIndent":           {reflect.TypeOf(q.TabIndent), constant.MakeInt64(int64(q.TabIndent))},
		},
		UntypedConsts: map[string]igop.UntypedConst{
			"Escape": {"untyped rune", constant.MakeInt64(int64(q.Escape))},
		},
	})
}
