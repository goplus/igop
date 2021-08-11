// export by github.com/goplus/interp/cmd/qexp

package quick

import (
	"testing/quick"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("testing/quick", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*testing/quick.CheckEqualError).Error": (*quick.CheckEqualError).Error,
	"(*testing/quick.CheckError).Error":      (*quick.CheckError).Error,
	"(testing/quick.Generator).Generate":     (quick.Generator).Generate,
	"(testing/quick.SetupError).Error":       (quick.SetupError).Error,
	"testing/quick.Check":                    quick.Check,
	"testing/quick.CheckEqual":               quick.CheckEqual,
	"testing/quick.Value":                    quick.Value,
}

var typList = []interface{}{
	(*quick.CheckEqualError)(nil),
	(*quick.CheckError)(nil),
	(*quick.Config)(nil),
	(*quick.Generator)(nil),
	(*quick.SetupError)(nil),
}
