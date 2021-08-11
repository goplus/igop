// export by github.com/goplus/interp/cmd/qexp

package list

import (
	"container/list"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("container/list", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*container/list.Element).Next":       (*list.Element).Next,
	"(*container/list.Element).Prev":       (*list.Element).Prev,
	"(*container/list.List).Back":          (*list.List).Back,
	"(*container/list.List).Front":         (*list.List).Front,
	"(*container/list.List).Init":          (*list.List).Init,
	"(*container/list.List).InsertAfter":   (*list.List).InsertAfter,
	"(*container/list.List).InsertBefore":  (*list.List).InsertBefore,
	"(*container/list.List).Len":           (*list.List).Len,
	"(*container/list.List).MoveAfter":     (*list.List).MoveAfter,
	"(*container/list.List).MoveBefore":    (*list.List).MoveBefore,
	"(*container/list.List).MoveToBack":    (*list.List).MoveToBack,
	"(*container/list.List).MoveToFront":   (*list.List).MoveToFront,
	"(*container/list.List).PushBack":      (*list.List).PushBack,
	"(*container/list.List).PushBackList":  (*list.List).PushBackList,
	"(*container/list.List).PushFront":     (*list.List).PushFront,
	"(*container/list.List).PushFrontList": (*list.List).PushFrontList,
	"(*container/list.List).Remove":        (*list.List).Remove,
	"container/list.New":                   list.New,
}

var typList = []interface{}{
	(*list.Element)(nil),
	(*list.List)(nil),
}
