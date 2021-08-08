package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan bool)
	go countDown(10, abort)
	getUserInput(abort)
	time.Sleep(100 * time.Second)
}

func test1() {
	//	select has multiple cases and a default. each case has a send or receive on some channel.
	// select is blocked until communication (send or receive) for at least one of cases is ready (dont block)
	// 		in this situation, only the ready case is executed and after that select statement will be finished.
	// if more than one cases are ready, select choose one of them randomly and next one is not executed. (put select in a loop to be sure second one will be executed in next iteration.)
	// a select with no case waits forever: select {}
	// a closed channel is always reday for receive.

	ch1 := make(chan int)
	select {
	case <-ch1:
		fmt.Println("ch1: data received and ignored")
		//-------
	case n1 := <-ch1:
		fmt.Println("ch2: data received:", n1)
	case ch1 <- 10:
		fmt.Println("ch3: data sent.")
	}

	// when a aselect hs default section, default section will be executed if no case is ready and select statement finish. in this situation select dont wait at all.
	//		when execUtion reach a select with default section, if one case is ready, it executed, if not, default is executed and program continue its execution afetr select
	// use select with default section for non-blocking send/receive of channels:
	select {
	case n := <-ch1:
		fmt.Println("data is ready on ch1 and received:", n)
	default:
		fmt.Println("data is not ready on channel ch1 and we dont wait for it and continue without ch1 data")
	}
}

func countDown(max int, abort <-chan bool) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for i := 1; i <= 10; i++ {

		select {
		case <-ticker.C:
			fmt.Println(i)
		case <-abort:
			fmt.Println("Aborted")
			os.Exit(1)
		}
	}
	fmt.Println("Launched")
	os.Exit(1)
}

func getUserInput(abort chan<- bool) {
	fmt.Println("Press neter to abort.")
	os.Stdin.Read(make([]byte, 1))
	abort <- true
}
