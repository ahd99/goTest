package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	test1()
	fmt.Println("----------------------------------- test2")
	test2()
	fmt.Println("----------------------------------- test3")
	test3()
}

// --------------------------------------------- implement with struct type
type CounterWriter struct {
	writer io.Writer
	count  int
}

func (c *CounterWriter) Write(p []byte) (int, error) {
	n, err := c.writer.Write(p)
	c.count += n
	return n, err
}

func (c CounterWriter) String() string {
	return fmt.Sprintf("CounterWriter: %d", c.count)
}

func NewCounterWriter(writer io.Writer) (io.Writer, *int) {
	cw := CounterWriter{writer, 0}
	return &cw, &cw.count
}

func test1() {
	cw := CounterWriter{os.Stdout, 0}
	fmt.Fprintf(&cw, "%s\n", "Hello")
	fmt.Fprintf(&cw, "%s\n", "ali")
	fmt.Println(cw)
}

func test2() {
	cw, c := NewCounterWriter(os.Stdout)
	fmt.Fprintf(cw, "%s\n", "Hello")
	fmt.Fprintf(cw, "%s\n", "ali")
	fmt.Println(cw, *c)
}

// ------------------------------------------- implelent with non struct type

type ByteCounter int

func (b *ByteCounter) Write(p []byte) (int, error) {
	*b += ByteCounter(len(p))
	return len(p), nil
}

func (b ByteCounter) String() string {
	return fmt.Sprintf("ByteCounter: %d", b)
}

func test3() {
	bc := ByteCounter(0)
	fmt.Fprintf(&bc, "%s\n", "Hello")
	fmt.Fprintf(&bc, "%s\n", "ali")
	bc.Write([]byte("!!"))
	fmt.Println(bc)
}
