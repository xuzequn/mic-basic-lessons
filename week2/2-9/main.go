package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func funcRecover() error {

	defer func() {
		if v := recover(); v != nil {
			fmt.Printf("Panic recover! v:%v", v)
		}
	}()
	return funcCook()
}

func funcCook() error {
	panic("停水了")
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
