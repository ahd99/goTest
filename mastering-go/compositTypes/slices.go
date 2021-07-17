package main

import "fmt"

func main() {
	test1()
	//test2()
}

func test1() {
	// slices are vlaue type containig a len, a cap and a pointer to underlying array (to its start indesx in underlying array)
	var s1 []int // len(s) == 0, s == nil  define a slice with zero value. zero value of slices are nil.
	logSlice_int(s1, "S1")
	s1 = []int{} // len(s) == 0, s != nil
	s1 = nil     // len(s) == 0, s == nil
	// nil slices and non-nil emapty slices behave the same. use len(s) == 0 for checking empty slice, not s == nil

	numbersArray := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} //array
	n0 := numbersArray[2:7]                                // [2 3 4 5 6]  len:5  cap:8  create slice from array. numbersArray is underlying array for slice n0
	logSlice_int(n0, "n0")

	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} //[0 1 2 3 4 5 6 7 8 9]  len:10  cap:10  create slice by slice literal. underlying array has 10 elements
	logSlice_int(numbers, "numbers")

	n1 := numbers[2:7] //[2 3 4 5 6]  len:5  cap:8  create slice from another slice. n1 and numbers share same underlying array
	logSlice_int(n1, "n1")

	n2 := numbers[3:9]
	logSlice_int(n2, "n2")

	n1[1] = 30
	logSlice_int(n1, "n1")
	logSlice_int(n2, "n2")

	fmt.Println(sliceEqualElementsCount(n1, n2))
	logSlice_int(sliceEqualElements(n1, n2))

	//n3 := n1[:9]		// panic
	//logSlice_int(n3)

	fmt.Println("--------------------------------")

	logSlice_int(numbers, "numbers") //[0 1 2 30 4 5 6 7 8 9]  len:10  cap:10
	logSlice_int(n1, "n1")           // [2 30 4 5 6]  len:5  cap:8
	//append an element at the end of slice. if slice has enough capacity (capacity >= len + 1), slice grows on current underlying array,
	// if not (capacity < len + 1), create a new underlying array with more size, copy elements to it and add new element.
	n1 = append(n1, -1)              //  cap(n1) >= len(n1) + 1 so underlying array has enough capacity. n1 slice extend by one on current underlying array and add -1 at last element
	logSlice_int(numbers, "numbers") //[0 1 2 30 4 5 6 -1 8 9]  len:10  cap:10
	logSlice_int(n1, "n1")           //[2 30 4 5 6 -1]  len:6  cap:8
	logSlice_int(n2, "n2")           //[30 4 5 6 -1 8] len:6 cap:7	n2 use same underlying array with n1 so change effects n2

	// append 3 items to n1. because cap(n1) < len(n1) + 3, there is not enough space in underlying array.
	// so a new array allocated with enough space, all data from n1 are copied to new array and 3 new elements are append after copied elements
	// and a new slice wil be returned. new slice len increased by 3. cap increase base on len of new underlying array. usually twice the new slice length.
	// and n1 point to new underlying array. so adding -1, -2, -3 has not effect on numbers and n2 slices that points to previous underlying array.
	// because in append may be new slice returned, always assign append return to the slice itself.
	n1 = append(n1, -2, -3, -4)
	logSlice_int(n1, "n1")           // [2 30 4 5 6 -1 -2 -3 -4]  len:9  cap:16
	logSlice_int(numbers, "numbers") //[0 1 2 30 4 5 6 -1 8 9]  len:10  cap:10
	logSlice_int(n2, "n2")           // [30 4 5 6 -1 8]  len:6  cap:7

	//copy(dest, src). copy elements from n2 to n1. number of copied elements is minimum of two slice lngths and no extend happen.
	numberOfElementsActuallyCopied := copy(n1, n2)
	fmt.Println(numberOfElementsActuallyCopied)

	// append a slice to another by following notation
	n1 = append(n1, n2...)
	n1 = append(n1, []int{-4, -5}...)

	// if(n1 == n2) {	//compile error: slices are not comparable.
	// 	//----
	// }

	if n1 == nil { // only slice comparison is against nil
		fmt.Println("nil")
	}

	fmt.Println(nonEmptyStrings([]string{"", "a", "", "b", "c", ""}))
	fmt.Println(nonEmptyStrings2([]string{"", "a", "", "b", "c", ""}))
}

func test2() {
	len, cap := 10, 20
	a := make([]int, len, cap) //[0 0 0 0 0 0 0 0 0 0]  len:10  cap:20  - create an underlying array of length = cap and create a slice of length = len on it. if remove capacity, capacity will be eqyal to len
	//a = make([]int, cap)[:len] //[0 0 0 0 0 0 0 0 0 0]  len:10  cap:20 - equals to make(int[], len, cap)
	logSlice_int(a)

	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	a1 := numbers[1:5]
	logSlice_int(a1, "a1")

	a1 = append(a1, -1)
	logSlice_int(numbers, "numbers")
	logSlice_int(a1, "a1")

}

func logSlice_int(a []int, desc ...string) {
	if len(desc) < 1 {
		desc = []string{""}
	}
	fmt.Printf("%s > %v  len:%d  cap:%d \n", desc[0], a, len(a), cap(a))
}

func sliceEqualElementsCount(x, y []int) (i int) {
	for a := range x {
		for b := range y {
			if a == b {
				i++
			}
		}
	}
	return
}

func sliceEqualElements(x, y []int) []int {
	s := []int{}
	for _, a := range x {
		for _, b := range y {
			if a == b {
				s = append(s, a)
			}
		}
	}
	return s
}

// remove epty strings in slice by in-place slice change technique.
// at the end new slice is returned because length of slice is changed.
// so caller must get and assign return value: a = nonEmptyStrings(a)
func nonEmptyStrings(s []string) []string {
	i := 0
	for _, v := range s {
		if v != "" {
			s[i] = v
			i++
		}
	}
	return s[:i]
}

func nonEmptyStrings2(s []string) []string {
	out := s[:0] //zero length slice from original
	for _, v := range s {
		if v != "" {
			out = append(out, v)
		}
	}
	return out
}
