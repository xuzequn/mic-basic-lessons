package main

import "fmt"

// func 函数名称(参数列表) (返回值列表)

func ShowName() string {
	return "从0到Go语言微服务架构师"
}

func ShowInfo(book, author string) string {
	return book + "-作者是" + author
}

func ShowInfoAndPrice(book, author string, price float64) (string, float64) {
	return book + "-作者是" + author, price
}

func ShowInfoAndPrice2(book, author string, price float64) (bookInfo string, finalPrice float64) {
	bookInfo = book + "-作者是" + author
	finalPrice = price
	return
}

func PriceBookInfo(do func(string, string, float64) (bookInfo string, finalPrice float64),
	bookName, author string, price float64) {
	bookNameInfo, finalPrice := do(bookName, author, price)
	fmt.Println(bookNameInfo)
	fmt.Println(finalPrice)
}

func PriceBookInfo2(bookName, author string, price float64) {
	do := func(string, string, float64) (bookInfo string, finalPrice float64) {
		bookInfo = bookName + author
		finalPrice = price
		return
	}
	bookNameInfo, finalPrice := do(bookName, author, price)
	fmt.Println(bookNameInfo)
	fmt.Println(finalPrice)
}

func ShowAll(pricelist ...string) string {
	r := ""
	for idx, item := range pricelist {
		fmt.Println(idx)
		r += item
	}
	return r
}

func main() {
	//fmt.Println(ShowName())
	//fmt.Println(ShowInfo("Go语言极简一本通","欢喜"))
	//fmt.Println(ShowInfoAndPrice2("Go语言极简一本通","欢喜",99.00))
	//info,_ := ShowInfoAndPrice2("Go语言极简一本通","欢喜",99.00)
	//fmt.Println(info)
	//PriceBookInfo(ShowInfoAndPrice2,"Go语言极简一本通","欢喜",99.00)
	PriceBookInfo2("Go语言极简一本通", "欢喜", 99.00)
	fmt.Println(ShowAll("葱烧海参", "九转大肠", "溏心鲍鱼", "麻辣兔头"))
}
