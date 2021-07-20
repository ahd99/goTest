package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	f1()
	combineUserdefinedTypes()
	ff()
}

type age int
type weight int
type Title string

func f1() {
	var a1 age = 12
	var w1 weight = 60

	fmt.Printf("a1: %v %T \n", a1, a1)
	fmt.Printf("w1: %v %T \n", w1, w1)

	// a1 = w1	Error

	var i int = 5
	math.Inf(i)
	// math.Inf(a1) Error

	var a2 age
	// a2 = a1 + w1 Error
	_ = a2

	a2 = a1 + 3
}

func combineUserdefinedTypes() {
	var a1 age = 10
	var w1 weight = 50
	var a2 age = 12

	a1 = a2
	a1 = a1 + a2

	a1 = 4
	a1 = a1 + 3

	// a1 = w1			Errot
	a1 = age(w1)
	// a1 = a1 + w1		Error
	a2 = a1 + age(w1)
	// if a1 == w1 {}	Error
	if a1 == age(w1) {
	}

	// func(aa age) {} (w1)		Error
	func(aa age) {}(age(w1))

	var i1 int = 2

	// a1 = i1			Error
	a1 = age(i1)
	//a2 = a1 + i1		Error
	a2 = a1 + age(i1)
	// if a1 == i1 {}	Error
	if a1 == age(i1) {
	}

	// i1 = a1			Error
	i1 = int(a1)
	// if i1 == a1 {}	Error
	if i1 == int(a1) {
	}

	// func(i int) {} (a1)		Error
	func(i int) {}(int(a1))

	// func(a age) {} (i1)		Error
	func(a age) {}(age(i1))

	a1.String1()
	// w1.String1()		ERROR
	// i1.String1()		ERROR

	var t1 Title = "abc"
	var t2 Title = "xyz"

	t2 = t1 + "d"

	_ = a1
	_ = w1
	_ = a2
	_ = i1
	_, _ = t1, t2

	fmt.Println(a1.String1())

}

func (a age) String1() string {
	return "age " + strconv.FormatInt(int64(a), 10)
}

type age1 = age

func aliasTest() {
	var a1 age1 = 10
	var a age = 20

	a = a1
	a1 = a
	var b = a1 + a

	func(g age) {}(a1)
	func(g age1) {}(a)

	_, _ = a, a1
	_ = b
}

func (a1 age1) double() age1 {
	return a1 * a1
}

func ff() {

	i := 1
	j := 2
	i1 := int8(2)
	s := "ali"
	a := age(10)
	w := weight(20)

	//i = i1	Error
	//i1 = i	Error
	//i = s		Error
	i = int(i1)
	i1 = int8(i)

	var ip *int = &i
	var jp *int = &j
	var i1p *int8 = &i1
	var sp *string = &s
	var ap *age = &a
	var wp *weight = &w

	ip = jp
	//ip = i1p		Error
	//ip = ap		Error
	//ap = ip		Error
	//ap = wp		Error

	//ip = (*int)(i1p)	Error
	//ip = (*int)(sp)	Error
	ip = (*int)(ap)
	ap = (*age)(wp)

	fmt.Printf("%T\t%[1]v\n", i)
	fmt.Printf("%T\t%[1]v\n", i1)
	fmt.Printf("%T\t%[1]v\n", s)
	fmt.Printf("%T\t%[1]v\n", a)
	fmt.Printf("%T\t%[1]v\n", w)

	fmt.Printf("%T\t%[1]v\n", *ip)
	fmt.Printf("%T\t%[1]v\n", *i1p)
	fmt.Printf("%T\t%[1]v\n", *sp)
	fmt.Printf("%T\t%[1]v\n", *ap)
	fmt.Printf("%T\t%[1]v\n", *wp)

}
