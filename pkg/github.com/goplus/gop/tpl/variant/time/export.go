// export by github.com/goplus/igop/cmd/qexp

package time

import (
	_ "github.com/goplus/gop/tpl/variant/time"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "time",
		Path: "github.com/goplus/gop/tpl/variant/time",
		Deps: map[string]string{
			"github.com/goplus/gop/tpl/variant": "variant",
			"time":                              "time",
		},
	})
}
