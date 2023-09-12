// export by github.com/goplus/igop/cmd/qexp

//go:build go1.20 && !go1.20
// +build go1.20,!go1.20

package strings

import (
	q "strings"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "strings",
		Path: "strings",
		Deps: map[string]string{
			"errors":           "errors",
			"internal/bytealg": "bytealg",
			"io":               "io",
			"sync":             "sync",
			"unicode":          "unicode",
			"unicode/utf8":     "utf8",
			"unsafe":           "unsafe",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Builder":  reflect.TypeOf((*q.Builder)(nil)).Elem(),
			"Reader":   reflect.TypeOf((*q.Reader)(nil)).Elem(),
			"Replacer": reflect.TypeOf((*q.Replacer)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Clone":          reflect.ValueOf(q.Clone),
			"Compare":        reflect.ValueOf(q.Compare),
			"Contains":       reflect.ValueOf(q.Contains),
			"ContainsAny":    reflect.ValueOf(q.ContainsAny),
			"ContainsRune":   reflect.ValueOf(q.ContainsRune),
			"Count":          reflect.ValueOf(q.Count),
			"Cut":            reflect.ValueOf(q.Cut),
			"CutPrefix":      reflect.ValueOf(q.CutPrefix),
			"CutSuffix":      reflect.ValueOf(q.CutSuffix),
			"EqualFold":      reflect.ValueOf(q.EqualFold),
			"Fields":         reflect.ValueOf(q.Fields),
			"FieldsFunc":     reflect.ValueOf(q.FieldsFunc),
			"HasPrefix":      reflect.ValueOf(q.HasPrefix),
			"HasSuffix":      reflect.ValueOf(q.HasSuffix),
			"Index":          reflect.ValueOf(q.Index),
			"IndexAny":       reflect.ValueOf(q.IndexAny),
			"IndexByte":      reflect.ValueOf(q.IndexByte),
			"IndexFunc":      reflect.ValueOf(q.IndexFunc),
			"IndexRune":      reflect.ValueOf(q.IndexRune),
			"Join":           reflect.ValueOf(q.Join),
			"LastIndex":      reflect.ValueOf(q.LastIndex),
			"LastIndexAny":   reflect.ValueOf(q.LastIndexAny),
			"LastIndexByte":  reflect.ValueOf(q.LastIndexByte),
			"LastIndexFunc":  reflect.ValueOf(q.LastIndexFunc),
			"Map":            reflect.ValueOf(q.Map),
			"NewReader":      reflect.ValueOf(q.NewReader),
			"NewReplacer":    reflect.ValueOf(q.NewReplacer),
			"Repeat":         reflect.ValueOf(q.Repeat),
			"Replace":        reflect.ValueOf(q.Replace),
			"ReplaceAll":     reflect.ValueOf(q.ReplaceAll),
			"Split":          reflect.ValueOf(q.Split),
			"SplitAfter":     reflect.ValueOf(q.SplitAfter),
			"SplitAfterN":    reflect.ValueOf(q.SplitAfterN),
			"SplitN":         reflect.ValueOf(q.SplitN),
			"Title":          reflect.ValueOf(q.Title),
			"ToLower":        reflect.ValueOf(q.ToLower),
			"ToLowerSpecial": reflect.ValueOf(q.ToLowerSpecial),
			"ToTitle":        reflect.ValueOf(q.ToTitle),
			"ToTitleSpecial": reflect.ValueOf(q.ToTitleSpecial),
			"ToUpper":        reflect.ValueOf(q.ToUpper),
			"ToUpperSpecial": reflect.ValueOf(q.ToUpperSpecial),
			"ToValidUTF8":    reflect.ValueOf(q.ToValidUTF8),
			"Trim":           reflect.ValueOf(q.Trim),
			"TrimFunc":       reflect.ValueOf(q.TrimFunc),
			"TrimLeft":       reflect.ValueOf(q.TrimLeft),
			"TrimLeftFunc":   reflect.ValueOf(q.TrimLeftFunc),
			"TrimPrefix":     reflect.ValueOf(q.TrimPrefix),
			"TrimRight":      reflect.ValueOf(q.TrimRight),
			"TrimRightFunc":  reflect.ValueOf(q.TrimRightFunc),
			"TrimSpace":      reflect.ValueOf(q.TrimSpace),
			"TrimSuffix":     reflect.ValueOf(q.TrimSuffix),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
