// export by github.com/goplus/interp/cmd/qexp

package smtp

import (
	"net/smtp"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("net/smtp", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*net/smtp.Client).Auth":               (*smtp.Client).Auth,
	"(*net/smtp.Client).Close":              (*smtp.Client).Close,
	"(*net/smtp.Client).Data":               (*smtp.Client).Data,
	"(*net/smtp.Client).Extension":          (*smtp.Client).Extension,
	"(*net/smtp.Client).Hello":              (*smtp.Client).Hello,
	"(*net/smtp.Client).Mail":               (*smtp.Client).Mail,
	"(*net/smtp.Client).Noop":               (*smtp.Client).Noop,
	"(*net/smtp.Client).Quit":               (*smtp.Client).Quit,
	"(*net/smtp.Client).Rcpt":               (*smtp.Client).Rcpt,
	"(*net/smtp.Client).Reset":              (*smtp.Client).Reset,
	"(*net/smtp.Client).StartTLS":           (*smtp.Client).StartTLS,
	"(*net/smtp.Client).TLSConnectionState": (*smtp.Client).TLSConnectionState,
	"(*net/smtp.Client).Verify":             (*smtp.Client).Verify,
	"(net/smtp.Auth).Next":                  (smtp.Auth).Next,
	"(net/smtp.Auth).Start":                 (smtp.Auth).Start,
	"net/smtp.CRAMMD5Auth":                  smtp.CRAMMD5Auth,
	"net/smtp.Dial":                         smtp.Dial,
	"net/smtp.NewClient":                    smtp.NewClient,
	"net/smtp.PlainAuth":                    smtp.PlainAuth,
	"net/smtp.SendMail":                     smtp.SendMail,
}

var typList = []interface{}{
	(*smtp.Auth)(nil),
	(*smtp.Client)(nil),
	(*smtp.ServerInfo)(nil),
}
