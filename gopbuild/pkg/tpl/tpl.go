package tpl

import (
	"github.com/goplus/igop"
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
	_ "github.com/goplus/igop/pkg/github.com/goplus/gop/tpl/scanner"
	_ "github.com/goplus/igop/pkg/github.com/goplus/gop/tpl/token"
	_ "github.com/goplus/igop/pkg/github.com/goplus/gop/tpl/types"
	_ "github.com/goplus/igop/pkg/github.com/goplus/gop/tpl/variant"
	_ "github.com/goplus/igop/pkg/github.com/goplus/gop/tpl/variant/builtin"
	_ "github.com/goplus/igop/pkg/github.com/goplus/gop/tpl/variant/delay"
	_ "github.com/goplus/igop/pkg/github.com/goplus/gop/tpl/variant/math"
	_ "github.com/goplus/igop/pkg/github.com/goplus/gop/tpl/variant/time"
)

var tpl_patch = `package tpl

// ListOp converts the matching result of (R % ",") to a flat list.
// R % "," means R *("," R)
func ListOp[T any](in []any, fn func(v any) T) []T {
	next := in[1].([]any)
	ret := make([]T, len(next)+1)
	ret[0] = fn(in[0])
	for i, v := range next {
		ret[i+1] = fn(v.([]any)[1])
	}
	return ret
}
`

var variant_patch = `
package variant

import . "github.com/goplus/gop/tpl/variant"

// ListOp converts the matching result of (R % ",") to a flat list.
// R % "," means R *("," R)
func ListOp[T any](in []any, fn func(v any) T) []T {
	next := in[1].([]any)
	ret := make([]T, len(next)+1)
	ret[0] = fn(Eval(in[0]))
	for i, v := range next {
		ret[i+1] = fn(Eval(v.([]any)[1]))
	}
	return ret
}
`

var delay_patch = `package delay

import "github.com/goplus/gop/tpl/variant"

// ListOp delays to convert the matching result of (R % ",") to a flat list.
// R % "," means R *("," R)
func ListOp[T any](in []any, op func(v any) T, fn func(flat []T)) any {
	return func() any {
		fn(variant.ListOp(in, op))
		return nil
	}
}`

func init() {
	igop.RegisterPatch("github.com/goplus/gop/tpl@patch", tpl_patch)
	igop.RegisterPatch("github.com/goplus/gop/tpl@patch.gop", tpl_patch)
	igop.RegisterPatch("github.com/goplus/gop/tpl/variant@patch", variant_patch)
	igop.RegisterPatch("github.com/goplus/gop/tpl/variant@patch.gop", variant_patch)
	igop.RegisterPatch("github.com/goplus/gop/tpl/variant/delay@patch", delay_patch)
	igop.RegisterPatch("github.com/goplus/gop/tpl/variant/delay@patch.gop", delay_patch)
}
