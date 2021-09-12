// export by github.com/goplus/interp/cmd/qexp

//go:build go1.16
// +build go1.16

package flag

import (
	"flag"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("flag", extMap_go116, typList_go116)
}

var extMap_go116 = map[string]interface{}{
	"(*flag.FlagSet).Func": (*flag.FlagSet).Func,
	"flag.Func":            flag.Func,
}

var typList_go116 = []interface{}{}
