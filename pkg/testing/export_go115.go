// export by github.com/goplus/gossa/cmd/qexp

//go:build go1.15
// +build go1.15

package testing

import (
	"testing"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("testing", extMap_go115, typList_go115)
}

var extMap_go115 = map[string]interface{}{
	"(*testing.B).TempDir":  (*testing.B).TempDir,
	"(*testing.T).Deadline": (*testing.T).Deadline,
	"(*testing.T).TempDir":  (*testing.T).TempDir,
	"(testing.TB).TempDir":  (testing.TB).TempDir,
}

var typList_go115 = []interface{}{}
