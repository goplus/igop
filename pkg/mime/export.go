// export by github.com/goplus/gossa/cmd/qexp

package mime

import (
	"mime"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("mime", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*mime.WordDecoder).Decode":       (*mime.WordDecoder).Decode,
	"(*mime.WordDecoder).DecodeHeader": (*mime.WordDecoder).DecodeHeader,
	"(mime.WordEncoder).Encode":        (mime.WordEncoder).Encode,
	"mime.AddExtensionType":            mime.AddExtensionType,
	"mime.ErrInvalidMediaParameter":    &mime.ErrInvalidMediaParameter,
	"mime.ExtensionsByType":            mime.ExtensionsByType,
	"mime.FormatMediaType":             mime.FormatMediaType,
	"mime.ParseMediaType":              mime.ParseMediaType,
	"mime.TypeByExtension":             mime.TypeByExtension,
}

var typList = []interface{}{
	(*mime.WordDecoder)(nil),
	(*mime.WordEncoder)(nil),
}
