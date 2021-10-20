// export by github.com/goplus/gossa/cmd/qexp

package maphash

import (
	"hash/maphash"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("hash/maphash", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*hash/maphash.Hash).BlockSize":   (*maphash.Hash).BlockSize,
	"(*hash/maphash.Hash).Reset":       (*maphash.Hash).Reset,
	"(*hash/maphash.Hash).Seed":        (*maphash.Hash).Seed,
	"(*hash/maphash.Hash).SetSeed":     (*maphash.Hash).SetSeed,
	"(*hash/maphash.Hash).Size":        (*maphash.Hash).Size,
	"(*hash/maphash.Hash).Sum":         (*maphash.Hash).Sum,
	"(*hash/maphash.Hash).Sum64":       (*maphash.Hash).Sum64,
	"(*hash/maphash.Hash).Write":       (*maphash.Hash).Write,
	"(*hash/maphash.Hash).WriteByte":   (*maphash.Hash).WriteByte,
	"(*hash/maphash.Hash).WriteString": (*maphash.Hash).WriteString,
	"hash/maphash.MakeSeed":            maphash.MakeSeed,
}

var typList = []interface{}{
	(*maphash.Hash)(nil),
	(*maphash.Seed)(nil),
}
