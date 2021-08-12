package main

import (
	"context"
	"fmt"
	"os"
	"time"
)

func main() {
	test1()
	fmt.Println("------------------------- test2() ------")
	test2()
	fmt.Println("------------------------- test3() ------")
	test3()

}

func test1() {
	ctx := context.Background()
	fmt.Printf("%T\t%v\n", ctx, ctx)

	ctx, cancelFunc := context.WithTimeout(ctx, 5*time.Second)
	go getConsoleCommand(cancelFunc)

	select {
	case <-ctx.Done():
		fmt.Println("context Done. reason:", ctx.Err())
		return
	case <-time.After(10 * time.Second):
		fmt.Println("after 10 seconds")
		return

	}

}

func getConsoleCommand(cancleFunc context.CancelFunc) {
	os.Stdin.Read(make([]byte, 1))
	cancleFunc()

}

func test2() {

	/*
		ctx_000 --> ctx_010 --> ctx_011
		                    --> ctx_012
		       --> cctx_020
	*/

	ctx_000, cancelFunc_000 := context.WithCancel(context.Background())
	
	ctx_010, cancelFunc_010 := context.WithCancel(ctx_000)
	ctx_011, cancelFunc_011 := context.WithCancel(ctx_010)
	ctx_012, cancelFunc_012 := context.WithCancel(ctx_010)

	ctx_020, cancelFunc_020 := context.WithCancel(ctx_000)
		
	_ = cancelFunc_000
	_ = cancelFunc_010
	_ = cancelFunc_011
	_ = cancelFunc_012
	_ = cancelFunc_020

	//var wg sync.WaitGroup
	//wg.Add(5)

	go waitForCtxClose(ctx_000, "ctx_000")
	go waitForCtxClose(ctx_010, "ctx_010")
	go waitForCtxClose(ctx_011, "ctx_011")
	go waitForCtxClose(ctx_012, "ctx_012")
	go waitForCtxClose(ctx_020, "ctx_020")

	time.Sleep(1 * time.Second)

	// cancelling a context or reaching its deadline, close Done channle of itself and all of its childrens (and children of childrens ...) 
	// and dont have effect on its parent and its siblings (nei) in hierarchy
	cancelFunc_000()		// cancelling parent closes all contexts
	// cancelFunc_010()		// close 010 (itself), 011 (child), 012 (child) and has not effect on 000 (parent), 020 (sibling)
	//cancelFunc_011()		// close only itself because has not any child
	//cancelFunc_020()		// close only itself because has not any child
	

	time.Sleep(2 * time.Second)

}

func waitForCtxClose(ctx context.Context, desc string) {
	<- ctx.Done()
	fmt.Println(desc, "closed. err:", ctx.Err(), time.Now())
}


func test3() {
	//when dedline of a context is reached, Done channel of context itself and all of its childrens will be closed

	ctx, _ := context.WithTimeout(context.Background(), 5 * time.Second)
	ctx1, _ := context.WithTimeout(ctx, 8* time.Second)  // ctx1 deadline is equal to parent ctx deadline because child deadline can not be after parent deadline
	ctx2, _ := context.WithTimeout(ctx, 3 * time.Second) // ctx2 dealine is before parent, so it is closed before parent
	fmt.Println(ctx.Deadline())
	fmt.Println(ctx1.Deadline())
	fmt.Println(ctx2.Deadline())

	go waitForCtxClose(ctx, "ctx ")
	go waitForCtxClose(ctx1, "ctx1")
	go waitForCtxClose(ctx2, "ctx2")

	time.Sleep(10 * time.Second)
}
