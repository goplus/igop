// export by github.com/goplus/gossa/cmd/qexp

package quotedprintable

import (
	"mime/quotedprintable"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("mime/quotedprintable", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*mime/quotedprintable.Reader).Read":  (*quotedprintable.Reader).Read,
	"(*mime/quotedprintable.Writer).Close": (*quotedprintable.Writer).Close,
	"(*mime/quotedprintable.Writer).Write": (*quotedprintable.Writer).Write,
	"mime/quotedprintable.NewReader":       quotedprintable.NewReader,
	"mime/quotedprintable.NewWriter":       quotedprintable.NewWriter,
}

var typList = []interface{}{
	(*quotedprintable.Reader)(nil),
	(*quotedprintable.Writer)(nil),
}
