// export by github.com/goplus/gossa/cmd/qexp

package strings

import (
	"strings"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("strings", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*strings.Builder).Cap":          (*strings.Builder).Cap,
	"(*strings.Builder).Grow":         (*strings.Builder).Grow,
	"(*strings.Builder).Len":          (*strings.Builder).Len,
	"(*strings.Builder).Reset":        (*strings.Builder).Reset,
	"(*strings.Builder).String":       (*strings.Builder).String,
	"(*strings.Builder).Write":        (*strings.Builder).Write,
	"(*strings.Builder).WriteByte":    (*strings.Builder).WriteByte,
	"(*strings.Builder).WriteRune":    (*strings.Builder).WriteRune,
	"(*strings.Builder).WriteString":  (*strings.Builder).WriteString,
	"(*strings.Reader).Len":           (*strings.Reader).Len,
	"(*strings.Reader).Read":          (*strings.Reader).Read,
	"(*strings.Reader).ReadAt":        (*strings.Reader).ReadAt,
	"(*strings.Reader).ReadByte":      (*strings.Reader).ReadByte,
	"(*strings.Reader).ReadRune":      (*strings.Reader).ReadRune,
	"(*strings.Reader).Reset":         (*strings.Reader).Reset,
	"(*strings.Reader).Seek":          (*strings.Reader).Seek,
	"(*strings.Reader).Size":          (*strings.Reader).Size,
	"(*strings.Reader).UnreadByte":    (*strings.Reader).UnreadByte,
	"(*strings.Reader).UnreadRune":    (*strings.Reader).UnreadRune,
	"(*strings.Reader).WriteTo":       (*strings.Reader).WriteTo,
	"(*strings.Replacer).Replace":     (*strings.Replacer).Replace,
	"(*strings.Replacer).WriteString": (*strings.Replacer).WriteString,
	"strings.Compare":                 strings.Compare,
	"strings.Contains":                strings.Contains,
	"strings.ContainsAny":             strings.ContainsAny,
	"strings.ContainsRune":            strings.ContainsRune,
	"strings.Count":                   strings.Count,
	"strings.EqualFold":               strings.EqualFold,
	"strings.Fields":                  strings.Fields,
	"strings.FieldsFunc":              strings.FieldsFunc,
	"strings.HasPrefix":               strings.HasPrefix,
	"strings.HasSuffix":               strings.HasSuffix,
	"strings.Index":                   strings.Index,
	"strings.IndexAny":                strings.IndexAny,
	"strings.IndexByte":               strings.IndexByte,
	"strings.IndexFunc":               strings.IndexFunc,
	"strings.IndexRune":               strings.IndexRune,
	"strings.Join":                    strings.Join,
	"strings.LastIndex":               strings.LastIndex,
	"strings.LastIndexAny":            strings.LastIndexAny,
	"strings.LastIndexByte":           strings.LastIndexByte,
	"strings.LastIndexFunc":           strings.LastIndexFunc,
	"strings.Map":                     strings.Map,
	"strings.NewReader":               strings.NewReader,
	"strings.NewReplacer":             strings.NewReplacer,
	"strings.Repeat":                  strings.Repeat,
	"strings.Replace":                 strings.Replace,
	"strings.ReplaceAll":              strings.ReplaceAll,
	"strings.Split":                   strings.Split,
	"strings.SplitAfter":              strings.SplitAfter,
	"strings.SplitAfterN":             strings.SplitAfterN,
	"strings.SplitN":                  strings.SplitN,
	"strings.Title":                   strings.Title,
	"strings.ToLower":                 strings.ToLower,
	"strings.ToLowerSpecial":          strings.ToLowerSpecial,
	"strings.ToTitle":                 strings.ToTitle,
	"strings.ToTitleSpecial":          strings.ToTitleSpecial,
	"strings.ToUpper":                 strings.ToUpper,
	"strings.ToUpperSpecial":          strings.ToUpperSpecial,
	"strings.ToValidUTF8":             strings.ToValidUTF8,
	"strings.Trim":                    strings.Trim,
	"strings.TrimFunc":                strings.TrimFunc,
	"strings.TrimLeft":                strings.TrimLeft,
	"strings.TrimLeftFunc":            strings.TrimLeftFunc,
	"strings.TrimPrefix":              strings.TrimPrefix,
	"strings.TrimRight":               strings.TrimRight,
	"strings.TrimRightFunc":           strings.TrimRightFunc,
	"strings.TrimSpace":               strings.TrimSpace,
	"strings.TrimSuffix":              strings.TrimSuffix,
}

var typList = []interface{}{
	(*strings.Builder)(nil),
	(*strings.Reader)(nil),
	(*strings.Replacer)(nil),
}
