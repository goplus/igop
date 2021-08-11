// export by github.com/goplus/interp/cmd/qexp

package context

import (
	"context"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("context", extMap, typList)
}

var extMap = map[string]interface{}{
	"(context.Context).Deadline": (context.Context).Deadline,
	"(context.Context).Done":     (context.Context).Done,
	"(context.Context).Err":      (context.Context).Err,
	"(context.Context).Value":    (context.Context).Value,
	"context.Background":         context.Background,
	"context.Canceled":           &context.Canceled,
	"context.DeadlineExceeded":   &context.DeadlineExceeded,
	"context.TODO":               context.TODO,
	"context.WithCancel":         context.WithCancel,
	"context.WithDeadline":       context.WithDeadline,
	"context.WithTimeout":        context.WithTimeout,
	"context.WithValue":          context.WithValue,
}

var typList = []interface{}{
	(*context.CancelFunc)(nil),
	(*context.Context)(nil),
}
