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

	for _, typ := range ints {
		makeFuncOp(&buf, "cvt"+strings.Title(typ), typ)
	}
	for _, typ := range floats {
		makeFuncOp(&buf, "cvt"+strings.Title(typ), typ)
	}
	for _, typ := range comps {
		makeFuncOp2(&buf, "cvt"+strings.Title(typ), typ)
	}

	data, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Println("format error", err)
		os.Exit(2)
	}
	ioutil.WriteFile("./opcvt.go", data, 0666)
}

func makeFuncOp(buf *bytes.Buffer, fnname string, typ string) {
	r := strings.NewReplacer("cvtInt", fnname, "type T = int", "type T = "+typ)
	r.WriteString(buf, cvt_func)
}

func makeFuncOp2(buf *bytes.Buffer, fnname string, typ string) {
	r := strings.NewReplacer("cvtComplex64", fnname, "type T = complex64", "type T = "+typ)
	r.WriteString(buf, cvt_func2)
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
)
`

var cvt_func = `func cvtInt(ir, ix register, xkind reflect.Kind, xtyp reflect.Type, typ reflect.Type) func(fr *frame) {
	type T = int
	t := xtype.TypeOfType(typ)
	isBasic := typ.PkgPath() == ""
	if xtyp.PkgPath() == "" {
		return func(fr *frame) {
			var v T
			switch xkind {
			case reflect.Int:
				v = T(fr.reg(ix).(int))
			case reflect.Int8:
				v = T(fr.reg(ix).(int8))
			case reflect.Int16:
				v = T(fr.reg(ix).(int16))
			case reflect.Int32:
				v = T(fr.reg(ix).(int32))
			case reflect.Int64:
				v = T(fr.reg(ix).(int64))
			case reflect.Uint:
				v = T(fr.reg(ix).(uint))
			case reflect.Uint8:
				v = T(fr.reg(ix).(uint8))
			case reflect.Uint16:
				v = T(fr.reg(ix).(uint16))
			case reflect.Uint32:
				v = T(fr.reg(ix).(uint32))
			case reflect.Uint64:
				v = T(fr.reg(ix).(uint64))
			case reflect.Uintptr:
				v = T(fr.reg(ix).(uintptr))
			case reflect.Float32:
				v = T(fr.reg(ix).(float32))
			case reflect.Float64:
				v = T(fr.reg(ix).(float64))
			}
			if isBasic {
				fr.setReg(ir, v)
			} else {
				fr.setReg(ir, xtype.Make(t, v))
			}
		}
	} else {
		return func(fr *frame) {
			var v T
			switch xkind {
			case reflect.Int:
				v = T(fr.int(ix))
			case reflect.Int8:
				v = T(fr.int8(ix))
			case reflect.Int16:
				v = T(fr.int16(ix))
			case reflect.Int32:
				v = T(fr.int32(ix))
			case reflect.Int64:
				v = T(fr.int64(ix))
			case reflect.Uint:
				v = T(fr.uint(ix))
			case reflect.Uint8:
				v = T(fr.uint8(ix))
			case reflect.Uint16:
				v = T(fr.uint16(ix))
			case reflect.Uint32:
				v = T(fr.uint32(ix))
			case reflect.Uint64:
				v = T(fr.uint64(ix))
			case reflect.Uintptr:
				v = T(fr.uintptr(ix))
			case reflect.Float32:
				v = T(fr.float32(ix))
			case reflect.Float64:
				v = T(fr.float64(ix))
			}
			if isBasic {
				fr.setReg(ir, v)
			} else {
				fr.setReg(ir, xtype.Make(t, v))
			}
		}
	}
}
`

var cvt_func2 = `func cvtComplex64(ir, ix register, xkind reflect.Kind, xtyp reflect.Type, typ reflect.Type) func(fr *frame) {
	type T = complex64
	t := xtype.TypeOfType(typ)
	isBasic := typ.PkgPath() == ""
	if xtyp.PkgPath() == "" {
		return func(fr *frame) {
			var v T
			switch xkind {
			case reflect.Complex64:
				v = T(fr.reg(ix).(complex64))
			case reflect.Complex128:
				v = T(fr.reg(ix).(complex128))
			}
			if isBasic {
				fr.setReg(ir, v)
			} else {
				fr.setReg(ir, xtype.Make(t, v))
			}
		}
	} else {
		return func(fr *frame) {
			var v T
			switch xkind {
			case reflect.Complex64:
				v = T(fr.complex64(ix))
			case reflect.Complex128:
				v = T(fr.complex128(ix))
			}
			if isBasic {
				fr.setReg(ir, v)
			} else {
				fr.setReg(ir, xtype.Make(t, v))
			}
		}
	}
}
`
