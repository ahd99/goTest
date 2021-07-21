package main

import "fmt"

func main() {
	test1()
	fmt.Println("--------------- test2 -----")
	test2()
}


/*

- function are first-class values in go. they have types, may assigned to variables, or passed or returnes from functions.
- function type include order and type of parameters and returns (not their names)
- zero value of a function type is nil. call a nil function cause Panic.

- named functions can only written in package level
- unnamed function can be defined insude any expression

*/

func test1() {

	// func variable declaration
	var f func(int) int	// f == nil
	fmt.Printf("%T\t%v\n", f, f == nil)	// "func(int) int	true"
	// n := f(1)	// Panic: because f == nil
	f = square
	fmt.Println(f(3))
	f = negative
	fmt.Println(f(7))
	// f = product	//ERROR: cannot use product (type func(int, int) int) as type func(int) int in assignment

	f1 := square	// type of f is: func(int) int
	fmt.Printf("%T\n", f1)	// "func(int) int"
	n1 := f1(2)  //equal to call square directly: n1 := square(2)
	fmt.Println(n1)	// "4"

	f1 = negative  // because type of negative is eqaul to type of f : 'func(int) int'
	n1 = f1(4)	//"-4"

	//f = product	//Compile error: cannot use product (type func(int, int) int) as type func(int) int in assignment 

	f2 := getFunction1()
	fmt.Printf("%T\n", f2)	// "func(int, int) int"
	n2 := f2(3,4)
	fmt.Println(n2)	// "12"
	//f1 = swap	//ERROR: type of swap is 'func(int, int) int, int' and its return types different from f1 that is 'func(int, int) int'

	//function values are not comparable. the only valid comprison is againt nil
	if(f == nil) {
		fmt.Println("f is nil")
	}


	fmt.Println(applyCalculation(5, square))	// "25"
	n3 := negative
	fmt.Println(applyCalculation(5, n3))	// "-5"
	//fmt.Println(applyCalculation(5, swap))	// ERROR : fmt.Println(applyCalculation(5, n3))	// "-5"

}

func test2() {
	// function literal or anonymous functions:
	f := func(n int) int {
		return n + 1
	}
	fmt.Printf("%T\n", f)	// "func(int) int"
	fmt.Println(f(2))	// "3"

	// function value as return value
	f1 := getFunction2()
	fmt.Println(f1(2))	// "1"

	// pass function value as argument
	n := applyCalculation(10, func(m int) int {return m * 2})
	fmt.Println(n)	// "20"

	// write and run an anonymous function 
	func(s string) {
		fmt.Printf("Hello %s", s)
	}("ali")
}

// functions can return function value
func getFunction1() func(int, int) int {
	return product
	//return swap	//ERROR because type of swap is 'func(int, int) int, int' and its return types different with getFunction1() return type
}

// create and return an anonymous function as return value
func getFunction2() func(int) int {
	return func(n int) int {
		return n - 1
	}
}

// function values can pass as parameter to functions
func applyCalculation(n int, calcFunc func(int) int) int {
	return calcFunc(n)
}

func square(n int) int {
	return n*n
}

func negative(n int) int {
	return -n
}

func product(n,m int) int {
	return n*m
}

func swap(n,m int) (int, int) {
	return m, n
}

