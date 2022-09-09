# iGo+ The Go/Go+ Interpreter

[![Go1.14](https://github.com/goplus/igop/workflows/Go1.14/badge.svg)](https://github.com/goplus/igop/actions/workflows/go114.yml)
[![Go1.15](https://github.com/goplus/igop/workflows/Go1.15/badge.svg)](https://github.com/goplus/igop/actions/workflows/go115.yml)
[![Go1.16](https://github.com/goplus/igop/workflows/Go1.16/badge.svg)](https://github.com/goplus/igop/actions/workflows/go116.yml)
[![Go1.17](https://github.com/goplus/igop/workflows/Go1.17/badge.svg)](https://github.com/goplus/igop/actions/workflows/go117.yml)
[![Go1.18](https://github.com/goplus/igop/workflows/Go1.18/badge.svg)](https://github.com/goplus/igop/actions/workflows/go118.yml)
[![Go1.19](https://github.com/goplus/igop/workflows/Go1.19/badge.svg)](https://github.com/goplus/igop/actions/workflows/go119.yml)

### ABI

support ABI0 and ABIInternal

- ABI0 stack-based ABI
- ABIInternal [register-based Go calling convention proposal](https://golang.org/design/40724-register-calling)

	- Go1.17: amd64
	- Go1.18: amd64 arm64 ppc64/ppc64le
	- Go1.19: amd64 arm64 ppc64/ppc64le riscv64

### Generics

support Go1.18 Go1.19 generics

### igop command line

Go version < 1.17:
```
go get -u github.com/goplus/igop/cmd/igop
```

Go version >= 1.17:
```
go install github.com/goplus/igop/cmd/igop@latest
```

Commands
```
igop             # igop repl mode
igop run         # run a Go/Go+ package
igop build       # compile a Go/Go+ package
igop test        # test a package
igop verson      # print version
igop export      # export Go package to igop builtin package
```

### igop repl mode
```
igop                       # run repl mode, support Go/Go+
igop repl                  # run repl mode, support Go/Go+
igop repl -gop=false       # run repl mode, disable Go+ syntax
```

### igop test unsupport features

- test -fuzz
- test -cover

### igop package

**run go source**
```
package main

import (
	"github.com/goplus/igop"
	_ "github.com/goplus/igop/pkg/fmt"
)

var source = `
package main

import "fmt"

func main() {
	fmt.Println("hello")
}
`

func main() {
	_, err := igop.RunFile("main.go", source, nil, 0)
	if err != nil {
		panic(err)
	}
}

```

**run gop source**
```
package main

import (
	"github.com/goplus/igop"
	_ "github.com/goplus/igop/gopbuild"
	_ "github.com/goplus/igop/pkg/fmt"
)

var source = `
println "Hello, Go+"
`

func main() {
	_, err := igop.RunFile("main.gop", source, nil, 0)
	if err != nil {
		panic(err)
	}
}
```
