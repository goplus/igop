//go:build ignore
// +build ignore

package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"strings"
)

var (
	ints = []string{"int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64", "uintptr"}
	floats = []string{"float32", "float64"}
	comps  = []string{"complex64", "complex128"}
	str    = []string{"string"}
)

func json(kinds ...[]string) (r []string) {
	for _, kind := range kinds {
		r = append(r, kind...)
	}
	return
}

func main() {
	var buf bytes.Buffer
	buf.WriteString(pkg_head)

	makeFuncOp(&buf, "makeBinOpADD", "+", json(ints, floats, comps, str))
	makeFuncOp(&buf, "makeBinOpSUB", "-", json(ints, floats, comps))
	makeFuncOp(&buf, "makeBinOpMUL", "*", json(ints, floats, comps))
	makeFuncOp(&buf, "makeBinOpQUO", "/", json(ints, floats, comps))

	makeFuncOp(&buf, "makeBinOpREM", "%", json(ints))
	makeFuncOp(&buf, "makeBinOpAND", "&", json(ints))
	makeFuncOp(&buf, "makeBinOpOR", "|", json(ints))
	makeFuncOp(&buf, "makeBinOpXOR", "^", json(ints))
	makeFuncOp(&buf, "makeBinOpANDNOT", "&^", json(ints))

	makeFuncCmp(&buf, "makeBinOpLSS", "<", json(ints, floats, str))
	makeFuncCmp(&buf, "makeBinOpLEQ", "<=", json(ints, floats, str))
	makeFuncCmp(&buf, "makeBinOpGTR", ">", json(ints, floats, str))
	makeFuncCmp(&buf, "makeBinOpGEQ", ">=", json(ints, floats, str))

	data, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Println("format error", err)
		os.Exit(2)
	}
	ioutil.WriteFile("./binop.go", data, 0666)
}

func makeFuncOp(buf *bytes.Buffer, fnname string, op string, kinds []string) {
	buf.WriteString(strings.Replace(func_head_op, "$NAME", fnname, 1))
	buf.WriteString(`
	if typ.PkgPath() == "" {
		switch typ.Kind() {
`)
	for _, kind := range kinds {
		r := strings.NewReplacer(
			"Int", strings.Title(kind),
			"int", kind,
			"+", op)
		r.WriteString(buf, func_case1)
	}
	buf.WriteString(`}
	} else {
	t := basic.TypeOfType(typ)
	switch typ.Kind() {
`)
	for _, kind := range kinds {
		r := strings.NewReplacer(
			"Int", strings.Title(kind),
			"+", op)
		r.WriteString(buf, func_case2)
	}
	buf.WriteString(`}
	}
	panic("unreachable")
}`)
}

func makeFuncCmp(buf *bytes.Buffer, fnname string, op string, kinds []string) {
	buf.WriteString(strings.Replace(func_head_cmp, "$NAME", fnname, 1))
	buf.WriteString(`
	if typ.PkgPath() == "" {
		switch typ.Kind() {
`)
	for _, kind := range kinds {
		r := strings.NewReplacer(
			"Int", strings.Title(kind),
			"int", kind,
			"+", op)
		r.WriteString(buf, func_case1)
	}
	buf.WriteString(`}
	} else {
	switch typ.Kind() {
`)
	for _, kind := range kinds {
		r := strings.NewReplacer(
			"Int", strings.Title(kind),
			"<", op)
		r.WriteString(buf, func_case2_cmp)
	}
	buf.WriteString(`}
	}
	panic("unreachable")
}`)
}

var pkg_head = `package gossa

import (
	"reflect"

	"github.com/goplus/gossa/basic"
	"golang.org/x/tools/go/ssa"
)
`

var func_head_op = `
func $NAME(pfn *function, instr *ssa.BinOp) func(fr *frame) {
	ir := pfn.regIndex(instr)
	ix, kx, vx := pfn.regIndex3(instr.X)
	iy, ky, vy := pfn.regIndex3(instr.Y)
	typ := pfn.Interp.preToType(instr.Type())
`

var func_head_cmp = `
func $NAME(pfn *function, instr *ssa.BinOp) func(fr *frame) {
	ir := pfn.regIndex(instr)
	ix, kx, vx := pfn.regIndex3(instr.X)
	iy, ky, vy := pfn.regIndex3(instr.Y)
	typ := pfn.Interp.preToType(instr.X.Type())
`

var func_case1 = `case reflect.Int:
if kx == kindConst {
	x := vx.(int)
	return func(fr *frame) {
		y := fr.reg(iy).(int)
		fr.setReg(ir, x+y)
	}
} else if ky == kindConst {
	y := vy.(int)
	return func(fr *frame) {
		x := fr.reg(ix).(int)
		fr.setReg(ir, x+y)
	}
} else {
	return func(fr *frame) {
		x := fr.reg(ix).(int)
		y := fr.reg(iy).(int)
		fr.setReg(ir, x+y)
	}
}
`

var func_case2 = `case reflect.Int:
if kx == kindConst {
	x := basic.Int(vx)
	return func(fr *frame) {
		y := basic.Int(fr.reg(iy))
		fr.setReg(ir, basic.Make(t,x+y))
	}
} else if ky == kindConst {
	y := basic.Int(vy)
	return func(fr *frame) {
		x := basic.Int(fr.reg(ix))
		fr.setReg(ir, basic.Make(t,x+y))
	}
} else {
	return func(fr *frame) {
		x := basic.Int(fr.reg(ix))
		y := basic.Int(fr.reg(iy))
		fr.setReg(ir, basic.Make(t,x+y))
	}
}
`

var func_case2_cmp = `case reflect.Int:
if kx == kindConst {
	x := basic.Int(vx)
	return func(fr *frame) {
		y := basic.Int(fr.reg(iy))
		fr.setReg(ir, x<y)
	}
} else if ky == kindConst {
	y := basic.Int(vy)
	return func(fr *frame) {
		x := basic.Int(fr.reg(ix))
		fr.setReg(ir, x<y)
	}
} else {
	return func(fr *frame) {
		x := basic.Int(fr.reg(ix))
		y := basic.Int(fr.reg(iy))
		fr.setReg(ir, x<y)
	}
}
`
