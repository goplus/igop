// export by github.com/goplus/gossa/cmd/qexp

package regexp

import (
	"regexp"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("regexp", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*regexp.Regexp).Copy":                       (*regexp.Regexp).Copy,
	"(*regexp.Regexp).Expand":                     (*regexp.Regexp).Expand,
	"(*regexp.Regexp).ExpandString":               (*regexp.Regexp).ExpandString,
	"(*regexp.Regexp).Find":                       (*regexp.Regexp).Find,
	"(*regexp.Regexp).FindAll":                    (*regexp.Regexp).FindAll,
	"(*regexp.Regexp).FindAllIndex":               (*regexp.Regexp).FindAllIndex,
	"(*regexp.Regexp).FindAllString":              (*regexp.Regexp).FindAllString,
	"(*regexp.Regexp).FindAllStringIndex":         (*regexp.Regexp).FindAllStringIndex,
	"(*regexp.Regexp).FindAllStringSubmatch":      (*regexp.Regexp).FindAllStringSubmatch,
	"(*regexp.Regexp).FindAllStringSubmatchIndex": (*regexp.Regexp).FindAllStringSubmatchIndex,
	"(*regexp.Regexp).FindAllSubmatch":            (*regexp.Regexp).FindAllSubmatch,
	"(*regexp.Regexp).FindAllSubmatchIndex":       (*regexp.Regexp).FindAllSubmatchIndex,
	"(*regexp.Regexp).FindIndex":                  (*regexp.Regexp).FindIndex,
	"(*regexp.Regexp).FindReaderIndex":            (*regexp.Regexp).FindReaderIndex,
	"(*regexp.Regexp).FindReaderSubmatchIndex":    (*regexp.Regexp).FindReaderSubmatchIndex,
	"(*regexp.Regexp).FindString":                 (*regexp.Regexp).FindString,
	"(*regexp.Regexp).FindStringIndex":            (*regexp.Regexp).FindStringIndex,
	"(*regexp.Regexp).FindStringSubmatch":         (*regexp.Regexp).FindStringSubmatch,
	"(*regexp.Regexp).FindStringSubmatchIndex":    (*regexp.Regexp).FindStringSubmatchIndex,
	"(*regexp.Regexp).FindSubmatch":               (*regexp.Regexp).FindSubmatch,
	"(*regexp.Regexp).FindSubmatchIndex":          (*regexp.Regexp).FindSubmatchIndex,
	"(*regexp.Regexp).LiteralPrefix":              (*regexp.Regexp).LiteralPrefix,
	"(*regexp.Regexp).Longest":                    (*regexp.Regexp).Longest,
	"(*regexp.Regexp).Match":                      (*regexp.Regexp).Match,
	"(*regexp.Regexp).MatchReader":                (*regexp.Regexp).MatchReader,
	"(*regexp.Regexp).MatchString":                (*regexp.Regexp).MatchString,
	"(*regexp.Regexp).NumSubexp":                  (*regexp.Regexp).NumSubexp,
	"(*regexp.Regexp).ReplaceAll":                 (*regexp.Regexp).ReplaceAll,
	"(*regexp.Regexp).ReplaceAllFunc":             (*regexp.Regexp).ReplaceAllFunc,
	"(*regexp.Regexp).ReplaceAllLiteral":          (*regexp.Regexp).ReplaceAllLiteral,
	"(*regexp.Regexp).ReplaceAllLiteralString":    (*regexp.Regexp).ReplaceAllLiteralString,
	"(*regexp.Regexp).ReplaceAllString":           (*regexp.Regexp).ReplaceAllString,
	"(*regexp.Regexp).ReplaceAllStringFunc":       (*regexp.Regexp).ReplaceAllStringFunc,
	"(*regexp.Regexp).Split":                      (*regexp.Regexp).Split,
	"(*regexp.Regexp).String":                     (*regexp.Regexp).String,
	"(*regexp.Regexp).SubexpNames":                (*regexp.Regexp).SubexpNames,
	"regexp.Compile":                              regexp.Compile,
	"regexp.CompilePOSIX":                         regexp.CompilePOSIX,
	"regexp.Match":                                regexp.Match,
	"regexp.MatchReader":                          regexp.MatchReader,
	"regexp.MatchString":                          regexp.MatchString,
	"regexp.MustCompile":                          regexp.MustCompile,
	"regexp.MustCompilePOSIX":                     regexp.MustCompilePOSIX,
	"regexp.QuoteMeta":                            regexp.QuoteMeta,
}

var typList = []interface{}{
	(*regexp.Regexp)(nil),
}
