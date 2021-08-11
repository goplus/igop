// export by github.com/goplus/interp/cmd/qexp

package errors

import (
	"errors"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("errors", extMap, typList)
}

var extMap = map[string]interface{}{
	"errors.As":     errors.As,
	"errors.Is":     errors.Is,
	"errors.New":    errors.New,
	"errors.Unwrap": errors.Unwrap,
}

var typList = []interface{}{}
