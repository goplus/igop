//go:build !go1.20
// +build !go1.20

package igop

const (
	errDeclNotUsed      = "declared but not used"
	errAppendOutOfRange = "growslice: cap out of range"
	errSliceToArrayPointer     = "cannot convert slice with length %v to pointer to array with length %v"
)
