//go:build !go1.18
// +build !go1.18

package export

import "go/types"

func IsTypeParam(t types.Type) bool {
	return false
}
