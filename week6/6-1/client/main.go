package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"mic-basic-lessons/week6/6-1/proto/pb"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:9092", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := pb.NewDefaultServiceClient(conn)
	res, err := client.GetDefault(context.Background(), &pb.DefaultRequest{})
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Msg)

}
