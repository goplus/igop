// export by github.com/goplus/gossa/cmd/qexp

package crypto

import (
	q "crypto"

	"go/constant"
	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "crypto",
		Path: "crypto",
		Deps: map[string]string{
			"hash":    "hash",
			"io":      "io",
			"strconv": "strconv",
		},
		Interfaces: map[string]reflect.Type{
			"Decrypter":     reflect.TypeOf((*q.Decrypter)(nil)).Elem(),
			"DecrypterOpts": reflect.TypeOf((*q.DecrypterOpts)(nil)).Elem(),
			"PrivateKey":    reflect.TypeOf((*q.PrivateKey)(nil)).Elem(),
			"PublicKey":     reflect.TypeOf((*q.PublicKey)(nil)).Elem(),
			"Signer":        reflect.TypeOf((*q.Signer)(nil)).Elem(),
			"SignerOpts":    reflect.TypeOf((*q.SignerOpts)(nil)).Elem(),
		},
		NamedTypes: map[string]gossa.NamedType{
			"Hash": {reflect.TypeOf((*q.Hash)(nil)).Elem(), "Available,HashFunc,New,Size,String", ""},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"RegisterHash": reflect.ValueOf(q.RegisterHash),
		},
		TypedConsts: map[string]gossa.TypedConst{
			"BLAKE2b_256": {reflect.TypeOf(q.BLAKE2b_256), constant.MakeInt64(17)},
			"BLAKE2b_384": {reflect.TypeOf(q.BLAKE2b_384), constant.MakeInt64(18)},
			"BLAKE2b_512": {reflect.TypeOf(q.BLAKE2b_512), constant.MakeInt64(19)},
			"BLAKE2s_256": {reflect.TypeOf(q.BLAKE2s_256), constant.MakeInt64(16)},
			"MD4":         {reflect.TypeOf(q.MD4), constant.MakeInt64(1)},
			"MD5":         {reflect.TypeOf(q.MD5), constant.MakeInt64(2)},
			"MD5SHA1":     {reflect.TypeOf(q.MD5SHA1), constant.MakeInt64(8)},
			"RIPEMD160":   {reflect.TypeOf(q.RIPEMD160), constant.MakeInt64(9)},
			"SHA1":        {reflect.TypeOf(q.SHA1), constant.MakeInt64(3)},
			"SHA224":      {reflect.TypeOf(q.SHA224), constant.MakeInt64(4)},
			"SHA256":      {reflect.TypeOf(q.SHA256), constant.MakeInt64(5)},
			"SHA384":      {reflect.TypeOf(q.SHA384), constant.MakeInt64(6)},
			"SHA3_224":    {reflect.TypeOf(q.SHA3_224), constant.MakeInt64(10)},
			"SHA3_256":    {reflect.TypeOf(q.SHA3_256), constant.MakeInt64(11)},
			"SHA3_384":    {reflect.TypeOf(q.SHA3_384), constant.MakeInt64(12)},
			"SHA3_512":    {reflect.TypeOf(q.SHA3_512), constant.MakeInt64(13)},
			"SHA512":      {reflect.TypeOf(q.SHA512), constant.MakeInt64(7)},
			"SHA512_224":  {reflect.TypeOf(q.SHA512_224), constant.MakeInt64(14)},
			"SHA512_256":  {reflect.TypeOf(q.SHA512_256), constant.MakeInt64(15)},
		},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}
