package main

import (
	"GRPC_Learn/03_Server/pb/person"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"sync"
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
	//m, _ := client.TestFunc2(context.Background())
	//for range 5 {
	//	time.Sleep(1 * time.Second)
	//	m.Send(&person.PersonReq{
	//		Name: "Hiiiii",
	//		Age:  66,
	//	})
	//}
	//res, _ := m.CloseAndRecv()
	//res, _ := client.TestFunc3(context.Background(), &person.PersonReq{
	//	Name: "HIIIiiii",
	//	Age:  12,
	//})
	//for {
	//	msg, err := res.Recv()
	//	fmt.Println(msg)
	//	if err == io.EOF {
	//		break
	//	}
	//}
	n, err := client.TestFunc4(context.Background())
	if err != nil {
		panic(err)
	}
	syn := sync.WaitGroup{}
	syn.Add(2)
	go func() {
		for {
			msg, _ := n.Recv()
			fmt.Println(msg)
			if msg == nil {
				break
			}
		}
		syn.Done()
	}()
	go func() {
		for range 7 {
			n.Send(&person.PersonReq{
				Name: "Hiiiiioo",
				Age:  199,
			})
		}
		syn.Done()
	}()
	syn.Wait()
}
