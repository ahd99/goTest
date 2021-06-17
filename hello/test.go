package main

import (
	"fmt"
	"math"
)

func test1() {
	fmt.Println(math.Pi)
	fmt.Println(swap("A", "B"))
	a, b := swap("ali", "baba")
	fmt.Println(a, b)
}

func swap(a, b string) (x string, y string) {
	x = b
	y = a
	return
}
