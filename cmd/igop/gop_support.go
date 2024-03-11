//go:build go1.18
// +build go1.18

package main

import (
	"go/ast"
	"go/constant"
	"go/token"
	"go/types"
	"io"
	"net/http"
	"time"
	_ "unsafe"

	"github.com/goplus/igop"
	_ "github.com/goplus/igop/cmd/igop/pkg/github.com/modern-go/reflect2"
)

//go:generate go run ../qexp -outdir pkg github.com/modern-go/reflect2

type operandMode byte
type builtinId int

type operand struct {
	mode operandMode
	expr ast.Expr
	typ  types.Type
	val  constant.Value
	id   builtinId
}

type positioner interface {
	Pos() token.Pos
}

//go:linkname checker_infer go/types.(*Checker).infer
func checker_infer(check *types.Checker, posn positioner, tparams []*types.TypeParam, targs []types.Type, params *types.Tuple, args []*operand) (result []types.Type)

//go:linkname serveContent net/http.serveContent
func serveContent(w http.ResponseWriter, r *http.Request, name string, modtime time.Time, sizeFunc func() (int64, error), content io.ReadSeeker)

//go:linkname toHTTPError net/http.toHTTPError
func toHTTPError(err error) (msg string, httpStatus int)

func init() {
	// support github.com/goplus/gogen.checker_infer
	igop.RegisterExternal("go/types.(*Checker).infer", checker_infer)
	igop.RegisterExternal("net/http.serveContent", serveContent)
	igop.RegisterExternal("net/http.toHTTPError", toHTTPError)
}
