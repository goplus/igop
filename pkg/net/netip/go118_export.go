// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

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
		NamedTypes: map[string]igop.NamedType{
			"Addr":     {reflect.TypeOf((*q.Addr)(nil)).Elem(), "AppendTo,As16,As4,AsSlice,BitLen,Compare,Is4,Is4In6,Is6,IsGlobalUnicast,IsInterfaceLocalMulticast,IsLinkLocalMulticast,IsLinkLocalUnicast,IsLoopback,IsMulticast,IsPrivate,IsUnspecified,IsValid,Less,MarshalBinary,MarshalText,Next,Prefix,Prev,String,StringExpanded,Unmap,WithZone,Zone,appendTo4,appendTo6,hasZone,isZero,lessOrEq,marshalBinaryWithTrailingBytes,string4,string6,v4,v6,v6u16,withoutZone", "UnmarshalBinary,UnmarshalText"},
			"AddrPort": {reflect.TypeOf((*q.AddrPort)(nil)).Elem(), "Addr,AppendTo,IsValid,MarshalBinary,MarshalText,Port,String,isZero", "UnmarshalBinary,UnmarshalText"},
			"Prefix":   {reflect.TypeOf((*q.Prefix)(nil)).Elem(), "Addr,AppendTo,Bits,Contains,IsSingleIP,IsValid,MarshalBinary,MarshalText,Masked,Overlaps,String,isZero", "UnmarshalBinary,UnmarshalText"},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"AddrFrom16":            reflect.ValueOf(q.AddrFrom16),
			"AddrFrom4":             reflect.ValueOf(q.AddrFrom4),
			"AddrFromSlice":         reflect.ValueOf(q.AddrFromSlice),
			"AddrPortFrom":          reflect.ValueOf(q.AddrPortFrom),
			"IPv4Unspecified":       reflect.ValueOf(q.IPv4Unspecified),
			"IPv6LinkLocalAllNodes": reflect.ValueOf(q.IPv6LinkLocalAllNodes),
			"IPv6Unspecified":       reflect.ValueOf(q.IPv6Unspecified),
			"MustParseAddr":         reflect.ValueOf(q.MustParseAddr),
			"MustParseAddrPort":     reflect.ValueOf(q.MustParseAddrPort),
			"MustParsePrefix":       reflect.ValueOf(q.MustParsePrefix),
			"ParseAddr":             reflect.ValueOf(q.ParseAddr),
			"ParseAddrPort":         reflect.ValueOf(q.ParseAddrPort),
			"ParsePrefix":           reflect.ValueOf(q.ParsePrefix),
			"PrefixFrom":            reflect.ValueOf(q.PrefixFrom),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
