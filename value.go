// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package interp

// Values
//
// All interpreter values are "boxed" in the empty interface, value.
// The range of possible dynamic types within value are:
//
// - bool
// - numbers (all built-in int/float/complex types are distinguished)
// - string
// - map[value]value --- maps for which  usesBuiltinMap(keyType)
//   *hashmap        --- maps for which !usesBuiltinMap(keyType)
// - chan value
// - []value --- slices
// - iface --- interfaces.
// - structure --- structs.  Fields are ordered and accessed by numeric indices.
// - array --- arrays.
// - *value --- pointers.  Careful: *value is a distinct type from *array etc.
// - *ssa.Function \
//   *ssa.Builtin   } --- functions.  A nil 'func' is always of type *ssa.Function.
//   *closure      /
// - tuple --- as returned by Return, Next, "value,ok" modes, etc.
// - iter --- iterators from 'range' over map or string.
// - bad --- a poison pill for locals that have gone out of scope.
// - rtype -- the interpreter's concrete implementation of reflect.Type
//
// Note that nil is not on this list.
//
// Pay close attention to whether or not the dynamic type is a pointer.
// The compiler cannot help you since value is an empty interface.

import (
	"bytes"
	"fmt"
	"go/types"
	"io"
	"reflect"
	"strings"
	"unsafe"

	"golang.org/x/tools/go/ssa"
)

type value interface{}

type tuple []value

type iface struct {
	t types.Type // never an "untyped" type
	v value
}

// For map, array, *array, slice, string or channel.
type iter interface {
	// next returns a Tuple (key, value, ok).
	// key and value are unaliased, e.g. copies of the sequence element.
	next() tuple
}

type closure struct {
	Fn  *ssa.Function
	Env []value
}

type bad struct{}

// Prints in the style of built-in println.
// (More or less; in gc println is actually a compiler intrinsic and
// can distinguish println(1) from println(interface{}(1)).)
func writeValue(buf *bytes.Buffer, v value) {
	switch v := v.(type) {
	case nil, bool, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, float32, float64, complex64, complex128, string:
		fmt.Fprintf(buf, "%v", v)

	case *ssa.Function, *ssa.Builtin, *closure:
		fmt.Fprintf(buf, "%p", v) // (an address)

	case tuple:
		// Unreachable in well-formed Go programs
		buf.WriteString("(")
		for i, e := range v {
			if i > 0 {
				buf.WriteString(", ")
			}
			writeValue(buf, e)
		}
		buf.WriteString(")")

	default:
		i := reflect.ValueOf(v)
		switch i.Kind() {
		case reflect.Map, reflect.Ptr, reflect.Func, reflect.Chan, reflect.UnsafePointer:
			fmt.Fprintf(buf, "%p", v)
		case reflect.Slice:
			fmt.Fprintf(buf, "[%v/%v]%p", i.Len(), i.Cap(), v)
		case reflect.String:
			fmt.Fprintf(buf, "%v", v)
		case reflect.Struct, reflect.Array:
			panic(fmt.Errorf("illegal types for operand: print %T", v))
		default:
			fmt.Fprintf(buf, "%v", v)
		}
	}
}

// Implements printing of Go values in the style of built-in println.
func toString(v value) string {
	var b bytes.Buffer
	writeValue(&b, v)
	return b.String()
}

// emptyInterface is the header for an interface{} value.
type emptyInterface struct {
	typ  unsafe.Pointer
	word unsafe.Pointer
}

func toInterface(i value) string {
	eface := *(*emptyInterface)(unsafe.Pointer(&i))
	return fmt.Sprintf("(%p,%p)", eface.typ, eface.word)
}

// ------------------------------------------------------------------------
// Iterators

type stringIter struct {
	*strings.Reader
	i int
}

func (it *stringIter) next() tuple {
	okv := make(tuple, 3)
	ch, n, err := it.ReadRune()
	ok := err != io.EOF
	okv[0] = ok
	if ok {
		okv[1] = it.i
		okv[2] = ch
	}
	it.i += n
	return okv
}

type mapIter struct {
	iter *reflect.MapIter
	ok   bool
}

func (it *mapIter) next() tuple {
	it.ok = it.iter.Next()
	if !it.ok {
		return []value{false, nil, nil}
	}
	k, v := it.iter.Key().Interface(), it.iter.Value().Interface()
	return []value{true, k, v}
}
