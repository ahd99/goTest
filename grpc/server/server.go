package server

import (
	"context"
	"fmt"
	"net"

	"aheydari.ir/gotest/grpc/proto"
	"google.golang.org/grpc"
)

type greetingServer struct {
	proto.UnimplementedGreeterServer
	helloStr string
}

func (s *greetingServer) SayHello(context context.Context, req *proto.GRequest) (resp *proto.GResponse, err error) {

	res := &proto.GResponse{Resp: s.helloStr + " " + req.GetName()}
	fmt.Println("Server -> received: ", req.GetName(), "   sent: ", res)
	return res, nil
}

// StartServer start server
func StartServer() {
	lis, err := net.Listen("tcp", "localhost:12520")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}
	defer lis.Close()

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	proto.RegisterGreeterServer(grpcServer, &greetingServer{helloStr: "Merhaba"})
	err1 := grpcServer.Serve(lis)
	if err1 != nil {
		fmt.Printf("failed to start: %v", err1)
		return
	}

	fmt.Println("Server started.")

}
