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

	buf.WriteString("if xIsBasic {")
	buf.WriteString(`
	switch xkind {
`)
	for _, kx := range kinds {
		buf.WriteString(fmt.Sprintf(`	case reflect.%v:
`, kx.kind))
		// if x const && y const
		buf.WriteString(fmt.Sprintf(`if kx == kindConst && ky == kindConst {
			x := vx.(%v)
			switch ykind {
`, kx.typ))
		for _, ky := range kinds {
			buf.WriteString(fmt.Sprintf(`case reflect.%v:
			var y %v
			if yIsBasic { y = vy.(%v) } else { y = basic.%v(vy) }
			v := x << y
			return func(fr *frame) {
				fr.setReg(ir,v)
			}
			`, ky.kind, ky.typ, ky.typ, ky.kind))
		}
		buf.WriteString(`		}
`)
		// else if x const
		buf.WriteString(fmt.Sprintf(`	} else if kx == kindConst {
		x := vx.(%v)
`, kx.typ))
		buf.WriteString(fmt.Sprintf("\t\t\tswitch ykind {\n"))
		for _, ky := range kinds {
			buf.WriteString(fmt.Sprintf(`case reflect.%v:
			if yIsBasic {
				return func(fr *frame) {
					y := fr.reg(iy).(%v)
					fr.setReg(ir,x<<y)
				}
			}
			return func(fr *frame) {
				y := basic.%v(fr.reg(iy))
				fr.setReg(ir,x<<y)
			}			
			`, ky.kind, ky.typ, ky.kind))
		}

		buf.WriteString(fmt.Sprintf("\t\t\t}\n"))

		// else if y const
		buf.WriteString("\t\t} else if ky == kindConst {\n")
		buf.WriteString(fmt.Sprintf("\t\t\tswitch ykind {\n"))
		for _, ky := range kinds {
			buf.WriteString(fmt.Sprintf(`case reflect.%v:
			var y %v
			if yIsBasic { y = vy.(%v) } else { y = basic.%v(vy) }
			return func(fr *frame) {
				x := fr.reg(ix).(%v)
				fr.setReg(ir,x<<y)
			}
			`, ky.kind, ky.typ, ky.typ, ky.kind, kx.typ))
		}
		buf.WriteString(fmt.Sprintf("\t\t\t}\n"))

		// else x no const and y no const
		buf.WriteString("\t} else {\n")
		buf.WriteString("switch ykind {\n")

		for _, ky := range kinds {
			buf.WriteString(fmt.Sprintf(`case reflect.%v:
			if yIsBasic {
				return func(fr *frame) {
					x := fr.reg(ix).(%v)
					y := fr.reg(iy).(%v)
					fr.setReg(ir,x<<y)
				}
			} 
			return func(fr *frame) {
				x := fr.reg(ix).(%v)
				y := basic.%v(fr.reg(iy))
				fr.setReg(ir,x<<y)
			}			
			`, ky.kind, kx.typ, ky.typ,
				kx.typ, ky.kind))
		}
		buf.WriteString("}\n") // end swith ykind
		buf.WriteString("}\n")
	}
	buf.WriteString("\t\t};") // end xkind

	buf.WriteString("} else {") //end xtypIsBasic

	// begin not basic
	buf.WriteString(`
		t := basic.TypeOfType(xtyp)
		switch xkind {
	`)
	for _, kx := range kinds {
		buf.WriteString(fmt.Sprintf(`	case reflect.%v:
`, kx.kind))
		// if x const && y const
		buf.WriteString(fmt.Sprintf(`if kx == kindConst && ky == kindConst {
				x := basic.%v(vx)
				switch ykind {
		`, kx.kind))
		for _, ky := range kinds {
			buf.WriteString(fmt.Sprintf(`case reflect.%v:
				var y %v
				if yIsBasic { y = vy.(%v) } else { y = basic.%v(vy) }
				v := basic.Make(t,x << y)
				return func(fr *frame) {
					fr.setReg(ir,v)
				}
				`, ky.kind, ky.typ, ky.typ, ky.kind))
		}
		buf.WriteString("}\n")
		// else if x const
		buf.WriteString(fmt.Sprintf(`	} else if kx == kindConst {
			x := basic.%v(vx)
	`, kx.kind))
		buf.WriteString(fmt.Sprintf("\t\t\tswitch ykind {\n"))
		for _, ky := range kinds {
			buf.WriteString(fmt.Sprintf(`case reflect.%v:
				if yIsBasic {
					return func(fr *frame) {
						y := fr.reg(iy).(%v)
						fr.setReg(ir,basic.Make(t,x<<y))
					}
				}
				return func(fr *frame) {
					y := basic.%v(fr.reg(iy))
					fr.setReg(ir,basic.Make(t,x<<y))
				}
				`, ky.kind, ky.typ, ky.kind))
		}

		buf.WriteString(fmt.Sprintf("\t\t\t}\n"))

		// else if y const
		buf.WriteString("\t\t} else if ky == kindConst {\n")
		buf.WriteString(fmt.Sprintf("\t\t\tswitch ykind {\n"))
		for _, ky := range kinds {
			buf.WriteString(fmt.Sprintf(`case reflect.%v:
				var y %v
				if yIsBasic { y = vy.(%v) } else { y = basic.%v(vy) }
				return func(fr *frame) {
					x := basic.%v(fr.reg(ix))
					fr.setReg(ir,basic.Make(t,x<<y))
				}
				`, ky.kind, ky.typ, ky.typ, ky.kind, kx.kind))
		}
		buf.WriteString(fmt.Sprintf("\t\t\t}\n"))

		// else x no const and y no const
		buf.WriteString("\t} else {\n")
		buf.WriteString("switch ykind {\n")

		for _, ky := range kinds {
			buf.WriteString(fmt.Sprintf(`case reflect.%v:
				if yIsBasic {
					return func(fr *frame) {
						x := basic.%v(fr.reg(ix))
						y := fr.reg(iy).(%v)
						fr.setReg(ir,basic.Make(t,x<<y))
					}
				}
				return func(fr *frame) {
					x := basic.%v(fr.reg(ix))
					y := basic.%v(fr.reg(iy))
					fr.setReg(ir,basic.Make(t,x<<y))
				}
				`, ky.kind, kx.kind, ky.typ,
				kx.kind, ky.kind))
		}
		buf.WriteString("}\n") // end swith ykind
		buf.WriteString("}\n")
	}
	buf.WriteString("\t\t};") // end xkind

	buf.WriteString("};") // end not basic

	buf.WriteString("\tpanic(\"unreachable\")\n}\n")
}

var pkg_head = `package gossa

import (
	"reflect"

	"github.com/goplus/gossa/basic"
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
	xIsBasic := xtyp.PkgPath() == ""
	yIsBasic := ytyp.PkgPath() == ""
`
