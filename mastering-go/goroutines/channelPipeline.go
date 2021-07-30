package main

import (
	"fmt"
	_ "time"
)

func main() {
	test1()
}

func test1() {
	fmt.Println("Pipeline ----------------")

	var ch1 chan int 
	ch1 = make(chan int)
	ch2 := make(chan int)
	done := make(chan bool)

	// counter. generates and sends numbers to squarer
	go func() {
		fmt.Println("counter start")
		for i:=1; i<= 1000; i++ {
			//fmt.Println("counter", i, "sending")
			ch1 <- i	
		}
		close(ch1)
	}()

	//squarer. receive numbers from counter and squre them and send to printer
	go func() {
		for n, ok := <- ch1; ok; n, ok = <- ch1 {
			//fmt.Println("squarer rec ", n, ok)
			s := n * n
			ch2 <- s
		} 
		//fmt.Println("squarer finish")
		close(ch2)
	}()

	//printer. 
	go func() {
		for n:= range ch2 {
			fmt.Println(n)
		}
		fmt.Println("Done")
		done <- true
	}()

	<- done
}