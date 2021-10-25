// export by github.com/goplus/gossa/cmd/qexp

//go:build go1.15
// +build go1.15

package driver

import (
	"database/sql/driver"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("database/sql/driver", extMap_go115, typList_go115)
}

var extMap_go115 = map[string]interface{}{
	"(database/sql/driver.Validator).IsValid": (driver.Validator).IsValid,
}

var typList_go115 = []interface{}{
	(*driver.Validator)(nil),
}
