//go:build go1.16
// +build go1.16

package igop_test

import (
	"testing"

	"github.com/goplus/igop"
	_ "github.com/goplus/igop/pkg/embed"
)

func TestEmbed(t *testing.T) {
	_, err := igop.Run("./testdata/embed", nil, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestEmbedErrorNoMatching(t *testing.T) {
	src := `package main

import (
	_ "embed"
)

//go:embed testdata/notfound.txt
var data string

func main() {
}
`
	_, err := igop.RunFile("main.go", src, nil, 0)
	if err == nil {
		t.Fatal("must panic")
	}
	t.Log(err)
}

func TestEmbedErrorMultipleVars(t *testing.T) {
	src := `package main

import (
	_ "embed"
)

//go:embed testdata
var data1, data2 string

func main() {
}
`
	_, err := igop.RunFile("main.go", src, nil, 0)
	if err == nil {
		t.Fatal("must panic")
	}
	t.Log(err)
}

func TestEmbedErrorMisplaced(t *testing.T) {
	src := `package main

import (
	_ "embed"
)

//go:embed testdata
//var data1 string

func main() {
}
`
	_, err := igop.RunFile("main.go", src, nil, 0)
	if err == nil {
		t.Fatal("must panic")
	}
	t.Log(err)
}

func TestEmbedErrorCannotApply(t *testing.T) {
	src := `package main

import (
	_ "embed"
)

//go:embed testdata
var data1 [][]byte

func main() {
}
`
	_, err := igop.RunFile("main.go", src, nil, 0)
	if err == nil {
		t.Fatal("must panic")
	}
	t.Log(err)
}
