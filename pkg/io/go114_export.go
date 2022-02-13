// export by github.com/goplus/gossa/cmd/qexp

//+build go1.14,!go1.15

package io

import (
	q "io"

	"go/constant"
	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "io",
		Path: "io",
		Deps: map[string]string{
			"errors": "errors",
			"sync":   "sync",
		},
		Interfaces: map[string]reflect.Type{
			"ByteReader":      reflect.TypeOf((*q.ByteReader)(nil)).Elem(),
			"ByteScanner":     reflect.TypeOf((*q.ByteScanner)(nil)).Elem(),
			"ByteWriter":      reflect.TypeOf((*q.ByteWriter)(nil)).Elem(),
			"Closer":          reflect.TypeOf((*q.Closer)(nil)).Elem(),
			"ReadCloser":      reflect.TypeOf((*q.ReadCloser)(nil)).Elem(),
			"ReadSeeker":      reflect.TypeOf((*q.ReadSeeker)(nil)).Elem(),
			"ReadWriteCloser": reflect.TypeOf((*q.ReadWriteCloser)(nil)).Elem(),
			"ReadWriteSeeker": reflect.TypeOf((*q.ReadWriteSeeker)(nil)).Elem(),
			"ReadWriter":      reflect.TypeOf((*q.ReadWriter)(nil)).Elem(),
			"Reader":          reflect.TypeOf((*q.Reader)(nil)).Elem(),
			"ReaderAt":        reflect.TypeOf((*q.ReaderAt)(nil)).Elem(),
			"ReaderFrom":      reflect.TypeOf((*q.ReaderFrom)(nil)).Elem(),
			"RuneReader":      reflect.TypeOf((*q.RuneReader)(nil)).Elem(),
			"RuneScanner":     reflect.TypeOf((*q.RuneScanner)(nil)).Elem(),
			"Seeker":          reflect.TypeOf((*q.Seeker)(nil)).Elem(),
			"StringWriter":    reflect.TypeOf((*q.StringWriter)(nil)).Elem(),
			"WriteCloser":     reflect.TypeOf((*q.WriteCloser)(nil)).Elem(),
			"WriteSeeker":     reflect.TypeOf((*q.WriteSeeker)(nil)).Elem(),
			"Writer":          reflect.TypeOf((*q.Writer)(nil)).Elem(),
			"WriterAt":        reflect.TypeOf((*q.WriterAt)(nil)).Elem(),
			"WriterTo":        reflect.TypeOf((*q.WriterTo)(nil)).Elem(),
		},
		NamedTypes: map[string]gossa.NamedType{
			"LimitedReader": {reflect.TypeOf((*q.LimitedReader)(nil)).Elem(), "", "Read"},
			"PipeReader":    {reflect.TypeOf((*q.PipeReader)(nil)).Elem(), "", "Close,CloseWithError,Read"},
			"PipeWriter":    {reflect.TypeOf((*q.PipeWriter)(nil)).Elem(), "", "Close,CloseWithError,Write"},
			"SectionReader": {reflect.TypeOf((*q.SectionReader)(nil)).Elem(), "", "Read,ReadAt,Seek,Size"},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"EOF":              reflect.ValueOf(&q.EOF),
			"ErrClosedPipe":    reflect.ValueOf(&q.ErrClosedPipe),
			"ErrNoProgress":    reflect.ValueOf(&q.ErrNoProgress),
			"ErrShortBuffer":   reflect.ValueOf(&q.ErrShortBuffer),
			"ErrShortWrite":    reflect.ValueOf(&q.ErrShortWrite),
			"ErrUnexpectedEOF": reflect.ValueOf(&q.ErrUnexpectedEOF),
		},
		Funcs: map[string]reflect.Value{
			"Copy":             reflect.ValueOf(q.Copy),
			"CopyBuffer":       reflect.ValueOf(q.CopyBuffer),
			"CopyN":            reflect.ValueOf(q.CopyN),
			"LimitReader":      reflect.ValueOf(q.LimitReader),
			"MultiReader":      reflect.ValueOf(q.MultiReader),
			"MultiWriter":      reflect.ValueOf(q.MultiWriter),
			"NewSectionReader": reflect.ValueOf(q.NewSectionReader),
			"Pipe":             reflect.ValueOf(q.Pipe),
			"ReadAtLeast":      reflect.ValueOf(q.ReadAtLeast),
			"ReadFull":         reflect.ValueOf(q.ReadFull),
			"TeeReader":        reflect.ValueOf(q.TeeReader),
			"WriteString":      reflect.ValueOf(q.WriteString),
		},
		TypedConsts: map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{
			"SeekCurrent": {"untyped int", constant.MakeInt64(int64(q.SeekCurrent))},
			"SeekEnd":     {"untyped int", constant.MakeInt64(int64(q.SeekEnd))},
			"SeekStart":   {"untyped int", constant.MakeInt64(int64(q.SeekStart))},
		},
	})
}
