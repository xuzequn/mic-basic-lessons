package main

import (
	"fmt"
	"mic-basic-lessons/week6/6-5/proto/pb"
)

func main() {
	req := pb.UpperRequest{
		Id:   1,
		Name: "面向加薪学习",
	}
	fmt.Println(req)
	platform := pb.Platform{
		Name:      "B站",
		FansCount: "100w",
	}
	res := pb.UpperResponse{
		Name:     "面向加薪学习",
		Platform: &platform,
	}
	fmt.Println(res)

}
