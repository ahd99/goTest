package main

import "fmt"

func main() {
	test1()
}

func test1() {
	// add function can be casted to arfunc because its signature is equal to arfunc type signature: func(int, int) int
	arf := arfunc(add)
	// arf is of type arfunc and arfunc type has a method Error(). so independent of arfunc is a struct or function, calling
	// String on it, calls its Error() method
	s1 := arf.Error() // s1 == "arfunc.String"
	fmt.Println(s1)

	var err error
	// it is ok to assign arf to err because arf implemented error interface
	err = arf
	fmt.Println(err)
}

type arfunc func(int, int) int

func (f arfunc) Error() string {
	return "arfunc.String"
}

func (f arfunc) multiply(x int, y int) int {
	return x * y
}

func add(x int, y int) int {
	return x + y
}

func subtract(x int, y int) int {
	return x - y
}
