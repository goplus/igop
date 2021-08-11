// export by github.com/goplus/interp/cmd/qexp

package utf16

import (
	"unicode/utf16"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("unicode/utf16", extMap, typList)
}

var extMap = map[string]interface{}{
	"unicode/utf16.Decode":      utf16.Decode,
	"unicode/utf16.DecodeRune":  utf16.DecodeRune,
	"unicode/utf16.Encode":      utf16.Encode,
	"unicode/utf16.EncodeRune":  utf16.EncodeRune,
	"unicode/utf16.IsSurrogate": utf16.IsSurrogate,
}

var typList = []interface{}{}
