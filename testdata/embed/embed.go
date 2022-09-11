package main

import (
	"embed"
)

//go:embed testdata/data1.txt
var data1 string

var (
	//go:embed testdata/data2.txt
	data2 []byte
	//go:embed testdata/*
	fs embed.FS
)

func main() {
	if data1 != "hello data1" {
		panic(data1)
	}
	if string(data2) != "hello data2" {
		panic(data2)
	}
	data, err := fs.ReadFile("testdata/sub/data2.txt")
	if err != nil {
		panic(err)
	}
	if string(data) != "sub data2" {
		panic(data)
	}
}
