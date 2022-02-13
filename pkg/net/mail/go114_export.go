// export by github.com/goplus/gossa/cmd/qexp

//+build go1.14,!go1.15

package mail

import (
	q "net/mail"

	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "mail",
		Path: "net/mail",
		Deps: map[string]string{
			"bufio":         "bufio",
			"errors":        "errors",
			"fmt":           "fmt",
			"io":            "io",
			"log":           "log",
			"mime":          "mime",
			"net/textproto": "textproto",
			"strings":       "strings",
			"sync":          "sync",
			"time":          "time",
			"unicode/utf8":  "utf8",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]gossa.NamedType{
			"Address":       {reflect.TypeOf((*q.Address)(nil)).Elem(), "", "String"},
			"AddressParser": {reflect.TypeOf((*q.AddressParser)(nil)).Elem(), "", "Parse,ParseList"},
			"Header":        {reflect.TypeOf((*q.Header)(nil)).Elem(), "AddressList,Date,Get", ""},
			"Message":       {reflect.TypeOf((*q.Message)(nil)).Elem(), "", ""},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"ErrHeaderNotPresent": reflect.ValueOf(&q.ErrHeaderNotPresent),
		},
		Funcs: map[string]reflect.Value{
			"ParseAddress":     reflect.ValueOf(q.ParseAddress),
			"ParseAddressList": reflect.ValueOf(q.ParseAddressList),
			"ParseDate":        reflect.ValueOf(q.ParseDate),
			"ReadMessage":      reflect.ValueOf(q.ReadMessage),
		},
		TypedConsts:   map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}
