package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"mic-basic-lessons/week6/6-3/proto/pb"
	"net"
)

type Mistake struct {
}

func (m *Mistake) MakeMistake(ctx context.Context, req *pb.MistakeRequest) (*pb.MistakeResponse, error) {
	s := fmt.Sprintf("title:%s, subTitle:%s", req.Title, req.SubTitle)
	fmt.Println(s)
	return &pb.MistakeResponse{}, nil

}

func main() {
	server := grpc.NewServer()
	pb.RegisterMistakeServiceServer(server, &Mistake{})
	listen, err := net.Listen("tcp", "0.0.0.0:9093")
	if err != nil {
		panic(err)
	}
	err = server.Serve(listen)
	if err != nil {
		panic(err)
	}

}
