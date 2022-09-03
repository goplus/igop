//go:build !go1.18
// +build !go1.18

package main

func IsTypeParam(t types.Type) bool {
	return false
}
