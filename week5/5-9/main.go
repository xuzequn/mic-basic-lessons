package main

import (
	"encoding/json"
	"fmt"
	"google.golang.org/protobuf/proto"
	"mic-basic-lessons/week5/5-9/proto/pb"
)

type BookInfo struct {
	Name string `json:"name"`
}

func main() {
	req := pb.BookRequest{Name: "Go语言极简一本通"}
	b, err := proto.Marshal(&req)
	if err != nil {
		panic(err)
	}
	fmt.Println(b)
	fmt.Println(string(b))

	info := BookInfo{
		Name: "Go语言极简一本通",
	}
	jsonByte, err := json.Marshal(info)
	if err != nil {
		panic(err)
	}
	fmt.Println(jsonByte)
	fmt.Println(string(jsonByte))
}
