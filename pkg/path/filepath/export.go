// export by github.com/goplus/interp/cmd/qexp

package filepath

import (
	"path/filepath"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("path/filepath", extMap, typList)
}

var extMap = map[string]interface{}{
	"path/filepath.Abs":           filepath.Abs,
	"path/filepath.Base":          filepath.Base,
	"path/filepath.Clean":         filepath.Clean,
	"path/filepath.Dir":           filepath.Dir,
	"path/filepath.ErrBadPattern": &filepath.ErrBadPattern,
	"path/filepath.EvalSymlinks":  filepath.EvalSymlinks,
	"path/filepath.Ext":           filepath.Ext,
	"path/filepath.FromSlash":     filepath.FromSlash,
	"path/filepath.Glob":          filepath.Glob,
	"path/filepath.HasPrefix":     filepath.HasPrefix,
	"path/filepath.IsAbs":         filepath.IsAbs,
	"path/filepath.Join":          filepath.Join,
	"path/filepath.Match":         filepath.Match,
	"path/filepath.Rel":           filepath.Rel,
	"path/filepath.SkipDir":       &filepath.SkipDir,
	"path/filepath.Split":         filepath.Split,
	"path/filepath.SplitList":     filepath.SplitList,
	"path/filepath.ToSlash":       filepath.ToSlash,
	"path/filepath.VolumeName":    filepath.VolumeName,
	"path/filepath.Walk":          filepath.Walk,
}

var typList = []interface{}{
	(*filepath.WalkFunc)(nil),
}
