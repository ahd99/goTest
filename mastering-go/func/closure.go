package main

import "fmt"

func main() {
	test1()
	fmt.Println("------------ test2 -------")
	test2()
}

var sequenceRatio int = 2

func test1() {
	f1 := getEvenSeries()
	fmt.Println("f1 ->  ", f1())	// "2"
	fmt.Println("f1 ->  ", f1())	// "4"
	fmt.Println("f1 ->  ", f1())	// "6"

	f2 := getEvenSeries()
	fmt.Println("f2 ->  ", f2())	// "2"
	fmt.Println("f2 ->  ", f2())	// "4"

	sequenceRatio = 20

	fmt.Println("f1 ->  ", f1())	// "80"
	fmt.Println("f2 ->  ", f2())	// "60"
}


// anonymous function have access to entire lixical block, so they have access to local variable of enclosing function.
// a hidden reference remained to local variable i from f1 and f2 funcation variables. each one of f1 and f2 have reference to their own instance of i and not to same variable 

func getEvenSeries() func() int {
	i := 0
	return func() int {
		i++
		return i * sequenceRatio
	}
}


func test2() {
	numbers := []int {10, 20, 30, 40, 50}
	var funcs  []func() int
	for i:=0; i < len(numbers); i++ {
		ii := i	// necessary, because i is not belong to internal for block and after finish loop, all funcs refer to last i value that is 5
		funcs = append(funcs, func() int{
			return numbers[ii]
		})
	}

	for _, f := range funcs {
		fmt.Println(f())
	}
}