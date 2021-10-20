// export by github.com/goplus/gossa/cmd/qexp

package cookiejar

import (
	"net/http/cookiejar"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("net/http/cookiejar", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*net/http/cookiejar.Jar).Cookies":                  (*cookiejar.Jar).Cookies,
	"(*net/http/cookiejar.Jar).SetCookies":               (*cookiejar.Jar).SetCookies,
	"(net/http/cookiejar.PublicSuffixList).PublicSuffix": (cookiejar.PublicSuffixList).PublicSuffix,
	"(net/http/cookiejar.PublicSuffixList).String":       (cookiejar.PublicSuffixList).String,
	"net/http/cookiejar.New":                             cookiejar.New,
}

var typList = []interface{}{
	(*cookiejar.Jar)(nil),
	(*cookiejar.Options)(nil),
	(*cookiejar.PublicSuffixList)(nil),
}
