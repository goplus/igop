// export by github.com/goplus/gossa/cmd/qexp

package reflect

import (
	q "reflect"

	"go/constant"
	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "reflect",
		Path: "reflect",
		Deps: map[string]string{
			"internal/unsafeheader": "unsafeheader",
			"math":                  "math",
			"runtime":               "runtime",
			"strconv":               "strconv",
			"sync":                  "sync",
			"unicode":               "unicode",
			"unicode/utf8":          "utf8",
			"unsafe":                "unsafe",
		},
		Interfaces: map[string]reflect.Type{
			"Type": reflect.TypeOf((*q.Type)(nil)).Elem(),
		},
		NamedTypes: map[string]gossa.NamedType{
			"ChanDir":      {reflect.TypeOf((*q.ChanDir)(nil)).Elem(), "String", ""},
			"Kind":         {reflect.TypeOf((*q.Kind)(nil)).Elem(), "String", ""},
			"MapIter":      {reflect.TypeOf((*q.MapIter)(nil)).Elem(), "", "Key,Next,Value"},
			"Method":       {reflect.TypeOf((*q.Method)(nil)).Elem(), "", ""},
			"SelectCase":   {reflect.TypeOf((*q.SelectCase)(nil)).Elem(), "", ""},
			"SelectDir":    {reflect.TypeOf((*q.SelectDir)(nil)).Elem(), "", ""},
			"SliceHeader":  {reflect.TypeOf((*q.SliceHeader)(nil)).Elem(), "", ""},
			"StringHeader": {reflect.TypeOf((*q.StringHeader)(nil)).Elem(), "", ""},
			"StructField":  {reflect.TypeOf((*q.StructField)(nil)).Elem(), "", ""},
			"StructTag":    {reflect.TypeOf((*q.StructTag)(nil)).Elem(), "Get,Lookup", ""},
			"Value":        {reflect.TypeOf((*q.Value)(nil)).Elem(), "Addr,Bool,Bytes,Call,CallSlice,CanAddr,CanInterface,CanSet,Cap,Close,Complex,Convert,Elem,Field,FieldByIndex,FieldByName,FieldByNameFunc,Float,Index,Int,Interface,InterfaceData,IsNil,IsValid,IsZero,Kind,Len,MapIndex,MapKeys,MapRange,Method,MethodByName,NumField,NumMethod,OverflowComplex,OverflowFloat,OverflowInt,OverflowUint,Pointer,Recv,Send,Set,SetBool,SetBytes,SetCap,SetComplex,SetFloat,SetInt,SetLen,SetMapIndex,SetPointer,SetString,SetUint,Slice,Slice3,String,TryRecv,TrySend,Type,Uint,UnsafeAddr,assignTo,call,pointer,recv,runes,send,setRunes", ""},
			"ValueError":   {reflect.TypeOf((*q.ValueError)(nil)).Elem(), "", "Error"},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Append":          reflect.ValueOf(q.Append),
			"AppendSlice":     reflect.ValueOf(q.AppendSlice),
			"ArrayOf":         reflect.ValueOf(q.ArrayOf),
			"ChanOf":          reflect.ValueOf(q.ChanOf),
			"Copy":            reflect.ValueOf(q.Copy),
			"DeepEqual":       reflect.ValueOf(q.DeepEqual),
			"FuncOf":          reflect.ValueOf(q.FuncOf),
			"Indirect":        reflect.ValueOf(q.Indirect),
			"MakeChan":        reflect.ValueOf(q.MakeChan),
			"MakeFunc":        reflect.ValueOf(q.MakeFunc),
			"MakeMap":         reflect.ValueOf(q.MakeMap),
			"MakeMapWithSize": reflect.ValueOf(q.MakeMapWithSize),
			"MakeSlice":       reflect.ValueOf(q.MakeSlice),
			"MapOf":           reflect.ValueOf(q.MapOf),
			"New":             reflect.ValueOf(q.New),
			"NewAt":           reflect.ValueOf(q.NewAt),
			"PtrTo":           reflect.ValueOf(q.PtrTo),
			"Select":          reflect.ValueOf(q.Select),
			"SliceOf":         reflect.ValueOf(q.SliceOf),
			"StructOf":        reflect.ValueOf(q.StructOf),
			"Swapper":         reflect.ValueOf(q.Swapper),
			"TypeOf":          reflect.ValueOf(q.TypeOf),
			"ValueOf":         reflect.ValueOf(q.ValueOf),
			"Zero":            reflect.ValueOf(q.Zero),
		},
		TypedConsts: map[string]gossa.TypedConst{
			"Array":         {reflect.TypeOf(q.Array), constant.MakeInt64(17)},
			"Bool":          {reflect.TypeOf(q.Bool), constant.MakeInt64(1)},
			"BothDir":       {reflect.TypeOf(q.BothDir), constant.MakeInt64(3)},
			"Chan":          {reflect.TypeOf(q.Chan), constant.MakeInt64(18)},
			"Complex128":    {reflect.TypeOf(q.Complex128), constant.MakeInt64(16)},
			"Complex64":     {reflect.TypeOf(q.Complex64), constant.MakeInt64(15)},
			"Float32":       {reflect.TypeOf(q.Float32), constant.MakeInt64(13)},
			"Float64":       {reflect.TypeOf(q.Float64), constant.MakeInt64(14)},
			"Func":          {reflect.TypeOf(q.Func), constant.MakeInt64(19)},
			"Int":           {reflect.TypeOf(q.Int), constant.MakeInt64(2)},
			"Int16":         {reflect.TypeOf(q.Int16), constant.MakeInt64(4)},
			"Int32":         {reflect.TypeOf(q.Int32), constant.MakeInt64(5)},
			"Int64":         {reflect.TypeOf(q.Int64), constant.MakeInt64(6)},
			"Int8":          {reflect.TypeOf(q.Int8), constant.MakeInt64(3)},
			"Interface":     {reflect.TypeOf(q.Interface), constant.MakeInt64(20)},
			"Invalid":       {reflect.TypeOf(q.Invalid), constant.MakeInt64(0)},
			"Map":           {reflect.TypeOf(q.Map), constant.MakeInt64(21)},
			"Ptr":           {reflect.TypeOf(q.Ptr), constant.MakeInt64(22)},
			"RecvDir":       {reflect.TypeOf(q.RecvDir), constant.MakeInt64(1)},
			"SelectDefault": {reflect.TypeOf(q.SelectDefault), constant.MakeInt64(3)},
			"SelectRecv":    {reflect.TypeOf(q.SelectRecv), constant.MakeInt64(2)},
			"SelectSend":    {reflect.TypeOf(q.SelectSend), constant.MakeInt64(1)},
			"SendDir":       {reflect.TypeOf(q.SendDir), constant.MakeInt64(2)},
			"Slice":         {reflect.TypeOf(q.Slice), constant.MakeInt64(23)},
			"String":        {reflect.TypeOf(q.String), constant.MakeInt64(24)},
			"Struct":        {reflect.TypeOf(q.Struct), constant.MakeInt64(25)},
			"Uint":          {reflect.TypeOf(q.Uint), constant.MakeInt64(7)},
			"Uint16":        {reflect.TypeOf(q.Uint16), constant.MakeInt64(9)},
			"Uint32":        {reflect.TypeOf(q.Uint32), constant.MakeInt64(10)},
			"Uint64":        {reflect.TypeOf(q.Uint64), constant.MakeInt64(11)},
			"Uint8":         {reflect.TypeOf(q.Uint8), constant.MakeInt64(8)},
			"Uintptr":       {reflect.TypeOf(q.Uintptr), constant.MakeInt64(12)},
			"UnsafePointer": {reflect.TypeOf(q.UnsafePointer), constant.MakeInt64(26)},
		},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}
