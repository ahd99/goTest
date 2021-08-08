package main

import (
	"fmt"
	"sync"
)

func main() {
	test1(45)
}

func test1(count int64) {

	var wg sync.WaitGroup
	results := make(chan int64)

	for i := int64(1); i < count; i++ {
		// add(n): adds n to waitGroup iternal counter
		wg.Add(1) //add must be called before goroutine, not within it. because we must be sure that all adds to wg must "happens before" calling wg.wait
		go func(i int64) {
			//done() decrease 1 from waitgroup internal counter
			defer wg.Done() // wg.Done() must call after finishing each goroutine.
			res := fibunacci(i)
			results <- res
		}(i)
	}

	// closer function
	// afeter making sure all calls to wg.add() happens, we call wg.wait to wait until all goroutines finish.
	// because we must be sure that all adds to wg must "happens before" calling wg.wait
	go func() {
		wg.Wait()      // block goroutine until internal counter be equal to zero
		close(results) //if we dont close results channel, following range will remian blocked forever
	}()

	for r := range results {
		fmt.Printf("%d\t", r)
	}
	fmt.Println()
}

func fibunacci(n int64) int64 {
	if n < 2 {
		return n
	}
	return fibunacci(n-1) + fibunacci(n-2)
}
