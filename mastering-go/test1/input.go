package main

import (
	"bufio"
	"io"
	"os"
)

func main() {
	//f := os.Stdin
	var f *os.File
	f = os.Stdin
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		io.WriteString(os.Stdout, scanner.Text())
		io.WriteString(os.Stdout, "\n")

		io.WriteString(os.Stderr, "My app Error output !!\n")

	}
}

// go run myapp.go 2>/tmp/myErrFile
// go run input.go > out.tmp 2>&1
