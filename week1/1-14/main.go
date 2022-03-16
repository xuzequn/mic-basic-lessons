package main

import (
	"fmt"
	"mic-basic-learning/week1/1-14/hello"
)

func main() {
	//var foodList []string // 私有
	//var Food string // 公开

	var Food model.Food
	fmt.Println(Food)

	price := 68
	newPrice := model.MyInt(price)
	fmt.Println(newPrice)

	// 组合
	food := model.Food{Name: "豆汁儿"}
	p := model.Person{Food: food}
	p.SayHello("老张")

	food2 := model.Food2{
		FoodName:   "葱烧海参",
		FoodSystem: model.FoodSystem{Name: "鲁菜"},
		Another:    model.FoodSystem{Name: "鲁菜"},
	}

	fmt.Println(food2.Another.Name)
	fmt.Println(food2.Name)
}
