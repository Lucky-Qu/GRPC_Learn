package main

import (
	hello_grpc "GRPC_Learn/01_HelloWorld/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	//创建连接
	conn, err := grpc.NewClient("localhost:7777", grpc.WithTransportCredentials(insecure.NewCredentials()))
	//代码执行完毕后及时关闭连接
	defer conn.Close()
	//处理错误
	if err != nil {
		panic(err)
	}
	//创建客户
	client := hello_grpc.NewHelloGRPCClient(conn)
	//调用方法
	req, err := client.SayHi(context.Background(), &hello_grpc.Req{Message: "我需要一个浆果！"})
	if err != nil {
		panic(err)
	}
	//获取返回值
	fmt.Println(req.GetMessage())

}
