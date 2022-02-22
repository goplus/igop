package main

// Test case where a slice of a user-defined byte type (not uint8 or byte) is
// converted to a string.  Same for slice of runes.

type MyByte byte

type MyRune rune

func main() {
	var y = []MyByte("hello")
	if string(y) != "hello" {
		panic("BUG")
	}

	var z = []MyRune("world")
	if string(z) != "world" {
		panic("BUG")
	}
}
