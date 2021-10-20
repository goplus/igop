// export by github.com/goplus/gossa/cmd/qexp

package suffixarray

import (
	"index/suffixarray"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("index/suffixarray", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*index/suffixarray.Index).Bytes":        (*suffixarray.Index).Bytes,
	"(*index/suffixarray.Index).FindAllIndex": (*suffixarray.Index).FindAllIndex,
	"(*index/suffixarray.Index).Lookup":       (*suffixarray.Index).Lookup,
	"(*index/suffixarray.Index).Read":         (*suffixarray.Index).Read,
	"(*index/suffixarray.Index).Write":        (*suffixarray.Index).Write,
	"index/suffixarray.New":                   suffixarray.New,
}

var typList = []interface{}{
	(*suffixarray.Index)(nil),
}
