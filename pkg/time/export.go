// export by github.com/goplus/gossa/cmd/qexp

package time

import (
	q "time"

	"go/constant"
	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "time",
		Path: "time",
		Deps: map[string]string{
			"errors":  "errors",
			"runtime": "runtime",
			"sync":    "sync",
			"syscall": "syscall",
			"unsafe":  "unsafe",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]gossa.NamedType{
			"Duration":   {reflect.TypeOf((*q.Duration)(nil)).Elem(), "Hours,Microseconds,Milliseconds,Minutes,Nanoseconds,Round,Seconds,String,Truncate", ""},
			"Location":   {reflect.TypeOf((*q.Location)(nil)).Elem(), "", "String,firstZoneUsed,get,lookup,lookupFirstZone,lookupName"},
			"Month":      {reflect.TypeOf((*q.Month)(nil)).Elem(), "String", ""},
			"ParseError": {reflect.TypeOf((*q.ParseError)(nil)).Elem(), "", "Error"},
			"Ticker":     {reflect.TypeOf((*q.Ticker)(nil)).Elem(), "", "Reset,Stop"},
			"Time":       {reflect.TypeOf((*q.Time)(nil)).Elem(), "Add,AddDate,After,AppendFormat,Before,Clock,Date,Day,Equal,Format,GobEncode,Hour,ISOWeek,In,IsZero,Local,Location,MarshalBinary,MarshalJSON,MarshalText,Minute,Month,Nanosecond,Round,Second,String,Sub,Truncate,UTC,Unix,UnixNano,Weekday,Year,YearDay,Zone,abs,date,locabs", "GobDecode,UnmarshalBinary,UnmarshalJSON,UnmarshalText,addSec,mono,nsec,sec,setLoc,setMono,stripMono,unixSec"},
			"Timer":      {reflect.TypeOf((*q.Timer)(nil)).Elem(), "", "Reset,Stop"},
			"Weekday":    {reflect.TypeOf((*q.Weekday)(nil)).Elem(), "String", ""},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"Local": reflect.ValueOf(&q.Local),
			"UTC":   reflect.ValueOf(&q.UTC),
		},
		Funcs: map[string]reflect.Value{
			"After":                  reflect.ValueOf(q.After),
			"AfterFunc":              reflect.ValueOf(q.AfterFunc),
			"Date":                   reflect.ValueOf(q.Date),
			"FixedZone":              reflect.ValueOf(q.FixedZone),
			"LoadLocation":           reflect.ValueOf(q.LoadLocation),
			"LoadLocationFromTZData": reflect.ValueOf(q.LoadLocationFromTZData),
			"NewTicker":              reflect.ValueOf(q.NewTicker),
			"NewTimer":               reflect.ValueOf(q.NewTimer),
			"Now":                    reflect.ValueOf(q.Now),
			"Parse":                  reflect.ValueOf(q.Parse),
			"ParseDuration":          reflect.ValueOf(q.ParseDuration),
			"ParseInLocation":        reflect.ValueOf(q.ParseInLocation),
			"Since":                  reflect.ValueOf(q.Since),
			"Sleep":                  reflect.ValueOf(q.Sleep),
			"Tick":                   reflect.ValueOf(q.Tick),
			"Unix":                   reflect.ValueOf(q.Unix),
			"Until":                  reflect.ValueOf(q.Until),
		},
		TypedConsts: map[string]gossa.TypedConst{
			"April":       {reflect.TypeOf(q.April), constant.MakeInt64(4)},
			"August":      {reflect.TypeOf(q.August), constant.MakeInt64(8)},
			"December":    {reflect.TypeOf(q.December), constant.MakeInt64(12)},
			"February":    {reflect.TypeOf(q.February), constant.MakeInt64(2)},
			"Friday":      {reflect.TypeOf(q.Friday), constant.MakeInt64(5)},
			"Hour":        {reflect.TypeOf(q.Hour), constant.MakeInt64(3600000000000)},
			"January":     {reflect.TypeOf(q.January), constant.MakeInt64(1)},
			"July":        {reflect.TypeOf(q.July), constant.MakeInt64(7)},
			"June":        {reflect.TypeOf(q.June), constant.MakeInt64(6)},
			"March":       {reflect.TypeOf(q.March), constant.MakeInt64(3)},
			"May":         {reflect.TypeOf(q.May), constant.MakeInt64(5)},
			"Microsecond": {reflect.TypeOf(q.Microsecond), constant.MakeInt64(1000)},
			"Millisecond": {reflect.TypeOf(q.Millisecond), constant.MakeInt64(1000000)},
			"Minute":      {reflect.TypeOf(q.Minute), constant.MakeInt64(60000000000)},
			"Monday":      {reflect.TypeOf(q.Monday), constant.MakeInt64(1)},
			"Nanosecond":  {reflect.TypeOf(q.Nanosecond), constant.MakeInt64(1)},
			"November":    {reflect.TypeOf(q.November), constant.MakeInt64(11)},
			"October":     {reflect.TypeOf(q.October), constant.MakeInt64(10)},
			"Saturday":    {reflect.TypeOf(q.Saturday), constant.MakeInt64(6)},
			"Second":      {reflect.TypeOf(q.Second), constant.MakeInt64(1000000000)},
			"September":   {reflect.TypeOf(q.September), constant.MakeInt64(9)},
			"Sunday":      {reflect.TypeOf(q.Sunday), constant.MakeInt64(0)},
			"Thursday":    {reflect.TypeOf(q.Thursday), constant.MakeInt64(4)},
			"Tuesday":     {reflect.TypeOf(q.Tuesday), constant.MakeInt64(2)},
			"Wednesday":   {reflect.TypeOf(q.Wednesday), constant.MakeInt64(3)},
		},
		UntypedConsts: map[string]gossa.UntypedConst{
			"ANSIC":       {"untyped string", constant.MakeString("Mon Jan _2 15:04:05 2006")},
			"Kitchen":     {"untyped string", constant.MakeString("3:04PM")},
			"RFC1123":     {"untyped string", constant.MakeString("Mon, 02 Jan 2006 15:04:05 MST")},
			"RFC1123Z":    {"untyped string", constant.MakeString("Mon, 02 Jan 2006 15:04:05 -0700")},
			"RFC3339":     {"untyped string", constant.MakeString("2006-01-02T15:04:05Z07:00")},
			"RFC3339Nano": {"untyped string", constant.MakeString("2006-01-02T15:04:05.999999999Z07:00")},
			"RFC822":      {"untyped string", constant.MakeString("02 Jan 06 15:04 MST")},
			"RFC822Z":     {"untyped string", constant.MakeString("02 Jan 06 15:04 -0700")},
			"RFC850":      {"untyped string", constant.MakeString("Monday, 02-Jan-06 15:04:05 MST")},
			"RubyDate":    {"untyped string", constant.MakeString("Mon Jan 02 15:04:05 -0700 2006")},
			"Stamp":       {"untyped string", constant.MakeString("Jan _2 15:04:05")},
			"StampMicro":  {"untyped string", constant.MakeString("Jan _2 15:04:05.000000")},
			"StampMilli":  {"untyped string", constant.MakeString("Jan _2 15:04:05.000")},
			"StampNano":   {"untyped string", constant.MakeString("Jan _2 15:04:05.000000000")},
			"UnixDate":    {"untyped string", constant.MakeString("Mon Jan _2 15:04:05 MST 2006")},
		},
	})
}
