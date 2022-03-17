package main

import "fmt"

func main() {
	test1()
	// test2()
	test3()

}

func test1() {
	/*
		- arrays are fixed length sequence of elements of a particular type
		- Array index is started from 0 to len-1
		- Size of an array is part of its type, so [3]int and [4]int are different types.
		- Two arrays are comparable and assignable if both element type and length are the same. 
			So [2]int and [3]int can not be compare or assign to each other
		- If 2 arrays are comparable, we can compare them by == and != operators. 
			2 arrays are equal when the values of all corresponding elements are equal.
	*/

	// define arrya of int with 3 elements from a[0] to a[2]. all elements get zero-value.
	// In array declaration, Len of array must be const (can evaluate at compile time)
	// By default Elements of a new array are set to zero-value.
	var a [3]int // [0 0 0]

	// len() method return number of elements in array
	fmt.Println(len(a)) // "3"

	// Arrays can not be nil and can not be compare with nil (cause compile error). array with zero len is empty array:
	var a0 [0]int            //len(a) == 0	a: []
	fmt.Println(len(a0), a0) // 0 []

	a[0] = 2 // [2 0 0 ]

	for i, v := range a { // i: index(0..2) v:value
		fmt.Printf("%d  %d\n", i, v)
	}

	for i := range a { // i: index(0..2)
		fmt.Println(i) //print indexes
	}

	for _, v := range a { // v:value
		fmt.Println(v) //print vlaues
	}

	a[1] = 3     // change an element value in the array
	a = [3]int{9, 8, 7}    // "[9 8 7]"  -> assig a new array to variable(not only change element value)
	fmt.Printf("%T  %v  len:%d\n", a, a, len(a)) //[3]int  [9 8 7]  len:3
	//a = [4]int{1, 2, 3, 4} //compile error. bcause type [4]int is different from [3]int (type a)

	var b [4]int = [4]int{1, 2, 3} // [1 2 3 0] -> declare array and initialize first 3 elements. b[3] will be initialized with zero-value:
	c := [2]int{1, 2}              // [1 2]

	d := [...]int{1, 2, 3, 4, 5} // [1 2 3 4 5] array length is determined by number of init elements. len(d) = 5
	fmt.Printf("%T \n", d)       //"[5]int"

	e := [...]int{0: 10, 1: 20, 2: 30} //[10 20 30]	initialize array by index.
	f := [...]int{4: 14, 2: 12}        // [0 0 12 0 14] //indices can be in any order. missed indices get zero-value
	fmt.Printf("%T  %v  len:%d\n", f, f, len(f))

	fmt.Println("------------------------------")
	fmt.Printf("%v\n", b, e)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", d)
}

func test2() {
	a := [3]int{10, 20, 30}
	incrementArray(a)
	fmt.Printf("%v\n", a)
}

func incrementArray(a [3]int) {
	for i := range a {
		fmt.Println(i)
		a[i] += 1
	}
}

func test3() {

	/*
		golang function arguments are passed by value and arrays are value type (not reference type).
		So in function calls, array arguments will be copied to function parameters.
		So a change in array elements has no effect on the array in the caller function.
		To solve this, use pointer to array (*[3]int) to send array reference.
	*/

	array := [3]int{1, 2, 3}

	increaseArray_V(array)
	fmt.Printf("%v\n", array)

	increaseArray_P(&array)
	fmt.Printf("%v\n", array)

	zeroArray(&array)
	fmt.Printf("%v\n", array)

	array = [3]int{1, 2, 3}
	zeroArray1(&array)
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
	*b = [3]int{} // change has effect on caller because b points to array sent from caller and *b will change that array value
}

func zeroArray1(b *[3]int) {
	b = &[3]int{} //change dont effect on caller. because now b point to another array not array received in argument
}
