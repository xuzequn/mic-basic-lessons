package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"mic-basic-lessons/week6/6-6/proto/pb"
	"net"
	"time"
)

type ToDo struct {
}

func (t *ToDo) DoWork(ctx context.Context, req *pb.TodoRequest) (*pb.TodoResponse, error) {
	//return nil,status.Error(codes.Unknown,"服务端内部错误")
	time.Sleep(10 * time.Second)
	fmt.Println(req.Todo + "已完成！")
	return &pb.TodoResponse{Done: true}, nil
}

func main() {
	server := grpc.NewServer()
	pb.RegisterToDoServiceServer(server, &ToDo{})
	listen, err := net.Listen("tcp", "0.0.0.0:9097")
	if err != nil {
		panic(err)
	}
	err = server.Serve(listen)
	if err != nil {
		panic(err)
	}

}
