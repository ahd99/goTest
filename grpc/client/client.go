package client

import (
	"context"
	"fmt"

	"aheydari.ir/gotest/grpc/proto"
	"google.golang.org/grpc"
)

// SendName send name
func SendName(name string) {
	//var opts []grpc.DialOption

	conn, err := grpc.Dial("localhost:12520", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Error connecting to server : %v", err)
		return
	}
	defer conn.Close()

	client := proto.NewGreeterClient(conn)

	gresp, err := client.SayHello(context.Background(), &proto.GRequest{Name: "reza"})
	if err != nil {
		fmt.Printf("Error in call method. %v", err)
		return
	}

	fmt.Printf("Client -> sent: %v   Response: %v\n", name, gresp.Resp)
}
