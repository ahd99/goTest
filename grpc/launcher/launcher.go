package main

import (
	"fmt"
	"time"

	"aheydari.ir/gotest/grpc/client"
	"aheydari.ir/gotest/grpc/server"
)

func main() {
	fmt.Println("GRPC test start !")
	go server.StartServer()

	time.Sleep(3 * time.Second)

	client.SendName("reza")
}
