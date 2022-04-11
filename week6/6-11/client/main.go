package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"mic-basic-lessons/week6/6-6/proto/pb"
	"time"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:9097", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := pb.NewToDoServiceClient(conn)
	ctx, cancelFunc := context.WithTimeout(context.Background(), 12*time.Second)
	defer cancelFunc()
	res, err := client.DoWork(ctx, &pb.TodoRequest{
		Todo: "Go语言微服务架构核心22讲",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Done)
}
