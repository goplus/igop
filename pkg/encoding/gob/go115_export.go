// export by github.com/goplus/ixgo/cmd/qexp

//+build go1.15,!go1.16

package gob

import (
	q "encoding/gob"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "gob",
		Path: "encoding/gob",
		Deps: map[string]string{
			"bufio":           "bufio",
			"encoding":        "encoding",
			"encoding/binary": "binary",
			"errors":          "errors",
			"fmt":             "fmt",
			"io":              "io",
			"math":            "math",
			"math/bits":       "bits",
			"os":              "os",
			"reflect":         "reflect",
			"sync":            "sync",
			"sync/atomic":     "atomic",
			"unicode":         "unicode",
			"unicode/utf8":    "utf8",
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
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
