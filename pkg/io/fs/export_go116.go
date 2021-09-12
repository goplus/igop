// export by github.com/goplus/interp/cmd/qexp

//go:build go1.16
// +build go1.16

package fs

import (
	"io/fs"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("io/fs", extMap_go116, typList_go116)
}

var extMap_go116 = map[string]interface{}{
	"(*io/fs.PathError).Error":    (*fs.PathError).Error,
	"(*io/fs.PathError).Timeout":  (*fs.PathError).Timeout,
	"(*io/fs.PathError).Unwrap":   (*fs.PathError).Unwrap,
	"(io/fs.DirEntry).Info":       (fs.DirEntry).Info,
	"(io/fs.DirEntry).IsDir":      (fs.DirEntry).IsDir,
	"(io/fs.DirEntry).Name":       (fs.DirEntry).Name,
	"(io/fs.DirEntry).Type":       (fs.DirEntry).Type,
	"(io/fs.FS).Open":             (fs.FS).Open,
	"(io/fs.File).Close":          (fs.File).Close,
	"(io/fs.File).Read":           (fs.File).Read,
	"(io/fs.File).Stat":           (fs.File).Stat,
	"(io/fs.FileInfo).IsDir":      (fs.FileInfo).IsDir,
	"(io/fs.FileInfo).ModTime":    (fs.FileInfo).ModTime,
	"(io/fs.FileInfo).Mode":       (fs.FileInfo).Mode,
	"(io/fs.FileInfo).Name":       (fs.FileInfo).Name,
	"(io/fs.FileInfo).Size":       (fs.FileInfo).Size,
	"(io/fs.FileInfo).Sys":        (fs.FileInfo).Sys,
	"(io/fs.FileMode).IsDir":      (fs.FileMode).IsDir,
	"(io/fs.FileMode).IsRegular":  (fs.FileMode).IsRegular,
	"(io/fs.FileMode).Perm":       (fs.FileMode).Perm,
	"(io/fs.FileMode).String":     (fs.FileMode).String,
	"(io/fs.FileMode).Type":       (fs.FileMode).Type,
	"(io/fs.GlobFS).Glob":         (fs.GlobFS).Glob,
	"(io/fs.GlobFS).Open":         (fs.GlobFS).Open,
	"(io/fs.ReadDirFS).Open":      (fs.ReadDirFS).Open,
	"(io/fs.ReadDirFS).ReadDir":   (fs.ReadDirFS).ReadDir,
	"(io/fs.ReadDirFile).Close":   (fs.ReadDirFile).Close,
	"(io/fs.ReadDirFile).Read":    (fs.ReadDirFile).Read,
	"(io/fs.ReadDirFile).ReadDir": (fs.ReadDirFile).ReadDir,
	"(io/fs.ReadDirFile).Stat":    (fs.ReadDirFile).Stat,
	"(io/fs.ReadFileFS).Open":     (fs.ReadFileFS).Open,
	"(io/fs.ReadFileFS).ReadFile": (fs.ReadFileFS).ReadFile,
	"(io/fs.StatFS).Open":         (fs.StatFS).Open,
	"(io/fs.StatFS).Stat":         (fs.StatFS).Stat,
	"(io/fs.SubFS).Open":          (fs.SubFS).Open,
	"(io/fs.SubFS).Sub":           (fs.SubFS).Sub,
	"io/fs.ErrClosed":             &fs.ErrClosed,
	"io/fs.ErrExist":              &fs.ErrExist,
	"io/fs.ErrInvalid":            &fs.ErrInvalid,
	"io/fs.ErrNotExist":           &fs.ErrNotExist,
	"io/fs.ErrPermission":         &fs.ErrPermission,
	"io/fs.Glob":                  fs.Glob,
	"io/fs.ReadDir":               fs.ReadDir,
	"io/fs.ReadFile":              fs.ReadFile,
	"io/fs.SkipDir":               &fs.SkipDir,
	"io/fs.Stat":                  fs.Stat,
	"io/fs.Sub":                   fs.Sub,
	"io/fs.ValidPath":             fs.ValidPath,
	"io/fs.WalkDir":               fs.WalkDir,
}

var typList_go116 = []interface{}{
	(*fs.DirEntry)(nil),
	(*fs.FS)(nil),
	(*fs.File)(nil),
	(*fs.FileInfo)(nil),
	(*fs.FileMode)(nil),
	(*fs.GlobFS)(nil),
	(*fs.PathError)(nil),
	(*fs.ReadDirFS)(nil),
	(*fs.ReadDirFile)(nil),
	(*fs.ReadFileFS)(nil),
	(*fs.StatFS)(nil),
	(*fs.SubFS)(nil),
	(*fs.WalkDirFunc)(nil),
}
