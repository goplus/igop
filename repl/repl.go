package repl

import (
	"bytes"
	"errors"
	"fmt"
	"go/token"
	"go/types"
	"io"
	"os/exec"
	"reflect"
	"regexp"
	"sort"
	"strings"

	"golang.org/x/tools/go/ssa"

	"github.com/goplus/igop"
)

const (
	// ContinuePrompt - the current code statement is not completed.
	ContinuePrompt string = "... "
	// NormalPrompt - start of a code statement.
	NormalPrompt string = ">>> "
)

type UI interface {
	SetPrompt(prompt string)
	Printf(format string, a ...interface{})
}

type REPL struct {
	*igop.Repl
	term UI
	more string
}

func NewREPL(mode igop.Mode) *REPL {
	r := &REPL{}
	igop.RegisterCustomBuiltin("__igop_repl_info__", func(v interface{}) {
		r.Printf("%v %v\n", v, reflect.TypeOf(v))
	})
	ctx := igop.NewContext(mode)
	r.Repl = igop.NewRepl(ctx)
	return r
}

func (r *REPL) SetUI(term UI) {
	r.term = term
	term.SetPrompt(NormalPrompt)
}

func (r *REPL) SetNormal() {
	r.more = ""
	r.SetPrompt(NormalPrompt)
}

func (r *REPL) SetPrompt(prompt string) {
	if r.term != nil {
		r.term.SetPrompt(prompt)
	}
}

func (r *REPL) IsNormal() bool {
	return r.more == ""
}

func (r *REPL) TryDump(expr string) bool {
	i := r.Interp()
	if i != nil {
		if m, v, ok := i.GetSymbol(expr); ok {
			switch p := m.(type) {
			case *ssa.NamedConst:
				r.Printf("%v %v (const)\n", v, p.Type())
			case *ssa.Global:
				r.Printf("%v %v (global var)\n", reflect.ValueOf(v).Elem().Interface(), p.Type().(*types.Pointer).Elem())
			case *ssa.Type:
				// r.Printf("%v %v\n", p.Type().Underlying(), v)
				return false
			case *ssa.Function:
				n := p.Signature.Params().Len()
				if n == 0 || (n == 1 && p.Signature.Variadic()) {
					return false
				}
				r.Printf("%v %v\n", v, p.Type())
			}
			return true
		}
	}
	return false
}

func (r *REPL) Dump(expr string) {
	i := r.Interp()
	if i != nil {
		if m, v, ok := i.GetSymbol(expr); ok {
			switch p := m.(type) {
			case *ssa.NamedConst:
				r.Printf("%v %v (const)\n", v, p.Type())
			case *ssa.Global:
				r.Printf("%v %v (global var)\n", reflect.ValueOf(v).Elem().Interface(), p.Type().(*types.Pointer).Elem())
			case *ssa.Type:
				r.Printf("%v %v\n", p.Type().Underlying(), v)
			case *ssa.Function:
				r.Printf("%v %v\n", v, p.Type())
			}
			return
		}
	}
	_, _, err := r.Eval(fmt.Sprintf("__igop_repl_info__(%v)", expr))
	if err != nil {
		r.tryPkg(expr)
		return
		if err := r.godoc(expr); err != nil {
			r.Printf("not found %v\n", expr)
		}
	}
}

// parseSymbol breaks str apart into a pkg, symbol and method
// fmt.Println => fmt Println
// fmt.Stringer.String => fmt Stringer.String
func parseSymbol(str string) (pkg, symbol string, method string, err error) {
	if str == "" {
		return
	}
	elem := strings.Split(str, ".")
	switch len(elem) {
	case 1:
	case 2:
		symbol = elem[1]
	case 3:
		symbol = elem[1]
		method = elem[2]
	default:
		err = errors.New("too many periods in symbol specification")
		return
	}
	pkg = elem[0]
	return
}

func findPkg(name string) (pkg string, found bool) {
	list := igop.PackageList()
	for _, v := range list {
		if name == v {
			return v, true
		}
	}
	for _, v := range list {
		if strings.HasSuffix(v, "/"+name) {
			return v, true
		}
	}
	return
}

func lookupSymbol(p *igop.Package, sym string) (info string, found bool) {
	if v, ok := p.UntypedConsts[sym]; ok {
		return fmt.Sprintf("const %v.%v %v = %v", p.Name, sym, v.Typ, v.Value), true
	}
	if v, ok := p.TypedConsts[sym]; ok {
		return fmt.Sprintf("const %v.%v %v = %v", p.Name, sym, v.Typ, v.Value), true
	}
	if v, ok := p.Vars[sym]; ok {
		return fmt.Sprintf("var %v.%v %v", p.Name, sym, v.Type().Elem()), true
	}
	if v, ok := p.Funcs[sym]; ok {
		var buf bytes.Buffer
		writeFunc(&buf, p.Name+"."+sym, v.Type())
		return buf.String(), true
	}
	if t, ok := p.NamedTypes[sym]; ok {
		var buf bytes.Buffer
		if t.Typ.Kind() == reflect.Struct {
			writeStruct(&buf, p.Name+"."+sym, t.Typ)
			buf.WriteByte('\n')
		} else {
			fmt.Fprintf(&buf, "type %v.%v %v\n", p.Name, sym, t.Typ.Kind())
		}
		for _, name := range strings.Split(t.Methods, ",") {
			if m, ok := t.Typ.MethodByName(name); ok {
				writeMethod(&buf, m)
				buf.WriteByte('\n')
			}
		}
		ptyp := reflect.PtrTo(t.Typ)
		for _, name := range strings.Split(t.PtrMethods, ",") {
			if m, ok := ptyp.MethodByName(name); ok {
				writeMethod(&buf, m)
				buf.WriteByte('\n')
			}
		}
		return buf.String(), true
	}
	if t, ok := p.Interfaces[sym]; ok {
		var buf bytes.Buffer
		writeInterface(&buf, p.Name+"."+sym, t)
		buf.WriteByte('\n')
		return buf.String(), true
	}
	return
}

func dumpPkg(p *igop.Package) string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "package %v // import %q\n\n", p.Name, p.Path)
	// untyped const
	var uconst []string
	for v := range p.UntypedConsts {
		uconst = append(uconst, v)
	}
	sort.Strings(uconst)
	for _, v := range uconst {
		t := p.UntypedConsts[v]
		fmt.Fprintf(&buf, "const %v = %v\n", v, t.Value)
	}
	// typed const
	var tconst []string
	for v := range p.TypedConsts {
		tconst = append(tconst, v)
	}
	sort.Strings(tconst)
	for _, v := range tconst {
		t := p.TypedConsts[v]
		fmt.Fprintf(&buf, "const %v %v = %v\n", v, t.Typ, t.Value)
	}
	// var
	var vars []string
	for v := range p.Vars {
		vars = append(vars, v)
	}
	sort.Strings(vars)
	for _, v := range vars {
		t := p.Vars[v]
		fmt.Fprintf(&buf, "var %v %v\n", v, t.Elem().Type())
	}
	// funcs
	var funcs []string
	for v := range p.Funcs {
		funcs = append(funcs, v)
	}
	sort.Strings(funcs)
	for _, v := range funcs {
		f := p.Funcs[v]
		writeFunc(&buf, v, f.Type())
		buf.WriteByte('\n')
	}
	// named types
	var types []string
	for v := range p.NamedTypes {
		types = append(types, v)
	}
	sort.Strings(types)
	for _, v := range types {
		t := p.NamedTypes[v]
		fmt.Fprintf(&buf, "type %v %v\n", v, t.Typ.Kind())
		for _, name := range strings.Split(t.Methods, ",") {
			if m, ok := t.Typ.MethodByName(name); ok {
				buf.WriteString("    ")
				writeMethod(&buf, m)
				buf.WriteByte('\n')
			}
		}
		ptyp := reflect.PtrTo(t.Typ)
		for _, name := range strings.Split(t.PtrMethods, ",") {
			if m, ok := ptyp.MethodByName(name); ok {
				buf.WriteString("    ")
				writeMethod(&buf, m)
				buf.WriteByte('\n')
			}
		}
	}
	// interface
	var ifaces []string
	for v := range p.Interfaces {
		ifaces = append(ifaces, v)
	}
	sort.Strings(ifaces)
	for _, v := range ifaces {
		t := p.Interfaces[v]
		fmt.Fprintf(&buf, "type %v interface\n", v)
		n := t.NumMethod()
		for i := 0; i < n; i++ {
			m := t.Method(i)
			buf.WriteString("    ")
			writeFunc(&buf, m.Name, m.Type)
			buf.WriteByte('\n')
		}
	}
	return buf.String()
}

func writeInterface(w io.Writer, name string, typ reflect.Type) {
	n := typ.NumMethod()
	if n == 0 {
		fmt.Fprintf(w, "type %v interface{}", name)
		return
	}
	fmt.Fprintf(w, "type %v interface{\n", name)
	for i := 0; i < n; i++ {
		m := typ.Method(i)
		w.Write([]byte("    "))
		writeFunc(w, m.Name, m.Type)
		w.Write([]byte{'\n'})
	}
	w.Write([]byte{'}'})
}

func writeStruct(w io.Writer, name string, typ reflect.Type) {
	n := typ.NumField()
	if n == 0 {
		fmt.Fprintf(w, "type %v struct{}", name)
		return
	}
	fmt.Fprintf(w, "type %v struct{\n", name)
	for i := 0; i < n; i++ {
		f := typ.Field(i)
		if f.PkgPath != "" {
			fmt.Fprintf(w, "    // ... other fields elided ...\n")
			break
		}
		if f.Anonymous {
			fmt.Fprintf(w, "    %v\n", f.Type)
		} else {
			fmt.Fprintf(w, "    %v %v\n", f.Name, f.Type)
		}
	}
	fmt.Fprintf(w, "}")
}

func writeMethod(w io.Writer, m reflect.Method) {
	typ := m.Type
	numIn := typ.NumIn()
	numOut := typ.NumOut()
	var ins []string
	if typ.IsVariadic() {
		for i := 1; i < numIn-1; i++ {
			ins = append(ins, typ.In(i).String())
		}
		ins = append(ins, "..."+typ.In(numIn-1).Elem().String())
	} else {
		for i := 1; i < numIn; i++ {
			ins = append(ins, typ.In(i).String())
		}
	}
	switch numOut {
	case 0:
		fmt.Fprintf(w, "func (%v) %v(%v)", typ.In(0), m.Name, strings.Join(ins, ", "))
	case 1:
		fmt.Fprintf(w, "func (%v) %v(%v) %v", typ.In(0), m.Name, strings.Join(ins, ", "), typ.Out(0).String())
	default:
		var outs []string
		for i := 0; i < numOut; i++ {
			outs = append(outs, typ.Out(i).String())
		}
		fmt.Fprintf(w, "func (%v) %v(%v) (%v)", typ.In(0), m.Name, strings.Join(ins, ", "), strings.Join(outs, ", "))
	}
}

func writeFunc(w io.Writer, name string, typ reflect.Type) {
	numIn := typ.NumIn()
	numOut := typ.NumOut()
	var ins []string
	if typ.IsVariadic() {
		for i := 0; i < numIn-1; i++ {
			ins = append(ins, typ.In(i).String())
		}
		ins = append(ins, "..."+typ.In(numIn-1).Elem().String())
	} else {
		for i := 0; i < numIn; i++ {
			ins = append(ins, typ.In(i).String())
		}
	}
	switch numOut {
	case 0:
		fmt.Fprintf(w, "func %v(%v)", name, strings.Join(ins, ", "))
	case 1:
		fmt.Fprintf(w, "func %v(%v) %v", name, strings.Join(ins, ", "), typ.Out(0).String())
	default:
		var outs []string
		for i := 0; i < numOut; i++ {
			outs = append(outs, typ.Out(i).String())
		}
		fmt.Fprintf(w, "func %v(%v) (%v)", name, strings.Join(ins, ", "), strings.Join(outs, ", "))
	}
}

func (r *REPL) tryPkg(expr string) error {
	pkg, sym, _, err := parseSymbol(expr)
	if err != nil {
		return err
	}
	if pkgPath, found := findPkg(pkg); found {
		if p, found := igop.LookupPackage(pkgPath); found {
			if sym == "" {
				r.Printf("%v\n", dumpPkg(p))
				return nil
			}
			if info, ok := lookupSymbol(p, sym); ok {
				r.Printf("%v\n", info)
				return nil
			}
			return fmt.Errorf("not found symbol %v.%v", pkg, sym)
		}
	}
	return fmt.Errorf("not found pkg %v", pkg)
}

func (r *REPL) godoc(expr string) error {
	gobin, err := exec.LookPath("go")
	if err != nil {
		return err
	}
	cmd := exec.Command(gobin, "doc", expr)
	data, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	r.Printf("%v\n", string(data))
	return nil
}

func (r *REPL) Printf(_fmt string, a ...interface{}) {
	if r.term != nil {
		r.term.Printf(_fmt, a...)
	} else {
		fmt.Printf(_fmt, a...)
	}
}

var (
	regWord = regexp.MustCompile("\\w+")
)

func (r *REPL) Run(line string) {
	var expr string
	if r.more != "" {
		if line == "" {
			r.SetNormal()
			return
		}
		expr = r.more + "\n" + line
	} else {
		if strings.HasPrefix(line, "?") {
			r.Dump(strings.TrimSpace(line[1:]))
			return
		} else if regWord.MatchString(line) {
			if r.TryDump(line) {
				return
			}
		}
		expr = line
	}
	tok, dump, err := r.Eval(expr)
	if err != nil {
		if checkMore(tok, err) {
			r.more += "\n" + line
			r.SetPrompt(ContinuePrompt)
		} else {
			r.SetNormal()
			r.Printf("error: %v\n", err)
		}
		return
	}
	switch len(dump) {
	case 0:
	case 1:
		r.Printf("%v\n", dump[0])
	default:
		r.Printf("(%v)\n", strings.Join(dump, ", "))
	}
	r.SetNormal()
}

func checkMore(tok token.Token, err error) bool {
	s := err.Error()
	if strings.Contains(s, `expected `) {
		return true
	}
	return false
}
