package main

import (
	"bytes"
	"fmt"
	"log"
	"time"
	"unicode/utf8"
)

func main() {
	//work()
	//test3()
	//test4()
	// test5()
	// test6()
	test7()
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

	s1 := "aعلی\xbd\u063Ab"
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

func test4() {
	
	// s := "aΣΔ\xbd\u03A9b"
	// fmt.Printf("%v\n", s)	// aΣΔ�Ωb"
	// fmt.Printf("%s\n", s)	// aΣΔ�Ωb"
	// fmt.Printf("% x\n", s)	// 61cea3ce94bdcea962"
	// fmt.Printf("%v\n", s)	// aΣΔ�Ωb"
	// fmt.Printf("%q\n", s)	// "aΣΔ\xbdΩb"
	// fmt.Printf("%+q\n", s)	// "a\u03a3\u0394\xbd\u03a9b"

	//s := "aΣΔ\xbd\u03A9b"
	// fmt.Printf("%s\n", s)	// aΣΔ�Ωb"
	// fmt.Printf("% x\n", s)
	// fmt.Printf("%+q\n", s)	// "a\u03a3\u0394\xbd\u03a9b"
	// for i, r := range s {
	// 	fmt.Printf("%d\t%#U\n", i, r)
	// }

	// s := "aΣΔ\xbd\u03A9b"
	// fmt.Printf("%s\n", s)	// aΣΔ�Ωb"
	// fmt.Printf("% x\n", s)
	// r := []rune(s)
	// fmt.Printf("%T\t%[1]v\t%[1]x\t%[1]q\n", r)

	// return~
	// for i, r := range s {
	// 	fmt.Printf("%d\t%#U\n", i, r)
	// }

	// fmt.Println("============================")
	

	s := "ع"
	fmt.Printf("%s\n", s)	// aΣΔ�Ωb"
	fmt.Printf("% x\n", s)
	for i, w := 0, 0; i < len(s); i += w {
		r1, width := utf8.DecodeRuneInString(s[i:])
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

func test5() {
	s := "ali heydari"
	prefix := "ali h"
	wrongPrefix := "alih"


	if !IsPrefix(s, prefix) {
		log.Fatalf("Error in isPrefix result. s:%s prefix:%s", s, prefix)
	}
	

	if IsPrefix(s, wrongPrefix) {
		log.Fatalf("Error in isPrefix result. s:%s prefix:%s", s, prefix)
	}

	log.Println("isPrefix success")
}

func test6() {
	s := "ali heydari"
	substr := "li h"
	wrongSubstr := "lih"

	if !IsContain(s, substr) {
		log.Fatalf("Error in IsCOntain result. s:%s substr:%s", s, substr)
	}
	

	if IsContain(s, wrongSubstr) {
		log.Fatalf("Error in IsCOntain result. s:%s substr:%s", s, wrongSubstr)
	}

	log.Println("IsContain success")

	fmt.Println(string(0x639))
}

func test7() {
	var buf bytes.Buffer

	buf.WriteByte('a')
	buf.WriteString("li")
	buf.WriteRune(0x0639)
	fmt.Fprintf(&buf, "%d", 66)
	s := buf.String()
	fmt.Println(s)	// "ali66"
	
	s1 := fmt.Sprintf("%d %[1]b %[1]x", 123)
	fmt.Println(s1)
}

func IsPrefix( s string, prefix string) bool {
	return (len(s) >= len(prefix) && (s[:len(prefix)] == prefix))
}

func IsContain(s string, substr string) bool{
	for i:=0; i<len(s)-len(substr); i++ {
		if IsPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}