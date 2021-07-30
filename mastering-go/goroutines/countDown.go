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

func countDown(max int, abort <-chan bool) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for i:=1; i<=10; i++ {
		select {
		case <- ticker.C:
			fmt.Println(i)
		case <- abort:
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