package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "我很好。你呢？"
	fmt.Printf("%s\n", []byte(s))
	for _, b := range s {
		fmt.Printf("%X ", b)
	}
	fmt.Println()
	for i, ch := range s {
		fmt.Printf("(%d %X) ", i, ch)
	}
	fmt.Println()

	fmt.Println("Rune Count:", utf8.RuneCountInString(s))

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		fmt.Printf("%c ", ch)
		bytes = bytes[size:]
	}
	fmt.Println()
}
