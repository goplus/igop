// export by github.com/goplus/interp/cmd/qexp

package url

import (
	"net/url"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("net/url", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*net/url.Error).Error":           (*url.Error).Error,
	"(*net/url.Error).Temporary":       (*url.Error).Temporary,
	"(*net/url.Error).Timeout":         (*url.Error).Timeout,
	"(*net/url.Error).Unwrap":          (*url.Error).Unwrap,
	"(*net/url.URL).EscapedPath":       (*url.URL).EscapedPath,
	"(*net/url.URL).Hostname":          (*url.URL).Hostname,
	"(*net/url.URL).IsAbs":             (*url.URL).IsAbs,
	"(*net/url.URL).MarshalBinary":     (*url.URL).MarshalBinary,
	"(*net/url.URL).Parse":             (*url.URL).Parse,
	"(*net/url.URL).Port":              (*url.URL).Port,
	"(*net/url.URL).Query":             (*url.URL).Query,
	"(*net/url.URL).RequestURI":        (*url.URL).RequestURI,
	"(*net/url.URL).ResolveReference":  (*url.URL).ResolveReference,
	"(*net/url.URL).String":            (*url.URL).String,
	"(*net/url.URL).UnmarshalBinary":   (*url.URL).UnmarshalBinary,
	"(*net/url.Userinfo).Password":     (*url.Userinfo).Password,
	"(*net/url.Userinfo).String":       (*url.Userinfo).String,
	"(*net/url.Userinfo).Username":     (*url.Userinfo).Username,
	"(net/url.EscapeError).Error":      (url.EscapeError).Error,
	"(net/url.InvalidHostError).Error": (url.InvalidHostError).Error,
	"(net/url.Values).Add":             (url.Values).Add,
	"(net/url.Values).Del":             (url.Values).Del,
	"(net/url.Values).Encode":          (url.Values).Encode,
	"(net/url.Values).Get":             (url.Values).Get,
	"(net/url.Values).Set":             (url.Values).Set,
	"net/url.Parse":                    url.Parse,
	"net/url.ParseQuery":               url.ParseQuery,
	"net/url.ParseRequestURI":          url.ParseRequestURI,
	"net/url.PathEscape":               url.PathEscape,
	"net/url.PathUnescape":             url.PathUnescape,
	"net/url.QueryEscape":              url.QueryEscape,
	"net/url.QueryUnescape":            url.QueryUnescape,
	"net/url.User":                     url.User,
	"net/url.UserPassword":             url.UserPassword,
}

var typList = []interface{}{
	(*url.Error)(nil),
	(*url.EscapeError)(nil),
	(*url.InvalidHostError)(nil),
	(*url.URL)(nil),
	(*url.Userinfo)(nil),
	(*url.Values)(nil),
}
