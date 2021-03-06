// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by genflags.go — DO NOT EDIT.

package test

type Type int

const (
	Bool Type = iota
	Int
	String
	Duration
)

// passFlagToTest contains the flags that should be forwarded to
// the test binary with the prefix "test.".
var passFlagToTest = map[string]Type{
	"bench":                String,
	"benchmem":             Bool,
	"benchtime":            Duration,
	"blockprofile":         String,
	"blockprofilerate":     Int,
	"count":                Int,
	"coverprofile":         String,
	"cpu":                  String,
	"cpuprofile":           String,
	"failfast":             Bool,
	"list":                 String,
	"memprofile":           String,
	"memprofilerate":       Int,
	"mutexprofile":         String,
	"mutexprofilefraction": Int,
	"outputdir":            String,
	"parallel":             Int,
	"run":                  String,
	"short":                Bool,
	"timeout":              Duration,
	"trace":                String,
	"v":                    Bool,
}
