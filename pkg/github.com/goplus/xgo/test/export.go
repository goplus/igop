// export by github.com/goplus/ixgo/cmd/qexp

package test

import (
	q "github.com/goplus/xgo/test"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "test",
		Path: "github.com/goplus/xgo/test",
		Deps: map[string]string{
			"os":      "os",
			"testing": "testing",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"App":  reflect.TypeOf((*q.App)(nil)).Elem(),
			"Case": reflect.TypeOf((*q.Case)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Gopt_App_TestMain":  reflect.ValueOf(q.Gopt_App_TestMain),
			"Gopt_Case_TestMain": reflect.ValueOf(q.Gopt_Case_TestMain),
		},
		TypedConsts: map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{
			"GopPackage": {"untyped bool", constant.MakeBool(bool(q.GopPackage))},
		},
		Source: source,
	})
}

var source = "package test\n\nimport (\n\t\"os\"\n\t\"testing\"\n)\n\nconst (\n\tGopPackage = true\n)\n\ntype testingT = testing.T\n\ntype Case struct {\n\tt *testingT\n}\n\nfunc (p *Case) initCase(t *testing.T) {\n\tp.t = t\n}\n\nfunc (p Case) T() *testing.T\t{ return p.t }\n\nfunc (p Case) Run(name string, f func(t *testing.T)) bool {\n\treturn p.t.Run(name, f)\n}\n\nfunc Gopt_Case_TestMain(c interface{ initCase(t *testing.T) }, t *testing.T) {\n\tc.initCase(t)\n\tc.(interface{ Main() }).Main()\n}\n\ntype App struct {\n\tm *testing.M\n}\n\nfunc (p *App) initApp(m *testing.M) {\n\tp.m = m\n}\n\nfunc (p App) M() *testing.M\t{ return p.m }\n\nfunc Gopt_App_TestMain(app interface{ initApp(m *testing.M) }, m *testing.M) {\n\tapp.initApp(m)\n\tif me, ok := app.(interface{ MainEntry() }); ok {\n\t\tme.MainEntry()\n\t}\n\tos.Exit(m.Run())\n}\n"
