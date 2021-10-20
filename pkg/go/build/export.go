// export by github.com/goplus/gossa/cmd/qexp

package build

import (
	"go/build"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("go/build", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*go/build.Context).Import":             (*build.Context).Import,
	"(*go/build.Context).ImportDir":          (*build.Context).ImportDir,
	"(*go/build.Context).MatchFile":          (*build.Context).MatchFile,
	"(*go/build.Context).SrcDirs":            (*build.Context).SrcDirs,
	"(*go/build.MultiplePackageError).Error": (*build.MultiplePackageError).Error,
	"(*go/build.NoGoError).Error":            (*build.NoGoError).Error,
	"(*go/build.Package).IsCommand":          (*build.Package).IsCommand,
	"go/build.ArchChar":                      build.ArchChar,
	"go/build.Default":                       &build.Default,
	"go/build.Import":                        build.Import,
	"go/build.ImportDir":                     build.ImportDir,
	"go/build.IsLocalImport":                 build.IsLocalImport,
	"go/build.ToolDir":                       &build.ToolDir,
}

var typList = []interface{}{
	(*build.Context)(nil),
	(*build.ImportMode)(nil),
	(*build.MultiplePackageError)(nil),
	(*build.NoGoError)(nil),
	(*build.Package)(nil),
}
