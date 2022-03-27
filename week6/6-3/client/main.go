package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"mic-basic-lessons/week6/6-3/proto/pb"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:9093", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := pb.NewMistakeServiceClient(conn)
	res, err := client.MakeMistake(context.Background(), &pb.MistakeRequest{
		Title:    "Go语言极简一本通",
		SubTitle: "一学就会，一本就通",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(res)

}
