// export by github.com/goplus/gossa/cmd/qexp

package user

import (
	"os/user"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("os/user", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*os/user.User).GroupIds":            (*user.User).GroupIds,
	"(os/user.UnknownGroupError).Error":   (user.UnknownGroupError).Error,
	"(os/user.UnknownGroupIdError).Error": (user.UnknownGroupIdError).Error,
	"(os/user.UnknownUserError).Error":    (user.UnknownUserError).Error,
	"(os/user.UnknownUserIdError).Error":  (user.UnknownUserIdError).Error,
	"os/user.Current":                     user.Current,
	"os/user.Lookup":                      user.Lookup,
	"os/user.LookupGroup":                 user.LookupGroup,
	"os/user.LookupGroupId":               user.LookupGroupId,
	"os/user.LookupId":                    user.LookupId,
}

var typList = []interface{}{
	(*user.Group)(nil),
	(*user.UnknownGroupError)(nil),
	(*user.UnknownGroupIdError)(nil),
	(*user.UnknownUserError)(nil),
	(*user.UnknownUserIdError)(nil),
	(*user.User)(nil),
}
