//go:build ignore
// +build ignore

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

package main

import (
	"log"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	ver := runtime.Version()[:6]
	var tags string
	var name string
	var fname string
	switch ver {
	case "go1.20":
		tags = "//+build go1.20"
		name = "go120_export"
		fname = "go120_pkgs.go"
	case "go1.19":
		tags = "//+build go1.19,!go1.20"
		name = "go119_export"
		fname = "go119_pkgs.go"
	case "go1.18":
		tags = "//+build go1.18,!go1.19"
		name = "go118_export"
		fname = "go118_pkgs.go"
	case "go1.17":
		tags = "//+build go1.17,!go1.18"
		name = "go117_export"
		fname = "go117_pkgs.go"
	case "go1.16":
		tags = "//+build go1.16,!go1.17"
		name = "go116_export"
		fname = "go116_pkgs.go"
	case "go1.15":
		tags = "//+build go1.15,!go1.16"
		name = "go115_export"
		fname = "go115_pkgs.go"
	case "go1.14":
		tags = "//+build go1.14,!go1.15"
		name = "go114_export"
		fname = "go114_pkgs.go"
	}

	log.Println(ver, name, fname, tags)

	// syscall/js
	cmd := exec.Command("qexp", "-outdir", ".", "-addtags", tags+";//+build js", "-filename", name, "-contexts", "js", "syscall/js")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "GOOS=js")
	cmd.Env = append(cmd.Env, "GOARCH=wasm")
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
