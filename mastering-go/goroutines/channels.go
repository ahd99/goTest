package main

import (
	"fmt"
	"time"
)

func main() {
	test1()
	fmt.Println("------------------- test2 ------")
	test2()
	fmt.Println("------------------- test3 ------")
	test3()
}

//channel is a thread safe communication mechanism for sending data from a goroutine to another goroutine
func test1() {

	// declare a channel with element type = int
	// chanels are reference type so its zero value is nil
	// when we copy a channel by assignment or pass it as argument to function, both refer to the same underlying channel
	var numbers chan int     // chan == nil
	numbers = make(chan int) // numbers != nil - create a unbuffered (queue size = 0) channel of type int.
	fmt.Printf("%T\t%v\n", numbers, numbers)

	numbers1 := make(chan int, 5) // create a buffered channel wit capacity 5. capacity means max size of queue inside buffered channel.
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
	// uni-directional channels ---------
	abort := make(chan bool) // abort is bidirectional so we can call both send and receive on it

	// send only channel. wrong use cause compile error
	// in calling following method with abort channel, abort convert implicitly to send-only channel
	// convert from bidirectional chanel to uni-directional is valid but the opposit is not possible
	go func(abort chan<- bool) {	// this abort (local to anonymous function) is a send-only channel
		time.Sleep(1 * time.Second)
		abort <- true
		abort <- true
	}(abort)

	// receive only channel. wrong use cause compile error
	// call close() on a receive only channel cause compile tine error
	go func(abort <-chan bool) {	// this abort (local to anonymous function) is a receive-only channel
		<- abort
		fmt.Println("received")
	}(abort)

	time.Sleep(2 * time.Second)
}


func test3() {
	// buffered channels
	numbers := make(chan int, 3) // a buffered channel with capacity 3 and current len 0
	//cap(channel) : channel capacity 
	fmt.Println("cap:", cap(numbers))	// "3"
	//len(channel) : numbers of values currently in channel queue
	l := len(numbers)	// "0"
	fmt.Println("len:", l)
	numbers <- 2	// send doesn't vlock because numbers is buffered channel with capacity 3
	numbers <- 3
	l = len(numbers)	// "2"  
	fmt.Println("len:", l)
	// cap(numbers)==3, len(numbers)==2, so channel is not full and is not empty, so both send and receive on it dont blocked
	numbers <- 4
	fmt.Println("len:", len(numbers))
	// now cap(numbers)==3, len(numbers)==3, so channel is full, so send on it blocks but receive on it does't block
	// numbers <- 5  //cause deadlock
	n1 := <-numbers		// n1==2	channel buffer like as queue (FIFO), first send item is received first
	fmt.Println("n1:", n1)
	n1 = <-numbers
	n1 = <- numbers
	fmt.Println("n1:", n1)
	fmt.Println("len:", len(numbers))
	// now len(numbers)==0, so channel is empty, so receive on it blocks but send does'nt block
	// <- numbers	// cause deadlock
	numbers <- 3	// len(numbers) == 1



	// goroutine leak: when a goroutine block sending or receiving from a channel and never another goroutine receive or send to that channel, 
	// and this goroutine will be remained forever. 
	// this blocked (leaked) goroutines doesn't reclaimed by GC automatically and if create a lot of them causes app to crash or run out of mempory
	
}