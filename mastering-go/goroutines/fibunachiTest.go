package main

import (
	"fmt"
	"time"
)

func main() {
	//fmt.Println(fibunacci(6))
	//printFibunacciSeries(45)

	go spinner(100 * time.Millisecond)
	res := fibunacci(50)
	fmt.Printf("\r%d\n", res)
	//time.Sleep(5 * time.Second)
}

func printFibunacciSeries(n int64) {
	for i := int64(1); i <= n; i++ {
		fmt.Printf("%d\t", fibunacci(i))
	}
	fmt.Println()
}

func fibunacci(n int64) int64 {
	if n < 2 {
		return n
	}
	return fibunacci(n-1) + fibunacci(n-2)
}

func spinner(delay time.Duration) {
	//fmt.Println("------")
	for {
		for _, c := range `-/|\` {
			fmt.Printf("\r%c", c)
			time.Sleep(delay)
		}
	}
}
