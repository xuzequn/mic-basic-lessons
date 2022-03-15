package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"strings"
)

func Show (){
	fmt.Println("从0到Go语言微服务架构师")
}

// 函数别名
type CheckOut func(int, int) int

func GotTotal(x int) func(int) int {
	return func(y int) int {
		// 闭包
		return  x+ y
	}
}

// 函数别名, 函数类型
type GenerateRandom func() int

func RandomSum() GenerateRandom{
	a, b := rand.Intn(10), rand.Intn(20)
	fmt.Println("---GenerateRandom---")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println("---GenerateRandom End---")
	return func() int {
		a,b = b ,a+b
		return a
	}
}

func (g GenerateRandom) Read(p []byte) (n int, err error){
	next := g()
	if next >21 {
		fmt.Println(">21")
		fmt.Println(next)
		fmt.Println(">21,end...")
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func PrintResult(reader io.Reader){
	scanner := bufio.NewScanner(reader)
	for scanner.Scan(){
		fmt.Println("---Scan()---")
		fmt.Println(scanner.Text())
		fmt.Println("---Scan() End---")
	}
}

func main(){

	// 函数变量
	Show()
	show := Show
	show()

	// 匿名函数
	show2 := func(){
		fmt.Println("Go语言核心22讲")
	}
	show2()

	// 函数类型 or 类型函数
	var checkOut CheckOut = func(a int, b int) int {
		return  a +b
	}
	fmt.Println(checkOut(68,98))

	// 函数返回值, 闭包
	fmt.Println("-----------------GotTotal---------------")
	total := GotTotal(68)
	sum := total(100)
	fmt.Println(sum)
	total = GotTotal(sum)
	sum = total(50)
	fmt.Println(sum)
	total = GotTotal(sum)
	sum = total(200)
	fmt.Println(sum)

	fmt.Println("-----------------RandomSum---------------")
	r := RandomSum()
	PrintResult(r)

}
