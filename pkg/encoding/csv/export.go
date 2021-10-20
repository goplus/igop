// export by github.com/goplus/gossa/cmd/qexp

package csv

import (
	"encoding/csv"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("encoding/csv", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*encoding/csv.ParseError).Error":  (*csv.ParseError).Error,
	"(*encoding/csv.ParseError).Unwrap": (*csv.ParseError).Unwrap,
	"(*encoding/csv.Reader).Read":       (*csv.Reader).Read,
	"(*encoding/csv.Reader).ReadAll":    (*csv.Reader).ReadAll,
	"(*encoding/csv.Writer).Error":      (*csv.Writer).Error,
	"(*encoding/csv.Writer).Flush":      (*csv.Writer).Flush,
	"(*encoding/csv.Writer).Write":      (*csv.Writer).Write,
	"(*encoding/csv.Writer).WriteAll":   (*csv.Writer).WriteAll,
	"encoding/csv.ErrBareQuote":         &csv.ErrBareQuote,
	"encoding/csv.ErrFieldCount":        &csv.ErrFieldCount,
	"encoding/csv.ErrQuote":             &csv.ErrQuote,
	"encoding/csv.ErrTrailingComma":     &csv.ErrTrailingComma,
	"encoding/csv.NewReader":            csv.NewReader,
	"encoding/csv.NewWriter":            csv.NewWriter,
}

var typList = []interface{}{
	(*csv.ParseError)(nil),
	(*csv.Reader)(nil),
	(*csv.Writer)(nil),
}
