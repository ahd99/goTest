package main

import (
	"fmt"
)

func test2() {
	forTest()
}

func forTest() {
	var i int = 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

}
