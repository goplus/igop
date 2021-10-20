// export by github.com/goplus/gossa/cmd/qexp

//go:build go1.15
// +build go1.15

package url

import (
	"net/url"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("net/url", extMap_go115, typList_go115)
}

var extMap_go115 = map[string]interface{}{
	"(*net/url.URL).EscapedFragment": (*url.URL).EscapedFragment,
	"(*net/url.URL).Redacted":        (*url.URL).Redacted,
}

var typList_go115 = []interface{}{}
