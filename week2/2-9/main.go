package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func funcRecover() error {

	defer func() {
		if v := recover(); v != nil { // 2 err 被处理掉了 err 为空
			fmt.Printf("Panic recover! v:%v", v) // 3
		}
	}()
	return funcCook()
}

func funcCook() error {
	panic("停水了") // 1
	return errors.New("发生错误了")
}

func main() {
	// panic recover defer

	err := funcRecover()
	if err != nil {
		fmt.Printf("err is %v", err)
	} else {
		fmt.Println("err is nil")
	}
}
