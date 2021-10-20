// export by github.com/goplus/gossa/cmd/qexp

//go:build go1.15
// +build go1.15

package sql

import (
	"database/sql"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("database/sql", extMap_go115, typList_go115)
}

var extMap_go115 = map[string]interface{}{
	"(*database/sql.DB).SetConnMaxIdleTime": (*sql.DB).SetConnMaxIdleTime,
	"(*database/sql.Row).Err":               (*sql.Row).Err,
}

var typList_go115 = []interface{}{}
