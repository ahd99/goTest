package main

import "fmt"

func main() {
	test1()
}

func test1() {
	array := [3]int{1,2,3}
	
	increaseArray_V(array)
	fmt.Printf("%v\n", array)

	increaseArray_P(&array)
	fmt.Printf("%v\n", array)

	zeroArray(&array)
	fmt.Printf("%v\n", array)
}

func increaseArray_V(a [3]int) {
	for i := range a {
		a[i] *= -1
	}
}

func increaseArray_P(a *[3]int) {
	for i := range a {
		a[i] *= -1
	}
}

func zeroArray(b *[3]int) {
	*b = [3]int{}
}