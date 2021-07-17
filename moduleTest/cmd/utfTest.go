package main

import "fmt"

func main() {
	f1()
}

func f1() {
	s:= "علی"

	fmt.Printf("% x\n", s)

	for i:=0; i<len(s); i++ {
		fmt.Printf("%x\n", s[i])
	}

	r := []rune(s)

	fmt.Printf("%x\n", r)
	fmt.Printf("%d\n", r)
}