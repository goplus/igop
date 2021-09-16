# interp - Golang SSA interpreter

[![Go1.14](https://github.com/goplus/interp/workflows/Go1.14/badge.svg)](https://github.com/goplus/interp/actions?query=workflow%3AGo1.14)
[![Go1.15](https://github.com/goplus/interp/workflows/Go1.15/badge.svg)](https://github.com/goplus/interp/actions?query=workflow%3AGo1.15)
[![Go1.16](https://github.com/goplus/interp/workflows/Go1.16/badge.svg)](https://github.com/goplus/interp/actions?query=workflow%3AGo1.16)

**Go1.17**

set env

`GOEXPERIMENT=noregabi`

### igop command line
```
go get -u github.com/goplus/interp/cmd/igop
```

Commands
```
igop run         # interpret package [gop|go]
igop test        # test package [gop|go]
```

### interp package
**run gop source**
```
package main

import (
	"github.com/goplus/interp"
	_ "github.com/goplus/interp/pkg"
)

var source = `
println("Hello, Go+")
`

func main() {
	err := interp.RunFile(0, "main.gop", source, nil)
	if err != nil {
		panic(err)
	}
}
```
**run go source**
```
package main

import (
	"github.com/goplus/interp"
	_ "github.com/goplus/interp/pkg"
)

var source = `
package main

import "fmt"

func main() {
	fmt.Println("hello")
}
`

func main() {
	err := interp.RunFile(0,"main.go", source, nil)
	if err != nil {
		panic(err)
	}
}

```
