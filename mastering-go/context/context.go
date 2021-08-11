package main

import (
	"context"
	"fmt"
	"os"
	"time"
)

func main() {
	test1()

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
