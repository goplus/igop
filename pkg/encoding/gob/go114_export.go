// export by github.com/goplus/gossa/cmd/qexp

//+build go1.14,!go1.15

package gob

import (
	q "encoding/gob"

	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "gob",
		Path: "encoding/gob",
		Deps: map[string]string{
			"bufio":           "bufio",
			"encoding":        "encoding",
			"encoding/binary": "binary",
			"errors":          "errors",
			"fmt":             "fmt",
			"io":              "io",
			"math":            "math",
			"math/bits":       "bits",
			"os":              "os",
			"reflect":         "reflect",
			"sync":            "sync",
			"sync/atomic":     "atomic",
			"unicode":         "unicode",
			"unicode/utf8":    "utf8",
		},
		Interfaces: map[string]reflect.Type{
			"GobDecoder": reflect.TypeOf((*q.GobDecoder)(nil)).Elem(),
			"GobEncoder": reflect.TypeOf((*q.GobEncoder)(nil)).Elem(),
		},
		NamedTypes: map[string]gossa.NamedType{
			"CommonType": {reflect.TypeOf((*q.CommonType)(nil)).Elem(), "", "id,name,safeString,setId,string"},
			"Decoder":    {reflect.TypeOf((*q.Decoder)(nil)).Elem(), "", "Decode,DecodeValue,compatibleType,compileDec,compileIgnoreSingle,compileSingle,decIgnoreOpFor,decOpFor,decodeArray,decodeArrayHelper,decodeGobDecoder,decodeIgnoredValue,decodeInterface,decodeMap,decodeSingle,decodeSlice,decodeStruct,decodeTypeSequence,decodeValue,freeDecoderState,getDecEnginePtr,getIgnoreEnginePtr,gobDecodeOpFor,ignoreArray,ignoreArrayHelper,ignoreGobDecoder,ignoreInterface,ignoreMap,ignoreSingle,ignoreSlice,ignoreStruct,newDecoderState,nextInt,nextUint,readMessage,recvMessage,recvType,typeString"},
			"Encoder":    {reflect.TypeOf((*q.Encoder)(nil)).Elem(), "", "Encode,EncodeValue,encode,encodeArray,encodeGobEncoder,encodeInterface,encodeMap,encodeSingle,encodeStruct,freeEncoderState,newEncoderState,popWriter,pushWriter,sendActualType,sendType,sendTypeDescriptor,sendTypeId,setError,writeMessage,writer"},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"NewDecoder":   reflect.ValueOf(q.NewDecoder),
			"NewEncoder":   reflect.ValueOf(q.NewEncoder),
			"Register":     reflect.ValueOf(q.Register),
			"RegisterName": reflect.ValueOf(q.RegisterName),
		},
		TypedConsts:   map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}
