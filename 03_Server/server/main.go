package main

import (
	"GRPC_Learn/03_Server/pb/person"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"net"
)

type personServer struct {
	person.UnimplementedTestServiceServer
}

func (*personServer) TestFunc1(ctx context.Context, req *person.PersonReq) (*person.PersonRes, error) {
	name := req.GetName()
	age := req.GetAge()
	return &person.PersonRes{
		Name: name,
		Age:  age,
	}, nil
}
func (*personServer) TestFunc2(client grpc.ClientStreamingServer[person.PersonReq, person.PersonRes]) error {
	for {
		rec, err := client.Recv()
		fmt.Println("我接受到消息：", rec)
		if err == io.EOF {
			fmt.Println("接受完毕来自客户端的消息")
			client.SendAndClose(&person.PersonRes{
				Name: "接受完毕你的消息",
				Age:  777,
			})
			break
		} else if err != nil {
			panic(err)
		}
	}
	return nil
}
func (*personServer) TestFunc3(*person.PersonReq, grpc.ServerStreamingServer[person.PersonRes]) error {
	return nil
}
func (*personServer) TestFunc4(grpc.BidiStreamingServer[person.PersonReq, person.PersonRes]) error {
	return nil
}

func main() {
	l, _ := net.Listen("tcp", "localhost:7777")
	grpcServer := grpc.NewServer()
	person.RegisterTestServiceServer(grpcServer, &personServer{})
	err := grpcServer.Serve(l)
	if err != nil {
		panic(err)
	}
}
