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
)

type TypeKind struct {
	typ  string
	kind string
}

func main() {
	var buf bytes.Buffer
	buf.WriteString(pkg_head)
	var kinds []*TypeKind
	for _, v := range ints {
		kinds = append(kinds, &TypeKind{v, strings.Title(v)})
	}

	var lbuf bytes.Buffer
	makeFuncSHL(&lbuf, kinds)
	buf.Write(lbuf.Bytes())
	r := strings.NewReplacer("makeBinOpSHL", "makeBinOpSHR", "<<", ">>")
	r.WriteString(&buf, lbuf.String())

	data, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Println("format error", err)
		fmt.Println(buf.String())
		os.Exit(2)
	}
	ioutil.WriteFile("./binop_shift.go", data, 0666)
}

func makeFuncSHL(buf *bytes.Buffer, kinds []*TypeKind) {
	buf.WriteString(func_head_op)
	// check x const && y const
	buf.WriteString(`if kx == kindConst && ky == kindConst {
	t := xtype.TypeOfType(xtyp)
	switch xkind {
`)
	for _, kx := range kinds {
		buf.WriteString(fmt.Sprintf(`	case reflect.%v:
		x := xtype.%v(vx)
		switch ykind {
`, kx.kind, kx.kind))
		for _, ky := range kinds {
			buf.WriteString(fmt.Sprintf(`case reflect.%v:
			v := xtype.Make(t, x << xtype.%v(vy))
			return func(fr *frame) { fr.setReg(ir,v) }
`, ky.kind, ky.kind))
		}
		// end switch ykind
		buf.WriteString("}\n")
	}

	// end switch xkind
	buf.WriteString("}\n")
	// end x const && y const
	buf.WriteString("}\n")

	buf.WriteString(`if xtyp.PkgPath() == "" {
	switch xkind {
`)
	for _, kx := range kinds {
		buf.WriteString(fmt.Sprintf(`	case reflect.%v:
`, kx.kind))
		// if x const
		buf.WriteString(fmt.Sprintf(`	if kx == kindConst {
		x := vx.(%v)
`, kx.typ))
		buf.WriteString(fmt.Sprintf("\t\t\tswitch ykind {\n"))
		for _, ky := range kinds {
			buf.WriteString(fmt.Sprintf(`case reflect.%v:
			return func(fr *frame) { fr.setReg(ir,x<<fr.%v(iy)) }
`, ky.kind, ky.typ))
		}

		buf.WriteString(fmt.Sprintf("\t\t\t}\n"))

		// else if y const
		buf.WriteString("\t\t} else if ky == kindConst {\n")
		buf.WriteString(fmt.Sprintf("\t\t\tswitch ykind {\n"))
		for _, ky := range kinds {
			buf.WriteString(fmt.Sprintf(`case reflect.%v:
			y := xtype.%v(vy)
			return func(fr *frame) { fr.setReg(ir,fr.reg(ix).(%v)<<y) }
`, ky.kind, ky.kind, kx.typ))
		}
		buf.WriteString(fmt.Sprintf("\t\t\t}\n"))

		// else x no const and y no const
		buf.WriteString("\t} else {\n")
		buf.WriteString("switch ykind {\n")

		for _, ky := range kinds {
			buf.WriteString(fmt.Sprintf(`case reflect.%v:
			return func(fr *frame) { fr.setReg(ir,fr.reg(ix).(%v)<<fr.%v(iy)) }
`, ky.kind, kx.typ, ky.typ))
		}
		buf.WriteString("}\n") // end swith ykind
		buf.WriteString("}\n")
	}
	buf.WriteString("\t\t};") // end xkind

	buf.WriteString("} else {") //end xtypIsBasic

	// begin not xtype
	buf.WriteString(`
		t := xtype.TypeOfType(xtyp)
		switch xkind {
	`)
	for _, kx := range kinds {
		buf.WriteString(fmt.Sprintf(`	case reflect.%v:
`, kx.kind))
		// if x const
		buf.WriteString(fmt.Sprintf(`	if kx == kindConst {
			x := xtype.%v(vx)
	`, kx.kind))
		buf.WriteString(fmt.Sprintf("\t\t\tswitch ykind {\n"))
		for _, ky := range kinds {
			buf.WriteString(fmt.Sprintf(`case reflect.%v:
				return func(fr *frame) { fr.setReg(ir,xtype.Make(t,x<<fr.%v(iy))) }
`, ky.kind, ky.typ))
		}

		buf.WriteString(fmt.Sprintf("\t\t\t}\n"))

		// else if y const
		buf.WriteString("\t\t} else if ky == kindConst {\n")
		buf.WriteString(fmt.Sprintf("\t\t\tswitch ykind {\n"))
		for _, ky := range kinds {
			buf.WriteString(fmt.Sprintf(`case reflect.%v:
				y := xtype.%v(vy)
				return func(fr *frame) { fr.setReg(ir,xtype.Make(t,fr.%v(ix)<<y)) }
`, ky.kind, ky.kind, kx.typ))
		}
		buf.WriteString(fmt.Sprintf("\t\t\t}\n"))

		// else x no const and y no const
		buf.WriteString("\t} else {\n")
		buf.WriteString("switch ykind {\n")

		for _, ky := range kinds {
			buf.WriteString(fmt.Sprintf(`case reflect.%v:
				return func(fr *frame) { fr.setReg(ir,xtype.Make(t,fr.%v(ix)<<fr.%v(iy))) }
`, ky.kind, kx.typ, ky.typ))
		}
		buf.WriteString("}\n") // end swith ykind
		buf.WriteString("}\n")
	}
	buf.WriteString("\t\t};") // end xkind

	buf.WriteString("};") // end not xtype

	buf.WriteString("\tpanic(\"unreachable\")\n}\n")
}

var pkg_head = `package igop

import (
	"reflect"

	"github.com/visualfc/xtype"
	"golang.org/x/tools/go/ssa"
)
`

var func_head_op = `
func makeBinOpSHL(pfn *function, instr *ssa.BinOp) func(fr *frame) {
	ir := pfn.regIndex(instr)
	ix, kx, vx := pfn.regIndex3(instr.X)
	iy, ky, vy := pfn.regIndex3(instr.Y)
	xtyp := pfn.Interp.preToType(instr.X.Type())
	ytyp := pfn.Interp.preToType(instr.Y.Type())
	xkind := xtyp.Kind()
	ykind := ytyp.Kind()
`
