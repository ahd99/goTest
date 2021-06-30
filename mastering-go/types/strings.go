package main

import (
	"fmt"
	"time"
	"unicode/utf8"
)

func main() {
	//work()
	test3()
}

func test1() {
	s := "عل"
	fmt.Printf("% x\n", s)
	fmt.Printf("%d\n", s[0])

	d1 := rune(s[0])
	fmt.Printf("%T >>  %v \n", d1, d1)

	r := []rune(s)
	fmt.Printf("%d\n", r)
}

func test2() {
	s := "ali\xbdع"
	fmt.Println(s)
	fmt.Printf("% x\n%[1]q\n%+[1]q\n------------\n", s)
	for i := 0; i < len(s); i++ {
		fmt.Printf("% x\t%[1]q\t%+[1]q\n", s[i])
	}
}

func test3() {
	s := "\u0061\u0300"
	fmt.Printf("%T\t%[1]v\t%[1]s\t% [1]x\t%[1]q\t%+[1]q\n", s)

	s1 := "aعلی\xbdb"
	fmt.Printf("%T\t%[1]v\t%[1]s\t% [1]x\t%[1]q\t%+[1]q\n", s1)
	r1 := []rune(s1)
	fmt.Printf("%T\t%[1]v\t%[1]s\t% [1]x\t%[1]q\t%+[1]q\n", r1)

	for i, r := range s1 {
		fmt.Printf("%d\t%#U\n", i, r)
	}

	fmt.Println("============================")
	for i, w := 0, 0; i < len(s1); i += w {
		r1, width := utf8.DecodeRuneInString(s1[i:])
		fmt.Printf("%d\t%d\t%#U\n", i, w, r1)
		w = width
	}
}

func work() {
	//defer time.Since(time.Now())

	//defer t1(t2("ali"))
	//fmt.Println("---------------------")

	// defer func(t time.Time) {
	// 	t2 := time.Since(t)
	// 	fmt.Println(t2)
	// }(time.Now())

	defer stat()()

	time.Sleep(3 * time.Second)
	fmt.Println("after sleep")
}

func t1(s string) string {
	r := "t1 -> " + s
	fmt.Println(r)
	return r
}

func t2(s string) string {
	r := "t2 -> " + s
	fmt.Println(r)
	return r
}

func stat() func() {
	t1 := time.Now()
	return func() {
		dt := time.Since(t1)
		fmt.Println(dt)
	}
}
