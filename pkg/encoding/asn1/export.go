// export by github.com/goplus/gossa/cmd/qexp

package asn1

import (
	q "encoding/asn1"

	"go/constant"
	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "asn1",
		Path: "encoding/asn1",
		Deps: map[string]string{
			"bytes":         "bytes",
			"errors":        "errors",
			"fmt":           "fmt",
			"math":          "math",
			"math/big":      "big",
			"reflect":       "reflect",
			"sort":          "sort",
			"strconv":       "strconv",
			"strings":       "strings",
			"time":          "time",
			"unicode/utf16": "utf16",
			"unicode/utf8":  "utf8",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]gossa.NamedType{
			"BitString":        {reflect.TypeOf((*q.BitString)(nil)).Elem(), "At,RightAlign", ""},
			"Enumerated":       {reflect.TypeOf((*q.Enumerated)(nil)).Elem(), "", ""},
			"Flag":             {reflect.TypeOf((*q.Flag)(nil)).Elem(), "", ""},
			"ObjectIdentifier": {reflect.TypeOf((*q.ObjectIdentifier)(nil)).Elem(), "Equal,String", ""},
			"RawContent":       {reflect.TypeOf((*q.RawContent)(nil)).Elem(), "", ""},
			"RawValue":         {reflect.TypeOf((*q.RawValue)(nil)).Elem(), "", ""},
			"StructuralError":  {reflect.TypeOf((*q.StructuralError)(nil)).Elem(), "Error", ""},
			"SyntaxError":      {reflect.TypeOf((*q.SyntaxError)(nil)).Elem(), "Error", ""},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"NullBytes":    reflect.ValueOf(&q.NullBytes),
			"NullRawValue": reflect.ValueOf(&q.NullRawValue),
		},
		Funcs: map[string]reflect.Value{
			"Marshal":             reflect.ValueOf(q.Marshal),
			"MarshalWithParams":   reflect.ValueOf(q.MarshalWithParams),
			"Unmarshal":           reflect.ValueOf(q.Unmarshal),
			"UnmarshalWithParams": reflect.ValueOf(q.UnmarshalWithParams),
		},
		TypedConsts: map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{
			"ClassApplication":     {"untyped int", constant.MakeInt64(1)},
			"ClassContextSpecific": {"untyped int", constant.MakeInt64(2)},
			"ClassPrivate":         {"untyped int", constant.MakeInt64(3)},
			"ClassUniversal":       {"untyped int", constant.MakeInt64(0)},
			"TagBMPString":         {"untyped int", constant.MakeInt64(30)},
			"TagBitString":         {"untyped int", constant.MakeInt64(3)},
			"TagBoolean":           {"untyped int", constant.MakeInt64(1)},
			"TagEnum":              {"untyped int", constant.MakeInt64(10)},
			"TagGeneralString":     {"untyped int", constant.MakeInt64(27)},
			"TagGeneralizedTime":   {"untyped int", constant.MakeInt64(24)},
			"TagIA5String":         {"untyped int", constant.MakeInt64(22)},
			"TagInteger":           {"untyped int", constant.MakeInt64(2)},
			"TagNull":              {"untyped int", constant.MakeInt64(5)},
			"TagNumericString":     {"untyped int", constant.MakeInt64(18)},
			"TagOID":               {"untyped int", constant.MakeInt64(6)},
			"TagOctetString":       {"untyped int", constant.MakeInt64(4)},
			"TagPrintableString":   {"untyped int", constant.MakeInt64(19)},
			"TagSequence":          {"untyped int", constant.MakeInt64(16)},
			"TagSet":               {"untyped int", constant.MakeInt64(17)},
			"TagT61String":         {"untyped int", constant.MakeInt64(20)},
			"TagUTCTime":           {"untyped int", constant.MakeInt64(23)},
			"TagUTF8String":        {"untyped int", constant.MakeInt64(12)},
		},
	})
}
