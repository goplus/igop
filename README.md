# interp

### example
```
package main

import (
	"fmt"

	"github.com/goplus/interp"
)

var souce = `
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
	err := interp.RunSource(interp.EnableTracing, souce)
	if err != nil {
		panic(err)
	}
}

```