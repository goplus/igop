// export by github.com/goplus/interp/cmd/qexp

package mail

import (
	"net/mail"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("net/mail", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*net/mail.Address).String":          (*mail.Address).String,
	"(*net/mail.AddressParser).Parse":     (*mail.AddressParser).Parse,
	"(*net/mail.AddressParser).ParseList": (*mail.AddressParser).ParseList,
	"(net/mail.Header).AddressList":       (mail.Header).AddressList,
	"(net/mail.Header).Date":              (mail.Header).Date,
	"(net/mail.Header).Get":               (mail.Header).Get,
	"net/mail.ErrHeaderNotPresent":        &mail.ErrHeaderNotPresent,
	"net/mail.ParseAddress":               mail.ParseAddress,
	"net/mail.ParseAddressList":           mail.ParseAddressList,
	"net/mail.ParseDate":                  mail.ParseDate,
	"net/mail.ReadMessage":                mail.ReadMessage,
}

var typList = []interface{}{
	(*mail.Address)(nil),
	(*mail.AddressParser)(nil),
	(*mail.Header)(nil),
	(*mail.Message)(nil),
}
