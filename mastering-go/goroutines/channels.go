package main

import "fmt"

func main() {
	test1()
	fmt.Println("------------------- test2 ------")
}

//channel is a thread safe communication mechanism for sending data from a goroutine to another goroutine
func test1() {

	// declare a channel with element type = int
	// chanels are reference type so its zero value is nil
	// when we copy a channel by assignment or pass it as argument to function, both refer to the same underlying channel
	var numbers chan int     // chan == nil
	numbers = make(chan int) // numbers != nil - create a unbuffered (queue size = 0) channel of type int.
	fmt.Printf("%T\t%v\n", numbers, numbers)

	numbers1 := make(chan int, 5) // create a buffered channel wit capacity 5
	fmt.Printf("%T\n", numbers1)

	// channel are comparable by == or !=
	// comparison is true if both are references to the same channel
	numbers10 := numbers
	if numbers == numbers10 {
		fmt.Println("equal") // "equal"
	}

	numbers11 := make(chan int)
	if numbers != numbers11 {
		fmt.Println("not equal") // "not equal"
	}

	// because numbers is a unbuffered channel, call send on it, blocks sender goroutin until anoher goroutine call receive on it, at whitch point
	// the value trnsmitted from sender to receiver and bith goroutines continue their execution
	// and call recieve on it, blocks receiver goroutin until anoher goroutine call send on it
	// sending and receiving on unbufered channel cause sender and receiver goroutines to synchronised.
	// when a value is sent on an unbuffered channel, the receipt of the value in receive goroutine "happens before" the reawakening of sending goroutine
	// so need to call one of send or receive in another goroutine. removing go keyword (so calling receive from numbers channel in main goroutine),
	// cause deadlock, because main goroutine blocks on receive command and no one send value on numbers channel
	go func() {
		numbers <- 10 // send. chan <- n
	}()
	n := <-numbers //receive. n <- cahn
	fmt.Println(n)

	//after close a channel, send on it cause panic.
	// but receive returns data in chnnale queue until queue goes empty,
	// after that each receive on channel returns channel element type zero value emmidiately
	// closing channel is necessery only when we want to tell receiving goroutine that there is no more data for send. if a channel
	// be unreachable, will be reclaimed by garbage collector whether or not it is closed
	close(numbers)

	// n: value received from channel.
	// ok: is true for a seccessful receive, and is false when channel is closed and drained (no more data is in it).
	// when ok is false, n has zero vlaue of channel element type
	n, ok := <-numbers // numbers is closed and no data is waiting to be received so ok is false
	fmt.Println(n, ok) // 0, false

	// we can use range loop on a channel for receiving data. range loop receive all values send on channel and block when waiting for data and
	// terinates when channel closed and drained

	numbers20 := make(chan int, 5)
	go func() {
		for i := 0; i < 10; i++ {
			numbers20 <- i
		}
		close(numbers20) // removing this line cause deadlock. because range loop in main goroutine never terminates and stay blocked forever.
	}()

	for n1 := range numbers20 {
		fmt.Print(n1, "\t") // 0       1       2       3       4       5       6       7       8       9
	}
	fmt.Println()

}

func test2() {

}
