//go:build ignore
// +build ignore

/*
 * Copyright (c) 2022 The GoPlus Authors (goplus.org). All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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
	buf.WriteString(`if typ.PkgPath() == "" {
		switch typ.Kind() {
`)
	var div_case1 string
	var div_case2 string
	if op == "/" {
		div_case1 = get_func_case1_div()
		div_case2 = get_func_case2_div()
	}
	for _, kind := range kinds {
		r := strings.NewReplacer(
			"Int", strings.Title(kind),
			"int", kind,
			"+", op)
		if op == "/" && strings.Contains(kind, "int") {
			r.WriteString(buf, div_case1)
		} else {
			r.WriteString(buf, func_case1)
		}
	}
	buf.WriteString(`}
	} else {
	t := xtype.TypeOfType(typ)
	switch typ.Kind() {
`)
	for _, kind := range kinds {
		r := strings.NewReplacer(
			"Int", strings.Title(kind),
			"int", kind,
			"+", op)
		if op == "/" && strings.Contains(kind, "int") {
			r.WriteString(buf, div_case2)
		} else {
			r.WriteString(buf, func_case2)
		}
	}
	buf.WriteString(`}
	}
	panic("unreachable")
}`)
}

func makeFuncCmp(buf *bytes.Buffer, fnname string, op string, kinds []string) {
	buf.WriteString(strings.Replace(func_head_cmp, "$NAME", fnname, 1))
	buf.WriteString(`if typ.PkgPath() == "" {
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
			"int", kind,
			"<", op)
		r.WriteString(buf, func_case2_cmp)
	}
	buf.WriteString(`}
	}
	panic("unreachable")
}`)
}

var pkg_head = `/*
 * Copyright (c) 2022 The GoPlus Authors (goplus.org). All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package igop

import (
	"reflect"

	"github.com/visualfc/xtype"
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

func get_func_case1_div() string {
	return strings.Replace(func_case1,
		"v := vx.(int)+vy.(int)",
		`x := vx.(int)
	y := vy.(int)
	if y == 0 {
		return func(fr *frame) { fr.setReg(ir, x/y) }
	}
	v := x/y`, 1)
}

var func_case1 = `case reflect.Int:
if kx == kindConst && ky == kindConst {
	v := vx.(int)+vy.(int)
	return func(fr *frame) { fr.setReg(ir, v) }
} else if kx == kindConst {
	x := vx.(int)
	return func(fr *frame) { fr.setReg(ir, x+fr.reg(iy).(int)) }
} else if ky == kindConst {
	y := vy.(int)
	return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)+y) }
} else {
	return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)+fr.reg(iy).(int)) }
}
`

func get_func_case2_div() string {
	return strings.Replace(func_case2,
		"v := xtype.Make(t,xtype.Int(vx)+xtype.Int(vy))",
		`x := xtype.Int(vx)
		y := xtype.Int(vy)
		if y == 0 {
		return func(fr *frame) { fr.setReg(ir, xtype.Make(t,x/y)) }
	}
	v := xtype.Make(t,x/y)`, 1)
}

var func_case2 = `case reflect.Int:
if kx == kindConst && ky == kindConst {
	v := xtype.Make(t,xtype.Int(vx)+xtype.Int(vy))
	return func(fr *frame) { fr.setReg(ir, v) }
} else if kx == kindConst {
	x := xtype.Int(vx)
	return func(fr *frame) { fr.setReg(ir, xtype.Make(t,x+fr.int(iy))) }
} else if ky == kindConst {
	y := xtype.Int(vy)
	return func(fr *frame) { fr.setReg(ir, xtype.Make(t,fr.int(ix)+y)) }
} else {
	return func(fr *frame) { fr.setReg(ir, xtype.Make(t,fr.int(ix)+fr.int(iy))) }
}
`

var func_case2_cmp = `case reflect.Int:
if kx == kindConst && ky == kindConst {
	v := xtype.Int(vx)<xtype.Int(vy)
	return func(fr *frame) { fr.setReg(ir, v) }
} else if kx == kindConst {
	x := xtype.Int(vx)
	return func(fr *frame) { fr.setReg(ir, x<fr.int(iy)) }
} else if ky == kindConst {
	y := xtype.Int(vy)
	return func(fr *frame) { fr.setReg(ir, fr.int(ix)<y) }
} else {
	return func(fr *frame) { fr.setReg(ir, fr.int(ix)<fr.int(iy)) }
}
`
