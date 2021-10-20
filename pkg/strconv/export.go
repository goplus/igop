// export by github.com/goplus/gossa/cmd/qexp

package strconv

import (
	"strconv"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("strconv", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*strconv.NumError).Error":        (*strconv.NumError).Error,
	"(*strconv.NumError).Unwrap":       (*strconv.NumError).Unwrap,
	"strconv.AppendBool":               strconv.AppendBool,
	"strconv.AppendFloat":              strconv.AppendFloat,
	"strconv.AppendInt":                strconv.AppendInt,
	"strconv.AppendQuote":              strconv.AppendQuote,
	"strconv.AppendQuoteRune":          strconv.AppendQuoteRune,
	"strconv.AppendQuoteRuneToASCII":   strconv.AppendQuoteRuneToASCII,
	"strconv.AppendQuoteRuneToGraphic": strconv.AppendQuoteRuneToGraphic,
	"strconv.AppendQuoteToASCII":       strconv.AppendQuoteToASCII,
	"strconv.AppendQuoteToGraphic":     strconv.AppendQuoteToGraphic,
	"strconv.AppendUint":               strconv.AppendUint,
	"strconv.Atoi":                     strconv.Atoi,
	"strconv.CanBackquote":             strconv.CanBackquote,
	"strconv.ErrRange":                 &strconv.ErrRange,
	"strconv.ErrSyntax":                &strconv.ErrSyntax,
	"strconv.FormatBool":               strconv.FormatBool,
	"strconv.FormatFloat":              strconv.FormatFloat,
	"strconv.FormatInt":                strconv.FormatInt,
	"strconv.FormatUint":               strconv.FormatUint,
	"strconv.IsGraphic":                strconv.IsGraphic,
	"strconv.IsPrint":                  strconv.IsPrint,
	"strconv.Itoa":                     strconv.Itoa,
	"strconv.ParseBool":                strconv.ParseBool,
	"strconv.ParseFloat":               strconv.ParseFloat,
	"strconv.ParseInt":                 strconv.ParseInt,
	"strconv.ParseUint":                strconv.ParseUint,
	"strconv.Quote":                    strconv.Quote,
	"strconv.QuoteRune":                strconv.QuoteRune,
	"strconv.QuoteRuneToASCII":         strconv.QuoteRuneToASCII,
	"strconv.QuoteRuneToGraphic":       strconv.QuoteRuneToGraphic,
	"strconv.QuoteToASCII":             strconv.QuoteToASCII,
	"strconv.QuoteToGraphic":           strconv.QuoteToGraphic,
	"strconv.Unquote":                  strconv.Unquote,
	"strconv.UnquoteChar":              strconv.UnquoteChar,
}

var typList = []interface{}{
	(*strconv.NumError)(nil),
}
