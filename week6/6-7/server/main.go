package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"mic-basic-lessons/week6/6-6/proto/pb"
	"net"
)

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
	fmt.Println(req.Todo + "已完成！")
	return &pb.TodoResponse{Done: true}, nil
}

func main() {
	// auth http header token gin . grpc metadata

	server := grpc.NewServer()
	pb.RegisterToDoServiceServer(server, &ToDo{})
	listen, err := net.Listen("tcp", "0.0.0.0:9094")
	if err != nil {
		panic(err)
	}
	err = server.Serve(listen)
	if err != nil {
		panic(err)
	}
}
