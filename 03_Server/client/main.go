package main

import (
	"GRPC_Learn/03_Server/pb/person"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

func main() {
	c, err := grpc.NewClient("localhost:7777", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer c.Close()
	client := person.NewTestServiceClient(c)
	//res, _ := client.TestFunc1(context.Background(), &person.PersonReq{
	//	Name: "LuckyQu",
	//	Age:  20,
	//})
	m, _ := client.TestFunc2(context.Background())
	for range 5 {
		time.Sleep(1 * time.Second)
		m.Send(&person.PersonReq{
			Name: "Hiiiii",
			Age:  66,
		})
	}
	res, _ := m.CloseAndRecv()
	fmt.Println(res)
}
