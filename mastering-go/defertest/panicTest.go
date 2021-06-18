package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func main()  {
	//defer fmt.Println("-------- main defer")
	defer printStack()
	fmt.Println("in main")
	//go func3()
	//ßtime.Sleep(2 * time.Second)
	func1()
}

func func1() {
	defer fmt.Println("--------- func1 defer")
	fmt.Println("in func1")
	func2()
}

func func2() {
	defer fmt.Println("--------- func2 defer")
	fmt.Println("in func2")
	//time.Sleep(10 * time.Second)
	panic("panic")
}

func func3() {
	defer fmt.Println("========== func3 defer")
	fmt.Println("== in func3")
	time.Sleep(5 * time.Second)
	panic("panic 1")
	fmt.Println("== func3 finish")
}


func printStack() {
var buf [4096]byte
n := runtime.Stack(buf[:], false) 
os.Stdout.Write(buf[:n])
}