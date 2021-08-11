// export by github.com/goplus/interp/cmd/qexp

package log

import (
	"log"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("log", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*log.Logger).Fatal":     (*log.Logger).Fatal,
	"(*log.Logger).Fatalf":    (*log.Logger).Fatalf,
	"(*log.Logger).Fatalln":   (*log.Logger).Fatalln,
	"(*log.Logger).Flags":     (*log.Logger).Flags,
	"(*log.Logger).Output":    (*log.Logger).Output,
	"(*log.Logger).Panic":     (*log.Logger).Panic,
	"(*log.Logger).Panicf":    (*log.Logger).Panicf,
	"(*log.Logger).Panicln":   (*log.Logger).Panicln,
	"(*log.Logger).Prefix":    (*log.Logger).Prefix,
	"(*log.Logger).Print":     (*log.Logger).Print,
	"(*log.Logger).Printf":    (*log.Logger).Printf,
	"(*log.Logger).Println":   (*log.Logger).Println,
	"(*log.Logger).SetFlags":  (*log.Logger).SetFlags,
	"(*log.Logger).SetOutput": (*log.Logger).SetOutput,
	"(*log.Logger).SetPrefix": (*log.Logger).SetPrefix,
	"(*log.Logger).Writer":    (*log.Logger).Writer,
	"log.Fatal":               log.Fatal,
	"log.Fatalf":              log.Fatalf,
	"log.Fatalln":             log.Fatalln,
	"log.Flags":               log.Flags,
	"log.New":                 log.New,
	"log.Output":              log.Output,
	"log.Panic":               log.Panic,
	"log.Panicf":              log.Panicf,
	"log.Panicln":             log.Panicln,
	"log.Prefix":              log.Prefix,
	"log.Print":               log.Print,
	"log.Printf":              log.Printf,
	"log.Println":             log.Println,
	"log.SetFlags":            log.SetFlags,
	"log.SetOutput":           log.SetOutput,
	"log.SetPrefix":           log.SetPrefix,
	"log.Writer":              log.Writer,
}

var typList = []interface{}{
	(*log.Logger)(nil),
}
