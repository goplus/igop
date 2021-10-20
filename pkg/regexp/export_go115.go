// export by github.com/goplus/gossa/cmd/qexp

//go:build go1.15
// +build go1.15

package regexp

import (
	"regexp"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("regexp", extMap_go115, typList_go115)
}

var extMap_go115 = map[string]interface{}{
	"(*regexp.Regexp).SubexpIndex": (*regexp.Regexp).SubexpIndex,
}

var typList_go115 = []interface{}{}
