# gossa - Golang SSA interpreter

[![Go1.14](https://github.com/goplus/gossa/workflows/Go1.14/badge.svg)](https://github.com/goplus/gossa/actions?query=workflow%3AGo1.14)
[![Go1.15](https://github.com/goplus/gossa/workflows/Go1.15/badge.svg)](https://github.com/goplus/gossa/actions?query=workflow%3AGo1.15)
[![Go1.16](https://github.com/goplus/gossa/workflows/Go1.16/badge.svg)](https://github.com/goplus/gossa/actions?query=workflow%3AGo1.16)

**Go1.17**

set env

`GOEXPERIMENT=noregabi`

### igop command line
```
go get -u github.com/goplus/gossa/cmd/gossa
```

Commands
```
gossa run         # interpret package [gop|go]
gossa test        # test package [gop|go]
```

### gossa package
**run gop source**
```
package main

import (
	"github.com/goplus/gossa"
	_ "github.com/goplus/gossa/pkg"
)

var source = `
println("Hello, Go+")
`

func main() {
	err := gossa.RunFile(0, "main.gop", source, nil)
	if err != nil {
		panic(err)
	}
}
```
**run go source**
```
package main

import (
	"github.com/goplus/gossa"
	_ "github.com/goplus/gossa/pkg"
)

var source = `
package main

import "fmt"

func main() {
	fmt.Println("hello")
}
`

func main() {
	err := gossa.RunFile(0,"main.go", source, nil)
	if err != nil {
		panic(err)
	}
}

```
