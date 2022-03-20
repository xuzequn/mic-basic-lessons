package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	c, err := rpc.Dial("tcp", "localhost:9090")
	if err != nil {
		fmt.Println(err)
		return
	}

	reply := ""

	err = c.Call("FoodService.SayName", "九转大肠", &reply)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(reply)

}
