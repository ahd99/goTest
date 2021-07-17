package main

import "fmt"

func main() {
	test1()
	// test2()

}

func test1() {
	var a [3]int // [0 0 0] define arrya of int with3 elements from a[0] to a[2]. all elements get zero-value

	a[0] = 2 // [2 0 0 ]

	for i, v := range a { // i: index(0..2) v:value
		fmt.Printf("%d  %d\n", i, v)
	}

	for i := range a {
		fmt.Println(i) //print indexes
	}

	for _, v := range a {
		fmt.Println(v) //print vlaues
	}

	a[1] = 3                                     // change an element value in the array
	a = [3]int{9, 8, 7}                          // "[9 8 7]"  -> assig a new array to variable(not only change element value)
	fmt.Printf("%T  %v  len:%d\n", a, a, len(a)) //[3]int  [9 8 7]  len:3
	//a = [4]int{1, 2, 3, 4} //compile error. bcause type [4]int is different from [3]int (type a)

	var b [4]int = [4]int{1, 2, 3} // [1 2 3 0] declare array and initialize first 3 elements. b[3] will be initialized with zero-value:
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
