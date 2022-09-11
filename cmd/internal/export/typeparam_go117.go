//go:build !go1.18
// +build !go1.18

package export

import "go/types"

func hasTypeParam(t types.Type) bool {
	return false
}
