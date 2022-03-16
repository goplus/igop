// export by github.com/goplus/gossa/cmd/qexp

//go:build go1.17 && !go1.18
// +build go1.17,!go1.18

package sql

import (
	q "database/sql"

	"go/constant"
	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "sql",
		Path: "database/sql",
		Deps: map[string]string{
			"context":             "context",
			"database/sql/driver": "driver",
			"errors":              "errors",
			"fmt":                 "fmt",
			"io":                  "io",
			"reflect":             "reflect",
			"runtime":             "runtime",
			"sort":                "sort",
			"strconv":             "strconv",
			"sync":                "sync",
			"sync/atomic":         "atomic",
			"time":                "time",
			"unicode":             "unicode",
			"unicode/utf8":        "utf8",
		},
		Interfaces: map[string]reflect.Type{
			"Result":  reflect.TypeOf((*q.Result)(nil)).Elem(),
			"Scanner": reflect.TypeOf((*q.Scanner)(nil)).Elem(),
		},
		NamedTypes: map[string]gossa.NamedType{
			"ColumnType":     {reflect.TypeOf((*q.ColumnType)(nil)).Elem(), "", "DatabaseTypeName,DecimalSize,Length,Name,Nullable,ScanType"},
			"Conn":           {reflect.TypeOf((*q.Conn)(nil)).Elem(), "", "BeginTx,Close,ExecContext,PingContext,PrepareContext,QueryContext,QueryRowContext,Raw,close,closemuRUnlockCondReleaseConn,grabConn,txCtx"},
			"DB":             {reflect.TypeOf((*q.DB)(nil)).Elem(), "", "Begin,BeginTx,Close,Conn,Driver,Exec,ExecContext,Ping,PingContext,Prepare,PrepareContext,Query,QueryContext,QueryRow,QueryRowContext,SetConnMaxIdleTime,SetConnMaxLifetime,SetMaxIdleConns,SetMaxOpenConns,Stats,addDep,addDepLocked,begin,beginDC,conn,connectionCleaner,connectionCleanerRunLocked,connectionOpener,exec,execDC,maxIdleConnsLocked,maybeOpenNewConnections,nextRequestKeyLocked,noteUnusedDriverStatement,openNewConnection,pingDC,prepare,prepareDC,putConn,putConnDBLocked,query,queryDC,removeDep,removeDepLocked,shortestIdleTimeLocked,startCleanerLocked"},
			"DBStats":        {reflect.TypeOf((*q.DBStats)(nil)).Elem(), "", ""},
			"IsolationLevel": {reflect.TypeOf((*q.IsolationLevel)(nil)).Elem(), "String", ""},
			"NamedArg":       {reflect.TypeOf((*q.NamedArg)(nil)).Elem(), "", ""},
			"NullBool":       {reflect.TypeOf((*q.NullBool)(nil)).Elem(), "Value", "Scan"},
			"NullByte":       {reflect.TypeOf((*q.NullByte)(nil)).Elem(), "Value", "Scan"},
			"NullFloat64":    {reflect.TypeOf((*q.NullFloat64)(nil)).Elem(), "Value", "Scan"},
			"NullInt16":      {reflect.TypeOf((*q.NullInt16)(nil)).Elem(), "Value", "Scan"},
			"NullInt32":      {reflect.TypeOf((*q.NullInt32)(nil)).Elem(), "Value", "Scan"},
			"NullInt64":      {reflect.TypeOf((*q.NullInt64)(nil)).Elem(), "Value", "Scan"},
			"NullString":     {reflect.TypeOf((*q.NullString)(nil)).Elem(), "Value", "Scan"},
			"NullTime":       {reflect.TypeOf((*q.NullTime)(nil)).Elem(), "Value", "Scan"},
			"Out":            {reflect.TypeOf((*q.Out)(nil)).Elem(), "", ""},
			"RawBytes":       {reflect.TypeOf((*q.RawBytes)(nil)).Elem(), "", ""},
			"Row":            {reflect.TypeOf((*q.Row)(nil)).Elem(), "", "Err,Scan"},
			"Rows":           {reflect.TypeOf((*q.Rows)(nil)).Elem(), "", "Close,ColumnTypes,Columns,Err,Next,NextResultSet,Scan,awaitDone,close,initContextClose,lasterrOrErrLocked,nextLocked"},
			"Stmt":           {reflect.TypeOf((*q.Stmt)(nil)).Elem(), "", "Close,Exec,ExecContext,Query,QueryContext,QueryRow,QueryRowContext,connStmt,finalClose,prepareOnConnLocked,removeClosedStmtLocked"},
			"Tx":             {reflect.TypeOf((*q.Tx)(nil)).Elem(), "", "Commit,Exec,ExecContext,Prepare,PrepareContext,Query,QueryContext,QueryRow,QueryRowContext,Rollback,Stmt,StmtContext,awaitDone,close,closePrepared,closemuRUnlockRelease,grabConn,isDone,rollback,txCtx"},
			"TxOptions":      {reflect.TypeOf((*q.TxOptions)(nil)).Elem(), "", ""},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"ErrConnDone": reflect.ValueOf(&q.ErrConnDone),
			"ErrNoRows":   reflect.ValueOf(&q.ErrNoRows),
			"ErrTxDone":   reflect.ValueOf(&q.ErrTxDone),
		},
		Funcs: map[string]reflect.Value{
			"Drivers":  reflect.ValueOf(q.Drivers),
			"Named":    reflect.ValueOf(q.Named),
			"Open":     reflect.ValueOf(q.Open),
			"OpenDB":   reflect.ValueOf(q.OpenDB),
			"Register": reflect.ValueOf(q.Register),
		},
		TypedConsts: map[string]gossa.TypedConst{
			"LevelDefault":         {reflect.TypeOf(q.LevelDefault), constant.MakeInt64(int64(q.LevelDefault))},
			"LevelLinearizable":    {reflect.TypeOf(q.LevelLinearizable), constant.MakeInt64(int64(q.LevelLinearizable))},
			"LevelReadCommitted":   {reflect.TypeOf(q.LevelReadCommitted), constant.MakeInt64(int64(q.LevelReadCommitted))},
			"LevelReadUncommitted": {reflect.TypeOf(q.LevelReadUncommitted), constant.MakeInt64(int64(q.LevelReadUncommitted))},
			"LevelRepeatableRead":  {reflect.TypeOf(q.LevelRepeatableRead), constant.MakeInt64(int64(q.LevelRepeatableRead))},
			"LevelSerializable":    {reflect.TypeOf(q.LevelSerializable), constant.MakeInt64(int64(q.LevelSerializable))},
			"LevelSnapshot":        {reflect.TypeOf(q.LevelSnapshot), constant.MakeInt64(int64(q.LevelSnapshot))},
			"LevelWriteCommitted":  {reflect.TypeOf(q.LevelWriteCommitted), constant.MakeInt64(int64(q.LevelWriteCommitted))},
		},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}
