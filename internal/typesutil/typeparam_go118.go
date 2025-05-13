//go:build go1.18 && !go1.23
// +build go1.18,!go1.23

package typesutil

import "go/types"

func HasTypeParam(typ types.Type) bool {
	switch t := typ.(type) {
	case *types.TypeParam:
		return true
	case *types.Named:
		return t.TypeParams() != nil
	case *types.Signature:
		return t.TypeParams() != nil
	}
	return false
}
