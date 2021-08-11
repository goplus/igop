// export by github.com/goplus/interp/cmd/qexp

package xml

import (
	"encoding/xml"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("encoding/xml", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*encoding/xml.Decoder).Decode":                  (*xml.Decoder).Decode,
	"(*encoding/xml.Decoder).DecodeElement":           (*xml.Decoder).DecodeElement,
	"(*encoding/xml.Decoder).InputOffset":             (*xml.Decoder).InputOffset,
	"(*encoding/xml.Decoder).RawToken":                (*xml.Decoder).RawToken,
	"(*encoding/xml.Decoder).Skip":                    (*xml.Decoder).Skip,
	"(*encoding/xml.Decoder).Token":                   (*xml.Decoder).Token,
	"(*encoding/xml.Encoder).Encode":                  (*xml.Encoder).Encode,
	"(*encoding/xml.Encoder).EncodeElement":           (*xml.Encoder).EncodeElement,
	"(*encoding/xml.Encoder).EncodeToken":             (*xml.Encoder).EncodeToken,
	"(*encoding/xml.Encoder).Flush":                   (*xml.Encoder).Flush,
	"(*encoding/xml.Encoder).Indent":                  (*xml.Encoder).Indent,
	"(*encoding/xml.SyntaxError).Error":               (*xml.SyntaxError).Error,
	"(*encoding/xml.TagPathError).Error":              (*xml.TagPathError).Error,
	"(*encoding/xml.UnsupportedTypeError).Error":      (*xml.UnsupportedTypeError).Error,
	"(encoding/xml.CharData).Copy":                    (xml.CharData).Copy,
	"(encoding/xml.Comment).Copy":                     (xml.Comment).Copy,
	"(encoding/xml.Directive).Copy":                   (xml.Directive).Copy,
	"(encoding/xml.Marshaler).MarshalXML":             (xml.Marshaler).MarshalXML,
	"(encoding/xml.MarshalerAttr).MarshalXMLAttr":     (xml.MarshalerAttr).MarshalXMLAttr,
	"(encoding/xml.ProcInst).Copy":                    (xml.ProcInst).Copy,
	"(encoding/xml.StartElement).Copy":                (xml.StartElement).Copy,
	"(encoding/xml.StartElement).End":                 (xml.StartElement).End,
	"(encoding/xml.TokenReader).Token":                (xml.TokenReader).Token,
	"(encoding/xml.UnmarshalError).Error":             (xml.UnmarshalError).Error,
	"(encoding/xml.Unmarshaler).UnmarshalXML":         (xml.Unmarshaler).UnmarshalXML,
	"(encoding/xml.UnmarshalerAttr).UnmarshalXMLAttr": (xml.UnmarshalerAttr).UnmarshalXMLAttr,
	"encoding/xml.CopyToken":                          xml.CopyToken,
	"encoding/xml.Escape":                             xml.Escape,
	"encoding/xml.EscapeText":                         xml.EscapeText,
	"encoding/xml.HTMLAutoClose":                      &xml.HTMLAutoClose,
	"encoding/xml.HTMLEntity":                         &xml.HTMLEntity,
	"encoding/xml.Marshal":                            xml.Marshal,
	"encoding/xml.MarshalIndent":                      xml.MarshalIndent,
	"encoding/xml.NewDecoder":                         xml.NewDecoder,
	"encoding/xml.NewEncoder":                         xml.NewEncoder,
	"encoding/xml.NewTokenDecoder":                    xml.NewTokenDecoder,
	"encoding/xml.Unmarshal":                          xml.Unmarshal,
}

var typList = []interface{}{
	(*xml.Attr)(nil),
	(*xml.CharData)(nil),
	(*xml.Comment)(nil),
	(*xml.Decoder)(nil),
	(*xml.Directive)(nil),
	(*xml.Encoder)(nil),
	(*xml.EndElement)(nil),
	(*xml.Marshaler)(nil),
	(*xml.MarshalerAttr)(nil),
	(*xml.Name)(nil),
	(*xml.ProcInst)(nil),
	(*xml.StartElement)(nil),
	(*xml.SyntaxError)(nil),
	(*xml.TagPathError)(nil),
	(*xml.Token)(nil),
	(*xml.TokenReader)(nil),
	(*xml.UnmarshalError)(nil),
	(*xml.Unmarshaler)(nil),
	(*xml.UnmarshalerAttr)(nil),
	(*xml.UnsupportedTypeError)(nil),
}
