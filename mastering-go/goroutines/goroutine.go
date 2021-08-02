package main

import (
	"fmt"
	"time"
)

// when program starts, its only goroutine is "main goroutine"
// when main functions returns (means main goroutine terminates) all other goroutines are suddenly terminated and program exits.
// there is not any way for one goroutine to stop another goroutine (each goroutine only able to stop itself)
func main() {
	spinner(100 * time.Millisecond)
}

func test1() {
	// when program starts, its only goroutine is "main goroutine"
	// new gouroutines aree creayed by "go func"
	// go statement method parameters or receivers evaluation are like defer. all receivers and parameters evaluated when  go statement executed
	// 		and final method will be execute concurrently.
	go spinner(100 * time.Millisecond)

}

func spinner(delay time.Duration) {
	for {
		for _, c := range `-/|\` {
			fmt.Printf("\r%c", c)
			time.Sleep(delay)
		}
	}
}
