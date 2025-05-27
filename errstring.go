//go:build !go1.20
// +build !go1.20

/*
 * Copyright (c) 2022 The GoPlus Authors (goplus.org). All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ixgo

import "strings"

const (
	errDeclaredNotUsed     = "declared but not used"
	errImportedNotUsed     = "imported but not used"
	errAppendOutOfRange    = "growslice: cap out of range"
	errSliceToArrayPointer = "cannot convert slice with length %v to pointer to array with length %v"
)

func hasTypesNotUsedError(msg string) bool {
	return strings.HasSuffix(msg, errDeclaredNotUsed) || strings.HasSuffix(msg, errImportedNotUsed)
}

func isTypesDeclaredNotUsed(msg string) (string, bool) {
	if strings.HasSuffix(msg, errDeclaredNotUsed) {
		return msg[0 : len(msg)-len(errDeclaredNotUsed)-1], true
	}
	return "", false
}
