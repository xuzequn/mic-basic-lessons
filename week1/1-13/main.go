package main

import "fmt"

// interface 契约

type Food struct {
	No   int32
	Name string
}

// 构造函数 工厂函数
func NewFood(no int32, name string) *Food {
	return &Food{
		No:   no,
		Name: name,
	}
}

// 方法   接受者
func (f *Food) Show() {
	fmt.Println(f.Name)
}

//  指针传递接收者
func (f *Food) SetName(name string) {
	f.Name = name
}

// 函数
//func Show(f* Food)  {
//	fmt.Println(f.Name)
//}

func main() {
	var food Food
	food = Food{
		No:   1,
		Name: "葱烧海参",
	}
	fmt.Println(food)
	fmt.Println(food.No)
	fmt.Println(food.Name)

	food2 := Food{
		No:   2,
		Name: "麻婆豆腐",
	}
	fmt.Println(food2)

	food3 := NewFood(3, "回锅肉")
	fmt.Println(*food3)

	food3.Show()
	food3.SetName("宫保鸡丁")
	food3.Show()

}
