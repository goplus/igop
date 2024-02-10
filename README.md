# iGo+ The Go/Go+ Interpreter

[![Go1.18](https://github.com/goplus/igop/workflows/Go1.18/badge.svg)](https://github.com/goplus/igop/actions/workflows/go118.yml)
[![Go1.19](https://github.com/goplus/igop/workflows/Go1.19/badge.svg)](https://github.com/goplus/igop/actions/workflows/go119.yml)
[![Go1.20](https://github.com/goplus/igop/workflows/Go1.20/badge.svg)](https://github.com/goplus/igop/actions/workflows/go120.yml)
[![Go1.21](https://github.com/goplus/igop/workflows/Go1.21/badge.svg)](https://github.com/goplus/igop/actions/workflows/go121.yml)
[![Go1.22](https://github.com/goplus/igop/workflows/Go1.22/badge.svg)](https://github.com/goplus/igop/actions/workflows/go122.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/goplus/igop.svg)](https://pkg.go.dev/github.com/goplus/igop)


### Go Version

- Go1.18 ~ Go1.22
- macOS Linux Windows  WebAssembly GopherJS and more.

### ABI

support ABI0 and ABIInternal

- ABI0 stack-based ABI
- ABIInternal [register-based Go calling convention proposal](https://golang.org/design/40724-register-calling)

    - Go1.17: amd64
    - Go1.18: amd64 arm64 ppc64/ppc64le
    - Go1.19: amd64 arm64 ppc64/ppc64le riscv64
    - Go1.20: amd64 arm64 ppc64/ppc64le riscv64
    - Go1.21: amd64 arm64 ppc64/ppc64le riscv64
    - Go1.22: amd64 arm64 ppc64/ppc64le riscv64

### Generics

- support generics (Go1.18 ~ Go1.22)
- support [Go1.20 nested type-parameterized declarations](https://github.com/golang/go/blob/master/test/typeparam/nested.go) on Go1.18/Go1.19 (Experimental)

### runtime.GC

`igop.Mode ExperimentalSupportGC`

experimental support runtime.GC and runtime.SetFinalizer

### install igop command line

Go version < 1.17:

```shell
go get -u github.com/goplus/igop/cmd/igop
```

Go version >= 1.17:

```shell
go install github.com/goplus/igop/cmd/igop@latest
```

### igop command

```
igop             # igop repl mode
igop run         # run a Go/Go+ package
igop build       # compile a Go/Go+ package
igop test        # test a package
igop verson      # print version
igop export      # export Go package to igop builtin package
```

### igop run mode
```
Usage: igop run [build flags] [package] [arguments...]
  -exp-gc
    	experimental support runtime.GC
  -mod value
    	module download mode to use: readonly, vendor, or mod.
  -ssa
    	print SSA instruction code
  -ssa-trace
    	trace SSA interpreter code
  -tags value
    	a comma-separated list of build tags to consider satisfied during the build
  -v	print the names of packages as they are compiled.
  -x	print the commands.
```

### igop repl mode

```shell
igop                       # run repl mode, support Go/Go+
igop repl                  # run repl mode, support Go/Go+
igop repl -go              # run repl mode, disable Go+ syntax
```

### igop test unsupport features

- test -fuzz
- test -cover

### igop demo

#### go js playground (gopherjs)

- <https://jsplay.goplus.org/>
- <https://github.com/goplusjs/play>

#### go repl playground (gopherjs/wasm)

- <https://repl.goplus.org/>
- <https://github.com/goplusjs/repl>

#### ispx

<https://github.com/goplus/ispx>

#### run simple Go source demo

```go
package main

import (
	"github.com/goplus/igop"
	_ "github.com/goplus/igop/pkg/fmt"
)

var source = `
package main
import "fmt"
type T struct {}
func (t T) String() string {
	return "Hello, World"
}
func main() {
	fmt.Println(&T{})
}
`

func main() {
	_, err := igop.RunFile("main.go", source, nil, 0)
	if err != nil {
		panic(err)
	}
}
```

#### run simple Go+ source demo

```go
package main

import (
	"github.com/goplus/igop"
	_ "github.com/goplus/igop/gopbuild"
)

var gopSrc string = `
fields := [
	"engineering",
	"STEM education", 
	"and data science",
]

println "The Go+ language for", fields.join(", ")
`

func main() {
	_, err := igop.RunFile("main.gop", gopSrc, nil, 0)
	if err != nil {
		panic(err)
	}
}
```

#### igop more demo

<https://github.com/visualfc/igop_demo>
