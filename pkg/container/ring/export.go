// export by github.com/goplus/gossa/cmd/qexp

package ring

import (
	"container/ring"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("container/ring", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*container/ring.Ring).Do":     (*ring.Ring).Do,
	"(*container/ring.Ring).Len":    (*ring.Ring).Len,
	"(*container/ring.Ring).Link":   (*ring.Ring).Link,
	"(*container/ring.Ring).Move":   (*ring.Ring).Move,
	"(*container/ring.Ring).Next":   (*ring.Ring).Next,
	"(*container/ring.Ring).Prev":   (*ring.Ring).Prev,
	"(*container/ring.Ring).Unlink": (*ring.Ring).Unlink,
	"container/ring.New":            ring.New,
}

var typList = []interface{}{
	(*ring.Ring)(nil),
}
