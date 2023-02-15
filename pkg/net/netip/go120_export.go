// export by github.com/goplus/igop/cmd/qexp

//go:build go1.20
// +build go1.20

package netip

import (
	q "net/netip"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "netip",
		Path: "net/netip",
		Deps: map[string]string{
			"errors":           "errors",
			"internal/bytealg": "bytealg",
			"internal/intern":  "intern",
			"internal/itoa":    "itoa",
			"math":             "math",
			"math/bits":        "bits",
			"strconv":          "strconv",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Addr":     reflect.TypeOf((*q.Addr)(nil)).Elem(),
			"AddrPort": reflect.TypeOf((*q.AddrPort)(nil)).Elem(),
			"Prefix":   reflect.TypeOf((*q.Prefix)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"AddrFrom16":              reflect.ValueOf(q.AddrFrom16),
			"AddrFrom4":               reflect.ValueOf(q.AddrFrom4),
			"AddrFromSlice":           reflect.ValueOf(q.AddrFromSlice),
			"AddrPortFrom":            reflect.ValueOf(q.AddrPortFrom),
			"IPv4Unspecified":         reflect.ValueOf(q.IPv4Unspecified),
			"IPv6LinkLocalAllNodes":   reflect.ValueOf(q.IPv6LinkLocalAllNodes),
			"IPv6LinkLocalAllRouters": reflect.ValueOf(q.IPv6LinkLocalAllRouters),
			"IPv6Loopback":            reflect.ValueOf(q.IPv6Loopback),
			"IPv6Unspecified":         reflect.ValueOf(q.IPv6Unspecified),
			"MustParseAddr":           reflect.ValueOf(q.MustParseAddr),
			"MustParseAddrPort":       reflect.ValueOf(q.MustParseAddrPort),
			"MustParsePrefix":         reflect.ValueOf(q.MustParsePrefix),
			"ParseAddr":               reflect.ValueOf(q.ParseAddr),
			"ParseAddrPort":           reflect.ValueOf(q.ParseAddrPort),
			"ParsePrefix":             reflect.ValueOf(q.ParsePrefix),
			"PrefixFrom":              reflect.ValueOf(q.PrefixFrom),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
