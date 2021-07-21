package main

import "fmt"

func main() {
	test1()

}

func test1() {
	printSum("")			// "0"
	printSum("", 1)			// "1"
	printSum("", 1,2,3,4)	// "10"

	// calling variadic mrthod with a slice
	m := []int {1,2,3,4,5,6}
	printSum("", m...)		// "21"
}

// variadic function. can accept any number of arguments
// only last parametrs can be in this way (variadic)
// n is a slice of type int
// type of a variadic function is different from type of a function that get a slice:
//	func f(...int) {}	type is func(...int) but
//	func f([]int) {}	type is func([]int)
func printSum(message string, numbers ...int) {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	fmt.Println(message, sum)
}