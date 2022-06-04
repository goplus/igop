// export by github.com/goplus/igop/cmd/qexp

//+build go1.16,!go1.17

package url

import (
	q "net/url"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "url",
		Path: "net/url",
		Deps: map[string]string{
			"errors":  "errors",
			"fmt":     "fmt",
			"sort":    "sort",
			"strconv": "strconv",
			"strings": "strings",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]igop.NamedType{
			"Error":            {reflect.TypeOf((*q.Error)(nil)).Elem(), "", "Error,Temporary,Timeout,Unwrap"},
			"EscapeError":      {reflect.TypeOf((*q.EscapeError)(nil)).Elem(), "Error", ""},
			"InvalidHostError": {reflect.TypeOf((*q.InvalidHostError)(nil)).Elem(), "Error", ""},
			"URL":              {reflect.TypeOf((*q.URL)(nil)).Elem(), "", "EscapedFragment,EscapedPath,Hostname,IsAbs,MarshalBinary,Parse,Port,Query,Redacted,RequestURI,ResolveReference,String,UnmarshalBinary,setFragment,setPath"},
			"Userinfo":         {reflect.TypeOf((*q.Userinfo)(nil)).Elem(), "", "Password,String,Username"},
			"Values":           {reflect.TypeOf((*q.Values)(nil)).Elem(), "Add,Del,Encode,Get,Set", ""},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Parse":           reflect.ValueOf(q.Parse),
			"ParseQuery":      reflect.ValueOf(q.ParseQuery),
			"ParseRequestURI": reflect.ValueOf(q.ParseRequestURI),
			"PathEscape":      reflect.ValueOf(q.PathEscape),
			"PathUnescape":    reflect.ValueOf(q.PathUnescape),
			"QueryEscape":     reflect.ValueOf(q.QueryEscape),
			"QueryUnescape":   reflect.ValueOf(q.QueryUnescape),
			"User":            reflect.ValueOf(q.User),
			"UserPassword":    reflect.ValueOf(q.UserPassword),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
