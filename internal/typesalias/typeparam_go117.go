//go:build !go1.18
// +build !go1.18

package typesalias

func HasTypeParam(typ types.Type) bool {
	return false
}
