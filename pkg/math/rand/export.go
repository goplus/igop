// export by github.com/goplus/interp/cmd/qexp

package rand

import (
	"math/rand"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("math/rand", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*math/rand.Rand).ExpFloat64":  (*rand.Rand).ExpFloat64,
	"(*math/rand.Rand).Float32":     (*rand.Rand).Float32,
	"(*math/rand.Rand).Float64":     (*rand.Rand).Float64,
	"(*math/rand.Rand).Int":         (*rand.Rand).Int,
	"(*math/rand.Rand).Int31":       (*rand.Rand).Int31,
	"(*math/rand.Rand).Int31n":      (*rand.Rand).Int31n,
	"(*math/rand.Rand).Int63":       (*rand.Rand).Int63,
	"(*math/rand.Rand).Int63n":      (*rand.Rand).Int63n,
	"(*math/rand.Rand).Intn":        (*rand.Rand).Intn,
	"(*math/rand.Rand).NormFloat64": (*rand.Rand).NormFloat64,
	"(*math/rand.Rand).Perm":        (*rand.Rand).Perm,
	"(*math/rand.Rand).Read":        (*rand.Rand).Read,
	"(*math/rand.Rand).Seed":        (*rand.Rand).Seed,
	"(*math/rand.Rand).Shuffle":     (*rand.Rand).Shuffle,
	"(*math/rand.Rand).Uint32":      (*rand.Rand).Uint32,
	"(*math/rand.Rand).Uint64":      (*rand.Rand).Uint64,
	"(*math/rand.Zipf).Uint64":      (*rand.Zipf).Uint64,
	"(math/rand.Source).Int63":      (rand.Source).Int63,
	"(math/rand.Source).Seed":       (rand.Source).Seed,
	"(math/rand.Source64).Int63":    (rand.Source64).Int63,
	"(math/rand.Source64).Seed":     (rand.Source64).Seed,
	"(math/rand.Source64).Uint64":   (rand.Source64).Uint64,
	"math/rand.ExpFloat64":          rand.ExpFloat64,
	"math/rand.Float32":             rand.Float32,
	"math/rand.Float64":             rand.Float64,
	"math/rand.Int":                 rand.Int,
	"math/rand.Int31":               rand.Int31,
	"math/rand.Int31n":              rand.Int31n,
	"math/rand.Int63":               rand.Int63,
	"math/rand.Int63n":              rand.Int63n,
	"math/rand.Intn":                rand.Intn,
	"math/rand.New":                 rand.New,
	"math/rand.NewSource":           rand.NewSource,
	"math/rand.NewZipf":             rand.NewZipf,
	"math/rand.NormFloat64":         rand.NormFloat64,
	"math/rand.Perm":                rand.Perm,
	"math/rand.Read":                rand.Read,
	"math/rand.Seed":                rand.Seed,
	"math/rand.Shuffle":             rand.Shuffle,
	"math/rand.Uint32":              rand.Uint32,
	"math/rand.Uint64":              rand.Uint64,
}

var typList = []interface{}{
	(*rand.Rand)(nil),
	(*rand.Source)(nil),
	(*rand.Source64)(nil),
	(*rand.Zipf)(nil),
}
