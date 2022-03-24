package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"mic-basic-lessons/week6/6-1/proto/pb"
	"net"
)

type DefaultValue struct {
}

func (d DefaultValue) GetDefault(ctx context.Context, req *pb.DefaultRequest) (*pb.DefaultResponse, error) {
	fmt.Println(req)
	return &pb.DefaultResponse{
		Msg: "调用成功",
	}, nil
}

func main() {
	server := grpc.NewServer()
	//  RegisterDefaultServiceServer(s *grpc.Server, srv DefaultServiceServer)
	pb.RegisterDefaultServiceServer(server, &DefaultValue{})
	listen, err := net.Listen("tcp", "0.0.0.0:9092")
	if err != nil {
		panic(err)
	}
	err = server.Serve(listen)
	if err != nil {
		panic(err)
	}
}
