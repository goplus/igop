// export by github.com/goplus/gossa/cmd/qexp

package bufio

import (
	"bufio"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("bufio", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*bufio.Reader).Buffered":       (*bufio.Reader).Buffered,
	"(*bufio.Reader).Discard":        (*bufio.Reader).Discard,
	"(*bufio.Reader).Peek":           (*bufio.Reader).Peek,
	"(*bufio.Reader).Read":           (*bufio.Reader).Read,
	"(*bufio.Reader).ReadByte":       (*bufio.Reader).ReadByte,
	"(*bufio.Reader).ReadBytes":      (*bufio.Reader).ReadBytes,
	"(*bufio.Reader).ReadLine":       (*bufio.Reader).ReadLine,
	"(*bufio.Reader).ReadRune":       (*bufio.Reader).ReadRune,
	"(*bufio.Reader).ReadSlice":      (*bufio.Reader).ReadSlice,
	"(*bufio.Reader).ReadString":     (*bufio.Reader).ReadString,
	"(*bufio.Reader).Reset":          (*bufio.Reader).Reset,
	"(*bufio.Reader).Size":           (*bufio.Reader).Size,
	"(*bufio.Reader).UnreadByte":     (*bufio.Reader).UnreadByte,
	"(*bufio.Reader).UnreadRune":     (*bufio.Reader).UnreadRune,
	"(*bufio.Reader).WriteTo":        (*bufio.Reader).WriteTo,
	"(*bufio.Scanner).Buffer":        (*bufio.Scanner).Buffer,
	"(*bufio.Scanner).Bytes":         (*bufio.Scanner).Bytes,
	"(*bufio.Scanner).Err":           (*bufio.Scanner).Err,
	"(*bufio.Scanner).Scan":          (*bufio.Scanner).Scan,
	"(*bufio.Scanner).Split":         (*bufio.Scanner).Split,
	"(*bufio.Scanner).Text":          (*bufio.Scanner).Text,
	"(*bufio.Writer).Available":      (*bufio.Writer).Available,
	"(*bufio.Writer).Buffered":       (*bufio.Writer).Buffered,
	"(*bufio.Writer).Flush":          (*bufio.Writer).Flush,
	"(*bufio.Writer).ReadFrom":       (*bufio.Writer).ReadFrom,
	"(*bufio.Writer).Reset":          (*bufio.Writer).Reset,
	"(*bufio.Writer).Size":           (*bufio.Writer).Size,
	"(*bufio.Writer).Write":          (*bufio.Writer).Write,
	"(*bufio.Writer).WriteByte":      (*bufio.Writer).WriteByte,
	"(*bufio.Writer).WriteRune":      (*bufio.Writer).WriteRune,
	"(*bufio.Writer).WriteString":    (*bufio.Writer).WriteString,
	"(bufio.ReadWriter).Available":   (bufio.ReadWriter).Available,
	"(bufio.ReadWriter).Discard":     (bufio.ReadWriter).Discard,
	"(bufio.ReadWriter).Flush":       (bufio.ReadWriter).Flush,
	"(bufio.ReadWriter).Peek":        (bufio.ReadWriter).Peek,
	"(bufio.ReadWriter).Read":        (bufio.ReadWriter).Read,
	"(bufio.ReadWriter).ReadByte":    (bufio.ReadWriter).ReadByte,
	"(bufio.ReadWriter).ReadBytes":   (bufio.ReadWriter).ReadBytes,
	"(bufio.ReadWriter).ReadFrom":    (bufio.ReadWriter).ReadFrom,
	"(bufio.ReadWriter).ReadLine":    (bufio.ReadWriter).ReadLine,
	"(bufio.ReadWriter).ReadRune":    (bufio.ReadWriter).ReadRune,
	"(bufio.ReadWriter).ReadSlice":   (bufio.ReadWriter).ReadSlice,
	"(bufio.ReadWriter).ReadString":  (bufio.ReadWriter).ReadString,
	"(bufio.ReadWriter).UnreadByte":  (bufio.ReadWriter).UnreadByte,
	"(bufio.ReadWriter).UnreadRune":  (bufio.ReadWriter).UnreadRune,
	"(bufio.ReadWriter).Write":       (bufio.ReadWriter).Write,
	"(bufio.ReadWriter).WriteByte":   (bufio.ReadWriter).WriteByte,
	"(bufio.ReadWriter).WriteRune":   (bufio.ReadWriter).WriteRune,
	"(bufio.ReadWriter).WriteString": (bufio.ReadWriter).WriteString,
	"(bufio.ReadWriter).WriteTo":     (bufio.ReadWriter).WriteTo,
	"bufio.ErrAdvanceTooFar":         &bufio.ErrAdvanceTooFar,
	"bufio.ErrBufferFull":            &bufio.ErrBufferFull,
	"bufio.ErrFinalToken":            &bufio.ErrFinalToken,
	"bufio.ErrInvalidUnreadByte":     &bufio.ErrInvalidUnreadByte,
	"bufio.ErrInvalidUnreadRune":     &bufio.ErrInvalidUnreadRune,
	"bufio.ErrNegativeAdvance":       &bufio.ErrNegativeAdvance,
	"bufio.ErrNegativeCount":         &bufio.ErrNegativeCount,
	"bufio.ErrTooLong":               &bufio.ErrTooLong,
	"bufio.NewReadWriter":            bufio.NewReadWriter,
	"bufio.NewReader":                bufio.NewReader,
	"bufio.NewReaderSize":            bufio.NewReaderSize,
	"bufio.NewScanner":               bufio.NewScanner,
	"bufio.NewWriter":                bufio.NewWriter,
	"bufio.NewWriterSize":            bufio.NewWriterSize,
	"bufio.ScanBytes":                bufio.ScanBytes,
	"bufio.ScanLines":                bufio.ScanLines,
	"bufio.ScanRunes":                bufio.ScanRunes,
	"bufio.ScanWords":                bufio.ScanWords,
}

var typList = []interface{}{
	(*bufio.ReadWriter)(nil),
	(*bufio.Reader)(nil),
	(*bufio.Scanner)(nil),
	(*bufio.SplitFunc)(nil),
	(*bufio.Writer)(nil),
}
