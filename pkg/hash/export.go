// export by github.com/goplus/interp/cmd/qexp

package hash

import (
	"hash"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("hash", extMap, typList)
}

var extMap = map[string]interface{}{
	"(hash.Hash).BlockSize":   (hash.Hash).BlockSize,
	"(hash.Hash).Reset":       (hash.Hash).Reset,
	"(hash.Hash).Size":        (hash.Hash).Size,
	"(hash.Hash).Sum":         (hash.Hash).Sum,
	"(hash.Hash).Write":       (hash.Hash).Write,
	"(hash.Hash32).BlockSize": (hash.Hash32).BlockSize,
	"(hash.Hash32).Reset":     (hash.Hash32).Reset,
	"(hash.Hash32).Size":      (hash.Hash32).Size,
	"(hash.Hash32).Sum":       (hash.Hash32).Sum,
	"(hash.Hash32).Sum32":     (hash.Hash32).Sum32,
	"(hash.Hash32).Write":     (hash.Hash32).Write,
	"(hash.Hash64).BlockSize": (hash.Hash64).BlockSize,
	"(hash.Hash64).Reset":     (hash.Hash64).Reset,
	"(hash.Hash64).Size":      (hash.Hash64).Size,
	"(hash.Hash64).Sum":       (hash.Hash64).Sum,
	"(hash.Hash64).Sum64":     (hash.Hash64).Sum64,
	"(hash.Hash64).Write":     (hash.Hash64).Write,
}

var typList = []interface{}{
	(*hash.Hash)(nil),
	(*hash.Hash32)(nil),
	(*hash.Hash64)(nil),
}
