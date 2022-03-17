package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	readEntireFile()
	openAndReadFile()
}

func readEntireFile() {
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

func openAndReadFile() {
	var f *os.File
	f, err := os.OpenFile("encodings/file.txt", os.O_RDONLY | os.O_CREATE, 0777)
	if err != nil {
		fmt.Println("Error openning file. err:", err)
		os.Exit(1)
	}
	defer f.Close()

	b := make([]byte, 20)

	reader := bufio.NewReader(f)

	n, err := reader.Read(b)
	if err != nil {
		fmt.Println("Error reading from file. err: ", err)
		os.Exit(1)
	}
	fmt.Printf("%d bytes read:\n %s\n", n, string(b[:n]))

}
