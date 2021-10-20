// export by github.com/goplus/gossa/cmd/qexp

package scanner

import (
	"go/scanner"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("go/scanner", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*go/scanner.ErrorList).Add":             (*scanner.ErrorList).Add,
	"(*go/scanner.ErrorList).RemoveMultiples": (*scanner.ErrorList).RemoveMultiples,
	"(*go/scanner.ErrorList).Reset":           (*scanner.ErrorList).Reset,
	"(*go/scanner.Scanner).Init":              (*scanner.Scanner).Init,
	"(*go/scanner.Scanner).Scan":              (*scanner.Scanner).Scan,
	"(go/scanner.Error).Error":                (scanner.Error).Error,
	"(go/scanner.ErrorList).Err":              (scanner.ErrorList).Err,
	"(go/scanner.ErrorList).Error":            (scanner.ErrorList).Error,
	"(go/scanner.ErrorList).Len":              (scanner.ErrorList).Len,
	"(go/scanner.ErrorList).Less":             (scanner.ErrorList).Less,
	"(go/scanner.ErrorList).Sort":             (scanner.ErrorList).Sort,
	"(go/scanner.ErrorList).Swap":             (scanner.ErrorList).Swap,
	"go/scanner.PrintError":                   scanner.PrintError,
}

var typList = []interface{}{
	(*scanner.Error)(nil),
	(*scanner.ErrorHandler)(nil),
	(*scanner.ErrorList)(nil),
	(*scanner.Mode)(nil),
	(*scanner.Scanner)(nil),
}
