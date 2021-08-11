// export by github.com/goplus/interp/cmd/qexp

// +build go1.15

package url

import (
	"net/url"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("net/url", extMap_go115, typList_go115)
}

var extMap_go115 = map[string]interface{}{
	"(*net/url.URL).EscapedFragment": (*url.URL).EscapedFragment,
	"(*net/url.URL).Redacted":        (*url.URL).Redacted,
}

var typList_go115 = []interface{}{}
