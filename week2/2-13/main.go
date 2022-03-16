package main

import "fmt"

func GivenFood() chan string {
	ch := make(chan string)
	go func() {
		ch <- "回锅肉"
		ch <- "炭烤生蚝"
		ch <- "担担面"
		close(ch)
	}()
	return ch

}

//发送
func OnlyReceive(ch chan<- int) {

}

// 读取
func OnlyRead(ch <-chan int) {

}
func main() {
	//channel
	//ch := make(chan string)
	//ch2 := make(chan string, 6)
	//ch <- "回锅肉"
	//<- ch2
	//close(ch)

	ch := make(chan string)
	ch = GivenFood()

	// first type
	//for {
	//	if name,ok:= <-ch; ok{
	//		fmt.Println(name)
	//	}else {
	//		break
	//	}
	//}

	//  second type
	for data := range ch {
		fmt.Println(data)
	}

	ch2 := make(chan<- int)
	OnlyReceive(ch2)

}
