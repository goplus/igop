// export by github.com/goplus/gossa/cmd/qexp

//go:build go1.17
// +build go1.17

package regexp

import (
	q "regexp"

	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "regexp",
		Path: "regexp",
		Deps: map[string]string{
			"bytes":         "bytes",
			"io":            "io",
			"regexp/syntax": "syntax",
			"sort":          "sort",
			"strconv":       "strconv",
			"strings":       "strings",
			"sync":          "sync",
			"unicode":       "unicode",
			"unicode/utf8":  "utf8",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]gossa.NamedType{
			"Regexp": {reflect.TypeOf((*q.Regexp)(nil)).Elem(), "", "Copy,Expand,ExpandString,Find,FindAll,FindAllIndex,FindAllString,FindAllStringIndex,FindAllStringSubmatch,FindAllStringSubmatchIndex,FindAllSubmatch,FindAllSubmatchIndex,FindIndex,FindReaderIndex,FindReaderSubmatchIndex,FindString,FindStringIndex,FindStringSubmatch,FindStringSubmatchIndex,FindSubmatch,FindSubmatchIndex,LiteralPrefix,Longest,Match,MatchReader,MatchString,NumSubexp,ReplaceAll,ReplaceAllFunc,ReplaceAllLiteral,ReplaceAllLiteralString,ReplaceAllString,ReplaceAllStringFunc,Split,String,SubexpIndex,SubexpNames,allMatches,backtrack,doExecute,doMatch,doOnePass,expand,get,pad,put,replaceAll,tryBacktrack"},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Compile":          reflect.ValueOf(q.Compile),
			"CompilePOSIX":     reflect.ValueOf(q.CompilePOSIX),
			"Match":            reflect.ValueOf(q.Match),
			"MatchReader":      reflect.ValueOf(q.MatchReader),
			"MatchString":      reflect.ValueOf(q.MatchString),
			"MustCompile":      reflect.ValueOf(q.MustCompile),
			"MustCompilePOSIX": reflect.ValueOf(q.MustCompilePOSIX),
			"QuoteMeta":        reflect.ValueOf(q.QuoteMeta),
		},
		TypedConsts:   map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}
