package main

import "fmt"

var (
	book = "\nGo语言极简一本通"
	lesson1 = "\nGo语言微服务核心22讲"
	lesson2 = "\n从0到Go语言微服务架构师路线图"
)
func main(){
	var price int32 = 68
	var name string = "宫保鸡丁"
	fmt.Println("hello world")
	fmt.Printf("%d, %s",price, name)

	var price2, wight = 68, 1
	fmt.Printf("\n%d,%d",price2, wight)

	price3 := 2
	wight2 := 1
	fmt.Printf("\n%d,%d",price3, wight2)

	fmt.Println(book,lesson1,lesson2)
}