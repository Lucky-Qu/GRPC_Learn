package main

import (
	"GRPC_Learn/03_Server/pb/person"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"net"
	"sync"
	"time"
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
func (*personServer) TestFunc3(clientReq *person.PersonReq, server grpc.ServerStreamingServer[person.PersonRes]) error {
	fmt.Println(clientReq)
	for range 5 {
		time.Sleep(time.Second)
		err := server.Send(&person.PersonRes{
			Name: "wow",
			Age:  10,
		})
		if err != nil {
			panic(err)
		}
	}
	return nil
}
func (*personServer) TestFunc4(serverAndClient grpc.BidiStreamingServer[person.PersonReq, person.PersonRes]) error {
	s := sync.WaitGroup{}
	s.Add(2)
	go func() {
		for {
			req, _ := serverAndClient.Recv()
			fmt.Println(req)
			if req == nil {
				break
			}
		}
		s.Done()
	}()
	go func() {
		for range 3 {
			time.Sleep(2)
			serverAndClient.Send(&person.PersonRes{
				Name: "here",
				Age:  1,
			})
		}
		s.Done()
	}()
	s.Wait()
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
