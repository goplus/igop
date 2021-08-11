// export by github.com/goplus/interp/cmd/qexp

// +build go1.15

package regexp

import (
	"regexp"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("regexp", extMap_go115, typList_go115)
}

var extMap_go115 = map[string]interface{}{
	"(*regexp.Regexp).SubexpIndex": (*regexp.Regexp).SubexpIndex,
}

var typList_go115 = []interface{}{}
