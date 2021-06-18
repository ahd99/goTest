package main

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

func main() {
	printStat()

	for i:=1; i<10; i++ {
		s := make([]byte, 50000000)
		if s == nil {
			log.Println("Operation failed")
		}
	}

	printStat()

	for i:=1; i<10; i++ {
		s := make([]byte, 100000000)
		if s == nil {
			log.Println("Operation failed")
		}
		time.Sleep(5 * time.Second)
	}

	printStat()

}

func printStat() {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	fmt.Println("mem.Alloc:", mem.Alloc) 
	fmt.Println("mem.TotalAlloc:", mem.TotalAlloc) 
	fmt.Println("mem.HeapAlloc:", mem.HeapAlloc) 
	fmt.Println("mem.NumGC:", mem.NumGC) 
	fmt.Println("-----")

}
