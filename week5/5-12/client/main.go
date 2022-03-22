package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"mic-basic-lessons/week5/5-12/proto/pb"
	"sync"
)

func main() {
	conn, err := grpc.Dial("localhost:9091", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := pb.NewFoodServiceClient(conn)
	// 服务端流模式
	//res, err := client.SayName(context.Background(), &pb.FoodStreamRequest{Name: "干锅肥肠"})
	//if err != nil {
	//	panic(err)
	//}
	//// 流模式接受
	//for {
	//	recv, err := res.Recv()
	//	if err != nil {
	//		fmt.Println(err)
	//		break
	//	}
	//	fmt.Println(recv.Msg)
	//}

	// 客户端流模式
	//                                                          10秒以后过期
	//ctx, cancelFunc:=context.WithTimeout(context.Background(), 10 *time.Second)
	//defer cancelFunc()
	//postNameClient, err := client.PostName(ctx)
	//if err != nil {
	//	panic(err)
	//}
	foods := []string{"东坡肘子", "回锅肉", "夫妻肺片", "鱼香茄子", "辣子鸡", "水煮牛肉", "宫保鸡丁"}
	//for _, item := range foods{
	//	fmt.Println("上菜"+item)
	//	err := postNameClient.Send(&pb.FoodStreamRequest{Name: item})
	//	time.Sleep(time.Second*1)
	//	if err != nil {
	//		fmt.Println(err)
	//		break
	//	}
	//}

	//双向流的模式
	fullClient, err := client.FullStream(context.Background())
	if err != nil {
		panic(err)
	}
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			item, err := fullClient.Recv()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(item.Msg)
		}
	}()

	go func(s []string) {
		defer wg.Done()
		for _, item := range s {
			err := fullClient.Send(&pb.FoodStreamRequest{Name: item})
			if err != nil {
				fmt.Println(err)
			}
		}
	}(foods)

	wg.Wait()

}
