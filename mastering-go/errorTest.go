package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("No argument error")
		return
	}

	var count int = 0
	for i := 1; i < len(os.Args); i++ {
		n, err := checkNumber(os.Args[i])
		if err != nil { // golang way of error handling: err != nil means there is error
			// there is error
			log.Println(err)
		} else {
			log.Println(n)
			count++
		}
	}
}

// check numbers !!
func checkNumber(numstr string) (float64, error) {
	n, err := strconv.ParseFloat(numstr, 64)
	if err != nil {
		// return 0, MyError{err.Error(), 1} return my own implementation of golang error interface
		return 0, errors.New("bad number " + numstr) //return golang default implementation
	}
	return n, nil //returning nil for error means no error
}

// our own custom implementation of golang error interface
type MyError struct {
	s string
	n int
}

func (e MyError) Error() string {
	return e.s + "(" + fmt.Sprintf("%d", e.n) + ")"
}
