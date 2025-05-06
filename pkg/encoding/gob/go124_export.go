// export by github.com/goplus/igop/cmd/qexp

//go:build go1.24
// +build go1.24

package gob

import (
	q "encoding/gob"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "gob",
		Path: "encoding/gob",
		Deps: map[string]string{
			"bufio":            "bufio",
			"encoding":         "encoding",
			"encoding/binary":  "binary",
			"errors":           "errors",
			"fmt":              "fmt",
			"internal/saferio": "saferio",
			"io":               "io",
			"maps":             "maps",
			"math":             "math",
			"math/bits":        "bits",
			"os":               "os",
			"reflect":          "reflect",
			"sync":             "sync",
			"sync/atomic":      "atomic",
			"unicode":          "unicode",
			"unicode/utf8":     "utf8",
		},
		Interfaces: map[string]reflect.Type{
			"GobDecoder": reflect.TypeOf((*q.GobDecoder)(nil)).Elem(),
			"GobEncoder": reflect.TypeOf((*q.GobEncoder)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"CommonType": reflect.TypeOf((*q.CommonType)(nil)).Elem(),
			"Decoder":    reflect.TypeOf((*q.Decoder)(nil)).Elem(),
			"Encoder":    reflect.TypeOf((*q.Encoder)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"NewDecoder":   reflect.ValueOf(q.NewDecoder),
			"NewEncoder":   reflect.ValueOf(q.NewEncoder),
			"Register":     reflect.ValueOf(q.Register),
			"RegisterName": reflect.ValueOf(q.RegisterName),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
