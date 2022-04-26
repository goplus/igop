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

	makeFuncOp(&buf, "makeUnOpSUB", "-", "Neg", json(ints, floats, comps))
	makeFuncOp(&buf, "makeUnOpXOR", "^", "Xor", json(ints))

	data, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Println("format error", err)
		os.Exit(2)
	}
	ioutil.WriteFile("./unop.go", data, 0666)
}

func makeFuncOp(buf *bytes.Buffer, fnname string, op string, neg string, kinds []string) {
	buf.WriteString(strings.Replace(func_head_op, "$NAME", fnname, 1))
	buf.WriteString(`
	if typ.PkgPath() == "" {
		switch typ.Kind() {
`)
	for _, kind := range kinds {
		r := strings.NewReplacer(
			"Int", strings.Title(kind),
			"int", kind,
			"-", op,
			"Neg", neg)
		r.WriteString(buf, func_case1)
	}
	buf.WriteString(`}
	} else {
	switch typ.Kind() {
`)
	for _, kind := range kinds {
		r := strings.NewReplacer(
			"Int", strings.Title(kind),
			"-", op,
			"Neg", neg)
		r.WriteString(buf, func_case2)
	}
	buf.WriteString(`}
	}
	panic("unreachable")
}`)
}

var pkg_head = `package gossa

import (
	"reflect"

	"github.com/goplus/gossa/internal/basic"
	"golang.org/x/tools/go/ssa"
)
`

var func_head_op = `
func $NAME(pfn *function, instr *ssa.UnOp) func(fr *frame) {
	ir := pfn.regIndex(instr)
	ix, kx, vx := pfn.regIndex3(instr.X)
	typ := pfn.Interp.preToType(instr.Type())
`

var func_case1 = `case reflect.Int:
if kx == kindConst {
	v := -vx.(int)
	return func(fr *frame) {
		fr.setReg(ir, v)
	}
} else {
	return func(fr *frame) {
		v := -fr.reg(ix).(int)
		fr.setReg(ir, v)
	}
}
`

var func_case2 = `case reflect.Int:
if kx == kindConst {
	v := basic.NegInt(vx)
	return func(fr *frame) {
		fr.setReg(ir, v)
	}
} else {
	return func(fr *frame) {
		v := basic.NegInt(fr.reg(ix))
		fr.setReg(ir, v)
	}
}
`
