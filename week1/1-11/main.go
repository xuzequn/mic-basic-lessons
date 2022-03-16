package main

import "fmt"

func SoldOut(foods []string) {
	foods[1] = "已售罄"
	fmt.Println(foods)
}

func main() {
	array5 := [...]string{"烧乳猪", "鲍鱼", "文昌鸡", "大龙虾", "烧鹅", "茶田鸭"}

	// slice 数组的view
	slice1 := array5[1:3]
	fmt.Println(slice1)

	slice2 := array5[1:]
	fmt.Println(slice2)

	slice3 := array5[:3]
	fmt.Println(slice3)

	slice4 := array5[:]
	fmt.Println(slice4)

	SoldOut(slice4)
	fmt.Println(slice4)

	// 剩下的切片会跟着来到slice5
	slice5 := slice4[2:5]
	fmt.Println(slice5)

	slice6 := slice5[2:4]
	fmt.Println(slice6)

	slice7 := slice5[2:]
	fmt.Println(slice7)

	slice8 := []string{}
	fmt.Println(len(slice8), cap(slice8))
	for i := 0; i < 10; i++ {
		slice8 = append(slice8, "蛋炒饭")
		fmt.Println(len(slice8), cap(slice8))
	}

	s1 := make([]string, 8)
	fmt.Println(s1)
	s2 := make([]string, 8, 16)
	fmt.Println(s2)

	copy(s1, slice4)
	copy(s2, slice4)
	fmt.Println(s1)
	fmt.Println(s2)

	// delete
	//slice4[1:]
	fmt.Println(slice4[:len(slice4)-1])
}
