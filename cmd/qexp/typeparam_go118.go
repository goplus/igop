//go:build go1.18
// +build go1.18

package main

import (
	"go/types"
)

func IsTypeParam(t types.Type) bool {
	switch t := t.(type) {
	case *types.TypeParam:
		return true
	case *types.Named:
		return t.TypeParams() != nil
	case *types.Signature:
		return t.RecvTypeParams() != nil || t.TypeParams() != nil
	}
	return false
}
