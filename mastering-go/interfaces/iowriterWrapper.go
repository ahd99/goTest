package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	test1()
	fmt.Println("---------------- test2 ------")
	test2()
}

// counter Writer

// implements iio.Writer and fmt.Stringer interfaces
type CounterWriter struct {
	writer io.Writer
	Count *int
}

// implment Write method of io.Writer interface
func (w CounterWriter) Write(p []byte) (n int, err error) {
	*(w.Count) += len(p)
	fmt.Println("cwr", w)
	return w.writer.Write(p)
}

// implement String method of fmt.Stringer interface`
func (w CounterWriter) String() string{
	return fmt.Sprintf("CounterWriter: %d", *w.Count)
}

func test1() {
	cwr := CounterWriter{os.Stdout, new(int)}
	fmt.Fprintf(cwr, "hello, %s\n", "ali")
	fmt.Println("test1", cwr)
	fmt.Fprintf(cwr, "hello, %s\n", "ali")
}


//--------------------------- Limit reader

type LimitReader struct {
	reader 	io.Reader
	limit 	int
	i		int
}

func (l *LimitReader) Read(p []byte) (int, error) {
	remained := l.limit - l.i
	readSize := len(p)
	if readSize > remained {
		readSize = remained
	}
	if readSize == 0 {
		return 0, io.EOF
	}
	n ,err := l.reader.Read(p[:readSize])
	l.i += n
	if (l.i == l.limit && err == nil) {
		err = io.EOF
	}
	return n, err
}

func NewLimitReader(r io.Reader, l int) *LimitReader {
	lReader := LimitReader{r, l, 0}
	return &lReader
}

func test2() {
	s := "1234567890"
	limitReader := NewLimitReader(strings.NewReader(s), 6)
	fmt.Printf("limitReader: %T\t%[1]v\n", limitReader)
	p := make([]byte, 4)
	n ,err := limitReader.Read(p)
	fmt.Println(n ,err)
	n ,err = limitReader.Read(p)
	fmt.Println(n ,err)
	n ,err = limitReader.Read(p)
	fmt.Println(n ,err)

}