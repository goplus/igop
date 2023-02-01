//go:build !go1.20
// +build !go1.20

package igop

const (
	errDeclaredNotUsed     = "declared but not used"
	errImportedNotUsed     = "imported but not used"
	errAppendOutOfRange    = "growslice: cap out of range"
	errSliceToArrayPointer = "cannot convert slice with length %v to pointer to array with length %v"
)
