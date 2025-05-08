package gopbuild

//go:generate go run ../cmd/qexp -outdir ../pkg github.com/qiniu/x/gop/...
//go:generate go run ../cmd/qexp -outdir ../pkg github.com/qiniu/x/errors
//go:generate go run ../cmd/qexp -outdir ../pkg github.com/qiniu/x/stringutil
//go:generate go run ../cmd/qexp -outdir ../pkg github.com/qiniu/x/stringslice

import (
	_ "github.com/goplus/igop/pkg/github.com/qiniu/x/gop"
	_ "github.com/goplus/igop/pkg/github.com/qiniu/x/gop/ng"
	_ "github.com/goplus/igop/pkg/github.com/qiniu/x/gop/osx"
	_ "github.com/goplus/igop/pkg/github.com/qiniu/x/stringslice"

	_ "github.com/goplus/igop/pkg/errors"
	_ "github.com/goplus/igop/pkg/fmt"
	_ "github.com/goplus/igop/pkg/io"
	_ "github.com/goplus/igop/pkg/math"
	_ "github.com/goplus/igop/pkg/math/big"
	_ "github.com/goplus/igop/pkg/os"
	_ "github.com/goplus/igop/pkg/reflect"
	_ "github.com/goplus/igop/pkg/strconv"
	_ "github.com/goplus/igop/pkg/strings"

	_ "github.com/goplus/igop/pkg/github.com/qiniu/x/errors"
	_ "github.com/goplus/igop/pkg/github.com/qiniu/x/stringutil"
)
