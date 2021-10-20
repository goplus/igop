// export by github.com/goplus/gossa/cmd/qexp

package utf8

import (
	"unicode/utf8"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("unicode/utf8", extMap, typList)
}

var extMap = map[string]interface{}{
	"unicode/utf8.DecodeLastRune":         utf8.DecodeLastRune,
	"unicode/utf8.DecodeLastRuneInString": utf8.DecodeLastRuneInString,
	"unicode/utf8.DecodeRune":             utf8.DecodeRune,
	"unicode/utf8.DecodeRuneInString":     utf8.DecodeRuneInString,
	"unicode/utf8.EncodeRune":             utf8.EncodeRune,
	"unicode/utf8.FullRune":               utf8.FullRune,
	"unicode/utf8.FullRuneInString":       utf8.FullRuneInString,
	"unicode/utf8.RuneCount":              utf8.RuneCount,
	"unicode/utf8.RuneCountInString":      utf8.RuneCountInString,
	"unicode/utf8.RuneLen":                utf8.RuneLen,
	"unicode/utf8.RuneStart":              utf8.RuneStart,
	"unicode/utf8.Valid":                  utf8.Valid,
	"unicode/utf8.ValidRune":              utf8.ValidRune,
	"unicode/utf8.ValidString":            utf8.ValidString,
}

var typList = []interface{}{}
