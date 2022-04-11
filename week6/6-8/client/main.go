package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"mic-basic-lessons/week6/6-6/proto/pb"
	"time"
)

func main() {
	ClientInterceptor := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		now := time.Now()
		// invoker 调用者
		err := invoker(ctx, method, req, reply, cc, opts...)
		d := time.Now().Sub(now)
		fmt.Printf("客户端调用执行时间%d", d.Milliseconds())
		return err
	}

	conn, err := grpc.Dial("127.0.0.1:9096", grpc.WithInsecure(), grpc.WithUnaryInterceptor(ClientInterceptor))
	defer conn.Close()
	if err != nil {
		panic(err)
	}

	client := pb.NewToDoServiceClient(conn)
	client.DoWork(context.Background(), &pb.TodoRequest{Todo: "go语言微服务核心22讲"})

}
