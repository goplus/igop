// export by github.com/goplus/interp/cmd/qexp

//go:build go1.16
// +build go1.16

package unicode

import (
	"unicode"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("unicode", extMap_go116, typList_go116)
}

var extMap_go116 = map[string]interface{}{
	"unicode.Chorasmian":          &unicode.Chorasmian,
	"unicode.Dives_Akuru":         &unicode.Dives_Akuru,
	"unicode.Khitan_Small_Script": &unicode.Khitan_Small_Script,
	"unicode.Yezidi":              &unicode.Yezidi,
}

var typList_go116 = []interface{}{}
