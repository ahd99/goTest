package main

import (
	"fmt"
)

func main() {
	test1()
}

const a = 1
const b float64 = 3.14
const (
	x = 1
	y
	z
)

const (
	_  = 1 + iota // 1 + 0
	a1            // 1 + 1
	a2            // 1 + 2
	a3            // 1 + 3
	a4            // 1 + 4
)

func test1() {
	fmt.Println(a1, a2, a3, a4)
}
