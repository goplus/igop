// export by github.com/goplus/gossa/cmd/qexp

//go:build go1.16
// +build go1.16

package flag

import (
	"flag"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("flag", extMap_go116, typList_go116)
}

var extMap_go116 = map[string]interface{}{
	"(*flag.FlagSet).Func": (*flag.FlagSet).Func,
	"flag.Func":            flag.Func,
}

var typList_go116 = []interface{}{}
