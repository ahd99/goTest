package main

import (
	"fmt"
	"os"
)

func main() {
	// read file as []byte (byte slice)
	bytes, err := os.ReadFile("encodings/file.txt")
	if err != nil {
		fmt.Println("Error Reading file. err:", err)
		os.Exit(1)
	}

	// print each byte in hex format
	for _, b := range bytes {
		fmt.Printf("% x", b)
	}
	fmt.Printf("\nbytes len: %d\n", len(bytes))

	// convert []byte to string. copy all bytes to string.
	var str string
	str = string(bytes)
	fmt.Println(str)

	fmt.Printf("string  len: %d\n", len(str))

	// convert string to []rune. bytes is interpreted in utf8 encoding.
	var runes []rune
	runes = []rune(str)
	for _, r := range runes {
		fmt.Printf("% x", r)
	}

	fmt.Printf("\nrunes len: %d\n", len(runes))

}
