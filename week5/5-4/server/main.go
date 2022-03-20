package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type FoodService struct {
}

func (f *FoodService) SayName(request string, resp *string) error {
	*resp = "您点的菜是" + request
	return nil
}

func main() {
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		// 不能panic
		fmt.Println(err)
		return
	}
	//                  名字                接受类型
	err = rpc.RegisterName("FoodService", &FoodService{})
	if err != nil {
		fmt.Println(err)
		return
	}
	// Accept waits for and returns the next connection to the listener.
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}
	// 服务链接，
	rpc.ServeConn(conn)

}
