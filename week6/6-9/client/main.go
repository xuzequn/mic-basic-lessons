package main

import (
	"context"
	"google.golang.org/grpc"
	"mic-basic-lessons/week6/6-6/proto/pb"
)

type MyCredntials struct {
}

func (c *MyCredntials) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{"token": "面向加薪学习"}, nil
}

func (c *MyCredntials) RequireTransportSecurity() bool {
	return false
}

func main() {
	//ClientInterceptor := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	//	md := metadata.Pairs("token", "欢喜-Go语言极简一本通")
	//	ctx = metadata.NewOutgoingContext(context.Background(), md)
	//	now := time.Now()
	//	// invoker 调用者
	//	err := invoker(ctx, method, req, reply, cc, opts...)
	//	d := time.Now().Sub(now)
	//	fmt.Printf("客户端调用-执行时间%d", d.Milliseconds())
	//	return err
	//}
	//opt := grpc.WithUnaryInterceptor(ClientInterceptor)

	opt := grpc.WithPerRPCCredentials(&MyCredntials{})
	conn, err := grpc.Dial("127.0.0.1:9096", grpc.WithInsecure(), opt)
	defer conn.Close()
	if err != nil {
		panic(err)
	}

	client := pb.NewToDoServiceClient(conn)
	client.DoWork(context.Background(), &pb.TodoRequest{Todo: "go语言微服务核心22讲"})

}
