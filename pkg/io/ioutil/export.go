// export by github.com/goplus/gossa/cmd/qexp

package ioutil

import (
	"io/ioutil"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("io/ioutil", extMap, typList)
}

var extMap = map[string]interface{}{
	"io/ioutil.Discard":   &ioutil.Discard,
	"io/ioutil.NopCloser": ioutil.NopCloser,
	"io/ioutil.ReadAll":   ioutil.ReadAll,
	"io/ioutil.ReadDir":   ioutil.ReadDir,
	"io/ioutil.ReadFile":  ioutil.ReadFile,
	"io/ioutil.TempDir":   ioutil.TempDir,
	"io/ioutil.TempFile":  ioutil.TempFile,
	"io/ioutil.WriteFile": ioutil.WriteFile,
}

var typList = []interface{}{}
