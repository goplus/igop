// export by github.com/goplus/igop/cmd/qexp

//+build go1.16,!go1.17

package quick

import (
	q "testing/quick"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "quick",
		Path: "testing/quick",
		Deps: map[string]string{
			"flag":      "flag",
			"fmt":       "fmt",
			"math":      "math",
			"math/rand": "rand",
			"reflect":   "reflect",
			"strings":   "strings",
			"time":      "time",
		},
		Interfaces: map[string]reflect.Type{
			"Generator": reflect.TypeOf((*q.Generator)(nil)).Elem(),
		},
		NamedTypes: map[string]igop.NamedType{
			"CheckEqualError": {reflect.TypeOf((*q.CheckEqualError)(nil)).Elem(), "", "Error"},
			"CheckError":      {reflect.TypeOf((*q.CheckError)(nil)).Elem(), "", "Error"},
			"Config":          {reflect.TypeOf((*q.Config)(nil)).Elem(), "", "getMaxCount,getRand"},
			"SetupError":      {reflect.TypeOf((*q.SetupError)(nil)).Elem(), "Error", ""},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Check":      reflect.ValueOf(q.Check),
			"CheckEqual": reflect.ValueOf(q.CheckEqual),
			"Value":      reflect.ValueOf(q.Value),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
