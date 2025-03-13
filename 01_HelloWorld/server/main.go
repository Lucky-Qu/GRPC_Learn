package main

import (
	hello_grpc "GRPC_Learn/01_HelloWorld/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	hello_grpc.UnimplementedHelloGRPCServer
}

func (s *server) SayHi(ctx context.Context, req *hello_grpc.Req) (res *hello_grpc.Res, err error) {
	fmt.Println(req.GetMessage())
	return &hello_grpc.Res{Message: "浆果*1"}, nil
}

func main() {
	l, err := net.Listen("tcp", "localhost:7777")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	hello_grpc.RegisterHelloGRPCServer(s, &server{})
	err = s.Serve(l)
	if err != nil {
		panic(err)
	}
}
