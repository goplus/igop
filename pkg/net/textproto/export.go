// export by github.com/goplus/gossa/cmd/qexp

package textproto

import (
	"net/textproto"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("net/textproto", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*net/textproto.Conn).Close":                    (*textproto.Conn).Close,
	"(*net/textproto.Conn).Cmd":                      (*textproto.Conn).Cmd,
	"(*net/textproto.Conn).DotReader":                (*textproto.Conn).DotReader,
	"(*net/textproto.Conn).DotWriter":                (*textproto.Conn).DotWriter,
	"(*net/textproto.Conn).EndRequest":               (*textproto.Conn).EndRequest,
	"(*net/textproto.Conn).EndResponse":              (*textproto.Conn).EndResponse,
	"(*net/textproto.Conn).Next":                     (*textproto.Conn).Next,
	"(*net/textproto.Conn).PrintfLine":               (*textproto.Conn).PrintfLine,
	"(*net/textproto.Conn).ReadCodeLine":             (*textproto.Conn).ReadCodeLine,
	"(*net/textproto.Conn).ReadContinuedLine":        (*textproto.Conn).ReadContinuedLine,
	"(*net/textproto.Conn).ReadContinuedLineBytes":   (*textproto.Conn).ReadContinuedLineBytes,
	"(*net/textproto.Conn).ReadDotBytes":             (*textproto.Conn).ReadDotBytes,
	"(*net/textproto.Conn).ReadDotLines":             (*textproto.Conn).ReadDotLines,
	"(*net/textproto.Conn).ReadLine":                 (*textproto.Conn).ReadLine,
	"(*net/textproto.Conn).ReadLineBytes":            (*textproto.Conn).ReadLineBytes,
	"(*net/textproto.Conn).ReadMIMEHeader":           (*textproto.Conn).ReadMIMEHeader,
	"(*net/textproto.Conn).ReadResponse":             (*textproto.Conn).ReadResponse,
	"(*net/textproto.Conn).StartRequest":             (*textproto.Conn).StartRequest,
	"(*net/textproto.Conn).StartResponse":            (*textproto.Conn).StartResponse,
	"(*net/textproto.Error).Error":                   (*textproto.Error).Error,
	"(*net/textproto.Pipeline).EndRequest":           (*textproto.Pipeline).EndRequest,
	"(*net/textproto.Pipeline).EndResponse":          (*textproto.Pipeline).EndResponse,
	"(*net/textproto.Pipeline).Next":                 (*textproto.Pipeline).Next,
	"(*net/textproto.Pipeline).StartRequest":         (*textproto.Pipeline).StartRequest,
	"(*net/textproto.Pipeline).StartResponse":        (*textproto.Pipeline).StartResponse,
	"(*net/textproto.Reader).DotReader":              (*textproto.Reader).DotReader,
	"(*net/textproto.Reader).ReadCodeLine":           (*textproto.Reader).ReadCodeLine,
	"(*net/textproto.Reader).ReadContinuedLine":      (*textproto.Reader).ReadContinuedLine,
	"(*net/textproto.Reader).ReadContinuedLineBytes": (*textproto.Reader).ReadContinuedLineBytes,
	"(*net/textproto.Reader).ReadDotBytes":           (*textproto.Reader).ReadDotBytes,
	"(*net/textproto.Reader).ReadDotLines":           (*textproto.Reader).ReadDotLines,
	"(*net/textproto.Reader).ReadLine":               (*textproto.Reader).ReadLine,
	"(*net/textproto.Reader).ReadLineBytes":          (*textproto.Reader).ReadLineBytes,
	"(*net/textproto.Reader).ReadMIMEHeader":         (*textproto.Reader).ReadMIMEHeader,
	"(*net/textproto.Reader).ReadResponse":           (*textproto.Reader).ReadResponse,
	"(*net/textproto.Writer).DotWriter":              (*textproto.Writer).DotWriter,
	"(*net/textproto.Writer).PrintfLine":             (*textproto.Writer).PrintfLine,
	"(net/textproto.MIMEHeader).Add":                 (textproto.MIMEHeader).Add,
	"(net/textproto.MIMEHeader).Del":                 (textproto.MIMEHeader).Del,
	"(net/textproto.MIMEHeader).Get":                 (textproto.MIMEHeader).Get,
	"(net/textproto.MIMEHeader).Set":                 (textproto.MIMEHeader).Set,
	"(net/textproto.MIMEHeader).Values":              (textproto.MIMEHeader).Values,
	"(net/textproto.ProtocolError).Error":            (textproto.ProtocolError).Error,
	"net/textproto.CanonicalMIMEHeaderKey":           textproto.CanonicalMIMEHeaderKey,
	"net/textproto.Dial":                             textproto.Dial,
	"net/textproto.NewConn":                          textproto.NewConn,
	"net/textproto.NewReader":                        textproto.NewReader,
	"net/textproto.NewWriter":                        textproto.NewWriter,
	"net/textproto.TrimBytes":                        textproto.TrimBytes,
	"net/textproto.TrimString":                       textproto.TrimString,
}

var typList = []interface{}{
	(*textproto.Conn)(nil),
	(*textproto.Error)(nil),
	(*textproto.MIMEHeader)(nil),
	(*textproto.Pipeline)(nil),
	(*textproto.ProtocolError)(nil),
	(*textproto.Reader)(nil),
	(*textproto.Writer)(nil),
}
