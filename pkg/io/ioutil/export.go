// export by github.com/goplus/interp/cmd/qexp

package ioutil

import (
	"io/ioutil"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("io/ioutil", extMap, typList)
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
