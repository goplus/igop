package main

import (
	"embed"
	"fmt"
	"reflect"
)

//go:embed testdata/data1.txt
var data1 string

var (
	//go:embed testdata/data2.txt
	data2 []byte
	//go:embed testdata/*
	fs embed.FS
)

var (
	//go:embed "testdata/data1.txt"
	helloT []T
	//go:embed "testdata/data1.txt"
	helloUint8 []uint8
	//go:embed "testdata/data1.txt"
	helloEUint8 []EmbedUint8
	//go:embed "testdata/data1.txt"
	helloBytes EmbedBytes
	//go:embed "testdata/data1.txt"
	helloString EmbedString
)

type T byte
type EmbedUint8 uint8
type EmbedBytes []byte
type EmbedString string

func checkAliases() {
	want := []byte("hello data1")
	check := func(g interface{}) {
		got := reflect.ValueOf(g)
		for i := 0; i < got.Len(); i++ {
			if byte(got.Index(i).Uint()) != want[i] {
				panic(fmt.Errorf("got %v want %v", got.Bytes(), want))
			}
		}
	}
	check(helloT)
	check(helloUint8)
	check(helloEUint8)
	check(helloBytes)
	check(helloString)
}

func checkEmbed() {
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

func main() {
	checkEmbed()
	checkAliases()
}
