//go:build go1.20
// +build go1.20

package igop

const (
	errDeclaredNotUsed     = "declared and not used"
	errImportedNotUsed     = "imported and not used"
	errAppendOutOfRange    = "len out of range"
	errSliceToArrayPointer = "cannot convert slice with length %v to array or pointer to array with length %v"
)
