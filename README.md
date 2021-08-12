# interp - Golang SSA interpreter

[![Go1.14](https://github.com/goplus/interp/workflows/Go1.14/badge.svg)](https://github.com/goplus/interp/actions?query=workflow%3AGo1.14)
[![Go1.15](https://github.com/goplus/interp/workflows/Go1.15/badge.svg)](https://github.com/goplus/interp/actions?query=workflow%3AGo1.15)
[![Go1.16](https://github.com/goplus/interp/workflows/Go1.16/badge.svg)](https://github.com/goplus/interp/actions?query=workflow%3AGo1.16)


### igop command line
```
go get -u github.com/goplus/interp/cmd/igop
```

Commands
```
igop run         # interpret package
igop test        # test package
```

### interp package
```
package main

import (
	"fmt"

	"github.com/goplus/interp"
)

var source = `
package main

import "fmt"

func main() {
	fmt.Println("hello")
}
`

func init() {
	interp.RegisterExternal("fmt.init", func() {})
	interp.RegisterExternal("fmt.Println", fmt.Println)
}

func main() {
	err := interp.RunSource(interp.EnableTracing, source, nil)
	if err != nil {
		panic(err)
	}
}

```