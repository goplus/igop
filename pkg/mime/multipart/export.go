// export by github.com/goplus/interp/cmd/qexp

package multipart

import (
	"mime/multipart"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("mime/multipart", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*mime/multipart.FileHeader).Open":            (*multipart.FileHeader).Open,
	"(*mime/multipart.Form).RemoveAll":             (*multipart.Form).RemoveAll,
	"(*mime/multipart.Part).Close":                 (*multipart.Part).Close,
	"(*mime/multipart.Part).FileName":              (*multipart.Part).FileName,
	"(*mime/multipart.Part).FormName":              (*multipart.Part).FormName,
	"(*mime/multipart.Part).Read":                  (*multipart.Part).Read,
	"(*mime/multipart.Reader).NextPart":            (*multipart.Reader).NextPart,
	"(*mime/multipart.Reader).NextRawPart":         (*multipart.Reader).NextRawPart,
	"(*mime/multipart.Reader).ReadForm":            (*multipart.Reader).ReadForm,
	"(*mime/multipart.Writer).Boundary":            (*multipart.Writer).Boundary,
	"(*mime/multipart.Writer).Close":               (*multipart.Writer).Close,
	"(*mime/multipart.Writer).CreateFormField":     (*multipart.Writer).CreateFormField,
	"(*mime/multipart.Writer).CreateFormFile":      (*multipart.Writer).CreateFormFile,
	"(*mime/multipart.Writer).CreatePart":          (*multipart.Writer).CreatePart,
	"(*mime/multipart.Writer).FormDataContentType": (*multipart.Writer).FormDataContentType,
	"(*mime/multipart.Writer).SetBoundary":         (*multipart.Writer).SetBoundary,
	"(*mime/multipart.Writer).WriteField":          (*multipart.Writer).WriteField,
	"(mime/multipart.File).Close":                  (multipart.File).Close,
	"(mime/multipart.File).Read":                   (multipart.File).Read,
	"(mime/multipart.File).ReadAt":                 (multipart.File).ReadAt,
	"(mime/multipart.File).Seek":                   (multipart.File).Seek,
	"mime/multipart.ErrMessageTooLarge":            &multipart.ErrMessageTooLarge,
	"mime/multipart.NewReader":                     multipart.NewReader,
	"mime/multipart.NewWriter":                     multipart.NewWriter,
}

var typList = []interface{}{
	(*multipart.File)(nil),
	(*multipart.FileHeader)(nil),
	(*multipart.Form)(nil),
	(*multipart.Part)(nil),
	(*multipart.Reader)(nil),
	(*multipart.Writer)(nil),
}
