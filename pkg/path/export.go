// export by github.com/goplus/interp/cmd/qexp

package path

import (
	"path"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("path", extMap, typList)
}

var extMap = map[string]interface{}{
	"path.Base":          path.Base,
	"path.Clean":         path.Clean,
	"path.Dir":           path.Dir,
	"path.ErrBadPattern": &path.ErrBadPattern,
	"path.Ext":           path.Ext,
	"path.IsAbs":         path.IsAbs,
	"path.Join":          path.Join,
	"path.Match":         path.Match,
	"path.Split":         path.Split,
}

var typList = []interface{}{}
