package main

import (
	"fmt"
	"runtime"
	"time"
)

func ShowBook() {
	fmt.Println("Go语言极简一本通")
}

func main() {
	// go
	go ShowBook()

	for i := 0; i < 10; i++ {
		go func(j int) {
			fmt.Println(fmt.Sprintf("i am %d", j))
		}(i)
	}
	runtime.Gosched()

	time.Sleep(time.Second * 1)

}
