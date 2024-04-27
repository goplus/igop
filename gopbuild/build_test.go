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

package gopbuild

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/goplus/igop"
	_ "github.com/goplus/igop/pkg/bytes"
)

func gopClTest(t *testing.T, gopcode, expected string) {
	gopClTestEx(t, "main.gop", gopcode, expected)
}

func gopClTestEx(t *testing.T, filename string, gopcode, expected string) {
	ctx := igop.NewContext(0)
	data, err := BuildFile(ctx, filename, gopcode)
	if err != nil {
		t.Fatalf("build gop error: %v", err)
	}
	if string(data) != expected {
		fmt.Println("build gop error:")
		fmt.Println(string(data))
		t.Fail()
	}
}

func TestGop(t *testing.T) {
	gopClTest(t, `
println "Go+"
`, `package main

import "fmt"
//line main.gop:2
func main() {
//line main.gop:2:1
	fmt.Println("Go+")
}
`)
}

func TestGox(t *testing.T) {
	gopClTestEx(t, "Rect.gox", `
println "Go+"
`, `package main

import "fmt"

type Rect struct {
}
//line Rect.gox:2
func (this *Rect) Main() {
//line Rect.gox:2:1
	fmt.Println("Go+")
}
func main() {
}
`)
	gopClTestEx(t, "Rect.gox", `
var (
	Buffer
	v int
)
type Buffer struct {
	buf []byte
}
println "Go+"
`, `package main

import "fmt"

type Buffer struct {
	buf []byte
}
type Rect struct {
	Buffer
	v int
}
//line Rect.gox:9
func (this *Rect) Main() {
//line Rect.gox:9:1
	fmt.Println("Go+")
}
func main() {
}
`)
	gopClTestEx(t, "Rect.gox", `
var (
	*Buffer
	v int
)
type Buffer struct {
	buf []byte
}
println "Go+"
`, `package main

import "fmt"

type Buffer struct {
	buf []byte
}
type Rect struct {
	*Buffer
	v int
}
//line Rect.gox:9
func (this *Rect) Main() {
//line Rect.gox:9:1
	fmt.Println("Go+")
}
func main() {
}
`)
	gopClTestEx(t, "Rect.gox", `
import "bytes"
var (
	*bytes.Buffer
	v int
)
println "Go+"
`, `package main

import (
	"bytes"
	"fmt"
)

type Rect struct {
	*bytes.Buffer
	v int
}
//line Rect.gox:7
func (this *Rect) Main() {
//line Rect.gox:7:1
	fmt.Println("Go+")
}
func main() {
}
`)
	gopClTestEx(t, "Rect.gox", `
import "bytes"
var (
	bytes.Buffer
	v int
)
println "Go+"
`, `package main

import (
	"bytes"
	"fmt"
)

type Rect struct {
	bytes.Buffer
	v int
}
//line Rect.gox:7
func (this *Rect) Main() {
//line Rect.gox:7:1
	fmt.Println("Go+")
}
func main() {
}
`)
}

func TestBig(t *testing.T) {
	gopClTest(t, `
a := 1/2r
println a+1/2r
`, `package main

import (
	"fmt"
	"github.com/goplus/gop/builtin/ng"
	"math/big"
)
//line main.gop:2
func main() {
//line main.gop:2:1
	a := ng.Bigrat_Init__2(big.NewRat(1, 2))
//line main.gop:3:1
	fmt.Println((ng.Bigrat).Gop_Add(a, ng.Bigrat_Init__2(big.NewRat(1, 2))))
}
`)
}

func TestBuiltin(t *testing.T) {
	igop.RegisterCustomBuiltin("typeof", reflect.TypeOf)
	gopClTest(t, `
v := typeof(100)
println(v)
`, `package main

import "fmt"
//line main.gop:2
func main() {
//line main.gop:2:1
	v := typeof(100)
//line main.gop:3:1
	fmt.Println(v)
}
`)
}

func TestIoxLines(t *testing.T) {
	gopClTest(t, `
import "io"

var r io.Reader

for line <- lines(r) {
	println line
}
`, `package main

import (
	"fmt"
	"github.com/goplus/gop/builtin/iox"
	"io"
)

var r io.Reader
//line main.gop:6
func main() {
//line main.gop:6:1
	for _gop_it := iox.Lines(r).Gop_Enum(); ; {
		var _gop_ok bool
		line, _gop_ok := _gop_it.Next()
		if !_gop_ok {
			break
		}
//line main.gop:7:1
		fmt.Println(line)
	}
}
`)
}

func TestErrorWrap(t *testing.T) {
	gopClTest(t, `
import (
    "strconv"
)

func add(x, y string) (int, error) {
    return strconv.Atoi(x)? + strconv.Atoi(y)?, nil
}

func addSafe(x, y string) int {
    return strconv.Atoi(x)?:0 + strconv.Atoi(y)?:0
}

println add("100", "23")!

sum, err := add("10", "abc")
println sum, err

println addSafe("10", "abc")
`, `package main

import (
	"fmt"
	"github.com/qiniu/x/errors"
	"strconv"
)
//line main.gop:6:1
func add(x string, y string) (int, error) {
//line main.gop:7:1
	var _autoGo_1 int
//line main.gop:7:1
	{
//line main.gop:7:1
		var _gop_err error
//line main.gop:7:1
		_autoGo_1, _gop_err = strconv.Atoi(x)
//line main.gop:7:1
		if _gop_err != nil {
//line main.gop:7:1
			_gop_err = errors.NewFrame(_gop_err, "strconv.Atoi(x)", "main.gop", 7, "main.add")
//line main.gop:7:1
			return 0, _gop_err
		}
//line main.gop:7:1
		goto _autoGo_2
	_autoGo_2:
//line main.gop:7:1
	}
//line main.gop:7:1
	var _autoGo_3 int
//line main.gop:7:1
	{
//line main.gop:7:1
		var _gop_err error
//line main.gop:7:1
		_autoGo_3, _gop_err = strconv.Atoi(y)
//line main.gop:7:1
		if _gop_err != nil {
//line main.gop:7:1
			_gop_err = errors.NewFrame(_gop_err, "strconv.Atoi(y)", "main.gop", 7, "main.add")
//line main.gop:7:1
			return 0, _gop_err
		}
//line main.gop:7:1
		goto _autoGo_4
	_autoGo_4:
//line main.gop:7:1
	}
//line main.gop:7:1
	return _autoGo_1 + _autoGo_3, nil
}
//line main.gop:10:1
func addSafe(x string, y string) int {
//line main.gop:11:1
	return func() (_gop_ret int) {
//line main.gop:11:1
		var _gop_err error
//line main.gop:11:1
		_gop_ret, _gop_err = strconv.Atoi(x)
//line main.gop:11:1
		if _gop_err != nil {
//line main.gop:11:1
			return 0
		}
//line main.gop:11:1
		return
	}() + func() (_gop_ret int) {
//line main.gop:11:1
		var _gop_err error
//line main.gop:11:1
		_gop_ret, _gop_err = strconv.Atoi(y)
//line main.gop:11:1
		if _gop_err != nil {
//line main.gop:11:1
			return 0
		}
//line main.gop:11:1
		return
	}()
}
//line main.gop:14
func main() {
//line main.gop:14:1
	fmt.Println(func() (_gop_ret int) {
//line main.gop:14:1
		var _gop_err error
//line main.gop:14:1
		_gop_ret, _gop_err = add("100", "23")
//line main.gop:14:1
		if _gop_err != nil {
//line main.gop:14:1
			_gop_err = errors.NewFrame(_gop_err, "add(\"100\", \"23\")", "main.gop", 14, "main.main")
//line main.gop:14:1
			panic(_gop_err)
		}
//line main.gop:14:1
		return
	}())
//line main.gop:16:1
	sum, err := add("10", "abc")
//line main.gop:17:1
	fmt.Println(sum, err)
//line main.gop:19:1
	fmt.Println(addSafe("10", "abc"))
}
`)
}

func TestGsh(t *testing.T) {
	gopClTestEx(t, "exec.gsh", `
gop "run", "./foo"
exec "gop run ./foo"
exec "FOO=100 gop run ./foo"
exec {"FOO": "101"}, "gop", "run", "./foo"
exec "gop", "run", "./foo"
exec "ls $HOME"
ls "${HOME}"

`, `package main

import "github.com/qiniu/x/gsh"

type exec struct {
	gsh.App
}
//line exec.gsh:2
func (this *exec) MainEntry() {
//line exec.gsh:2:1
	this.Gop_Exec("gop", "run", "./foo")
//line exec.gsh:3:1
	this.Exec__1("gop run ./foo")
//line exec.gsh:4:1
	this.Exec__1("FOO=100 gop run ./foo")
//line exec.gsh:5:1
	this.Exec__0(map[string]string{"FOO": "101"}, "gop", "run", "./foo")
//line exec.gsh:6:1
	this.Exec__2("gop", "run", "./foo")
//line exec.gsh:7:1
	this.Exec__1("ls $HOME")
//line exec.gsh:8:1
	this.Gop_Exec("ls", this.Gop_Env("HOME"))
}
func (this *exec) Main() {
	gsh.Gopt_App_Main(this)
}
func main() {
	new(exec).Main()
}
`)
}
