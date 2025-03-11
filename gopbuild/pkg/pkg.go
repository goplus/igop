package pkg

//go:generate go run ../../cmd/qexp -outdir ../../pkg github.com/goplus/gop/builtin/...
//go:generate go run ../../cmd/qexp -outdir ../../pkg github.com/qiniu/x/stringutil
//go:generate go run ../../cmd/qexp -outdir ../../pkg github.com/qiniu/x/errors
//go:generate go run ../../cmd/qexp -outdir ../../pkg github.com/qiniu/x/gsh
//go:generate go run ../../cmd/qexp -outdir ../../pkg github.com/goplus/gop/ast
//go:generate go run ../../cmd/qexp -outdir ../../pkg github.com/goplus/gop/parser
//go:generate go run ../../cmd/qexp -outdir ../../pkg github.com/goplus/gop/printer
//go:generate go run ../../cmd/qexp -outdir ../../pkg github.com/goplus/gop/scanner
//go:generate go run ../../cmd/qexp -outdir ../../pkg github.com/goplus/gop/token
//go:generate go run ../../cmd/qexp -outdir ../../pkg github.com/goplus/gop/tpl/...

import (
	_ "github.com/goplus/igop/pkg/github.com/goplus/gop/ast"
	_ "github.com/goplus/igop/pkg/github.com/goplus/gop/builtin"
	_ "github.com/goplus/igop/pkg/github.com/goplus/gop/builtin/iox"
	_ "github.com/goplus/igop/pkg/github.com/goplus/gop/builtin/ng"
	_ "github.com/goplus/igop/pkg/github.com/goplus/gop/builtin/stringslice"
	_ "github.com/goplus/igop/pkg/github.com/goplus/gop/parser"
	_ "github.com/goplus/igop/pkg/github.com/goplus/gop/printer"
	_ "github.com/goplus/igop/pkg/github.com/goplus/gop/scanner"
	_ "github.com/goplus/igop/pkg/github.com/goplus/gop/token"
	_ "github.com/goplus/igop/pkg/github.com/goplus/gop/tpl"
	_ "github.com/goplus/igop/pkg/github.com/goplus/gop/tpl/ast"
	_ "github.com/goplus/igop/pkg/github.com/goplus/gop/tpl/cl"
	_ "github.com/goplus/igop/pkg/github.com/goplus/gop/tpl/encoding/csv"
	_ "github.com/goplus/igop/pkg/github.com/goplus/gop/tpl/encoding/json"
	_ "github.com/goplus/igop/pkg/github.com/goplus/gop/tpl/encoding/regexp"
	_ "github.com/goplus/igop/pkg/github.com/goplus/gop/tpl/encoding/regexposix"
	_ "github.com/goplus/igop/pkg/github.com/goplus/gop/tpl/encoding/xml"
	_ "github.com/goplus/igop/pkg/github.com/goplus/gop/tpl/matcher"
	_ "github.com/goplus/igop/pkg/github.com/goplus/gop/tpl/parser"
	_ "github.com/goplus/igop/pkg/github.com/goplus/gop/tpl/token"
	_ "github.com/goplus/igop/pkg/github.com/goplus/gop/tpl/types"
	_ "github.com/goplus/igop/pkg/github.com/qiniu/x/errors"
	_ "github.com/goplus/igop/pkg/github.com/qiniu/x/gsh"
	_ "github.com/goplus/igop/pkg/github.com/qiniu/x/stringutil"

	_ "github.com/goplus/igop/pkg/errors"
	_ "github.com/goplus/igop/pkg/fmt"
	_ "github.com/goplus/igop/pkg/io"
	_ "github.com/goplus/igop/pkg/math"
	_ "github.com/goplus/igop/pkg/math/big"
	_ "github.com/goplus/igop/pkg/os"
	_ "github.com/goplus/igop/pkg/os/exec"
	_ "github.com/goplus/igop/pkg/reflect"
	_ "github.com/goplus/igop/pkg/strconv"
	_ "github.com/goplus/igop/pkg/strings"
)
