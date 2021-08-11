// export by github.com/goplus/interp/cmd/qexp

package fmt

import (
	"fmt"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("fmt", extMap, typList)
}

var extMap = map[string]interface{}{
	"(fmt.Formatter).Format":     (fmt.Formatter).Format,
	"(fmt.GoStringer).GoString":  (fmt.GoStringer).GoString,
	"(fmt.ScanState).Read":       (fmt.ScanState).Read,
	"(fmt.ScanState).ReadRune":   (fmt.ScanState).ReadRune,
	"(fmt.ScanState).SkipSpace":  (fmt.ScanState).SkipSpace,
	"(fmt.ScanState).Token":      (fmt.ScanState).Token,
	"(fmt.ScanState).UnreadRune": (fmt.ScanState).UnreadRune,
	"(fmt.ScanState).Width":      (fmt.ScanState).Width,
	"(fmt.Scanner).Scan":         (fmt.Scanner).Scan,
	"(fmt.State).Flag":           (fmt.State).Flag,
	"(fmt.State).Precision":      (fmt.State).Precision,
	"(fmt.State).Width":          (fmt.State).Width,
	"(fmt.State).Write":          (fmt.State).Write,
	"(fmt.Stringer).String":      (fmt.Stringer).String,
	"fmt.Errorf":                 fmt.Errorf,
	"fmt.Fprint":                 fmt.Fprint,
	"fmt.Fprintf":                fmt.Fprintf,
	"fmt.Fprintln":               fmt.Fprintln,
	"fmt.Fscan":                  fmt.Fscan,
	"fmt.Fscanf":                 fmt.Fscanf,
	"fmt.Fscanln":                fmt.Fscanln,
	"fmt.Print":                  fmt.Print,
	"fmt.Printf":                 fmt.Printf,
	"fmt.Println":                fmt.Println,
	"fmt.Scan":                   fmt.Scan,
	"fmt.Scanf":                  fmt.Scanf,
	"fmt.Scanln":                 fmt.Scanln,
	"fmt.Sprint":                 fmt.Sprint,
	"fmt.Sprintf":                fmt.Sprintf,
	"fmt.Sprintln":               fmt.Sprintln,
	"fmt.Sscan":                  fmt.Sscan,
	"fmt.Sscanf":                 fmt.Sscanf,
	"fmt.Sscanln":                fmt.Sscanln,
}

var typList = []interface{}{
	(*fmt.Formatter)(nil),
	(*fmt.GoStringer)(nil),
	(*fmt.ScanState)(nil),
	(*fmt.Scanner)(nil),
	(*fmt.State)(nil),
	(*fmt.Stringer)(nil),
}
