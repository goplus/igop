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
