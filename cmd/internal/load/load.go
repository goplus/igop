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

package load

import (
	"os"
	"path/filepath"
)

// IsDir checks a target path is dir or not.
func IsDir(target string) (bool, error) {
	fi, err := os.Stat(target)
	if err != nil {
		return false, err
	}
	return fi.IsDir(), nil
}

func ContainsExt(srcDir string, exts ...string) bool {
	if f, err := os.Open(srcDir); err == nil {
		defer f.Close()
		fis, _ := f.Readdir(-1)
		for _, fi := range fis {
			if !fi.IsDir() {
				ext := filepath.Ext(fi.Name())
				for _, v := range exts {
					if v == ext {
						return true
					}
				}
			}
		}
	}
	return false
}
