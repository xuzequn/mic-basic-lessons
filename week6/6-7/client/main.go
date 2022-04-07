package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"mic-basic-lessons/week6/6-6/proto/pb"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:9094", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := pb.NewToDoServiceClient(conn)
	// 通过matedata 将信息放到context 中传到服务端，比如将token穿过去做认证
	//MD1 := metadata.New(map[string]string{
	//	"name": "从0到GO语言微服务架构师",
	//})
	MD2 := metadata.Pairs("name", "从0到GO语言微服务架构师")
	ctx := metadata.NewOutgoingContext(context.Background(), MD2)
	res, err := client.DoWork(ctx, &pb.TodoRequest{Todo: "我的学习计划"})
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Done)
}
