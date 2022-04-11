package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"mic-basic-lessons/week6/6-6/proto/pb"
	"net"
	"time"
)

func MyInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Println("没有metadata, 错误")
		return nil, status.Error(codes.Unauthenticated, "认证失败")
	}
	token, ok := md["token"]
	if !ok {
		fmt.Println("没有token, 错误")
		return nil, status.Error(codes.Unauthenticated, "认证失败")
	}
	fmt.Printf("token:%s", token)
	now := time.Now()
	resp, err = handler(ctx, req)
	d := time.Now().Sub(now)
	fmt.Printf("执行时间%d", d.Milliseconds())
	return
}

type ToDo struct {
}

func (t *ToDo) DoWork(c context.Context, req *pb.TodoRequest) (*pb.TodoResponse, error) {
	// 获取 matedata 通过context
	md, ok := metadata.FromIncomingContext(c)
	if !ok {
		fmt.Println("没有metadata")
	}
	for k, v := range md {
		fmt.Printf("%s:%s", k, v)
	}
	time.Sleep(time.Second * 2)
	fmt.Println(req.Todo + "已完成！")
	return &pb.TodoResponse{Done: true}, nil
}

func main() {
	serverOption := grpc.UnaryInterceptor(MyInterceptor)
	server := grpc.NewServer(serverOption)
	pb.RegisterToDoServiceServer(server, &ToDo{})
	listen, err := net.Listen("tcp", "0.0.0.0:9096")
	if err != nil {
		panic(err)
	}
	err = server.Serve(listen)
	if err != nil {
		panic(err)
	}

}
