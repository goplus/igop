// export by github.com/goplus/gossa/cmd/qexp

package scanner

import (
	"text/scanner"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("text/scanner", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*text/scanner.Position).IsValid":  (*scanner.Position).IsValid,
	"(*text/scanner.Scanner).Init":      (*scanner.Scanner).Init,
	"(*text/scanner.Scanner).IsValid":   (*scanner.Scanner).IsValid,
	"(*text/scanner.Scanner).Next":      (*scanner.Scanner).Next,
	"(*text/scanner.Scanner).Peek":      (*scanner.Scanner).Peek,
	"(*text/scanner.Scanner).Pos":       (*scanner.Scanner).Pos,
	"(*text/scanner.Scanner).Scan":      (*scanner.Scanner).Scan,
	"(*text/scanner.Scanner).TokenText": (*scanner.Scanner).TokenText,
	"(text/scanner.Position).String":    (scanner.Position).String,
	"(text/scanner.Scanner).String":     (scanner.Scanner).String,
	"text/scanner.TokenString":          scanner.TokenString,
}

var typList = []interface{}{
	(*scanner.Position)(nil),
	(*scanner.Scanner)(nil),
}
