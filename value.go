// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gossa

import (
	"io"
	"reflect"
	"strings"
	"unsafe"
)

type Value = value
type Tuple = tuple

type value = interface{}

type tuple []value

type closure struct {
	pfn *Function
	env []value
}

// emptyInterface is the header for an interface{} value.
type emptyInterface struct {
	typ  unsafe.Pointer
	word unsafe.Pointer
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
