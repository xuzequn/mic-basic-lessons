package main

import "fmt"

func byValue(price int) {
	price += 20
}

func byRef(price *int) {
	*price += 20
}

func main() {
	var price int = 68
	//var ptr * int = &price // & 取地址  * 取值
	//fmt.Println(ptr)
	//
	//*ptr = 88
	//fmt.Println(price)

	byValue(price)
	fmt.Println(price)
	byRef(&price)
	fmt.Println(price)
}
