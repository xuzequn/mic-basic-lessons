package main

import "fmt"

type Where interface {
	WhereDinner()
}

type GoToHappy interface {
	GoToDinner(dest string)
}

func Happy(w Where, h GoToHappy, dest string) {
	w.WhereDinner()
	fmt.Printf("%T", "%v", w, w)
	if instance, ok := w.(*Boy); ok {
		fmt.Println("类型判断正确：" + instance.Name)
	}
	h.GoToDinner(dest)
	fmt.Printf("%T", "%v\n", h, h)
}

type Boy struct {
	Name string
}

func (b *Boy) WhereDinner() {
	fmt.Println(b.Name + ",周末去吃大餐吗？")
}

type Girl struct {
	Name string
}

func (g *Girl) GoToDinner(dest string) {
	fmt.Println(g.Name + "去" + dest)
}

type Pet struct {
	Name string
}

func (p *Pet) GoToDinner(dest string) {
	fmt.Println(p.Name + dest)
}

func SwitchInterface(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("输入为整形")
	case string:
		fmt.Println("输入为字符串")
	case Where:
		fmt.Println("实现了Where接口")

	}

}

func main() {
	b := Boy{Name: "小明"}
	g := Girl{Name: "小兰"}
	p := Pet{Name: "大黄"}
	Happy(&b, &g, "野炊吧")
	Happy(&b, &p, "在家吃肉肉")

	SwitchInterface("Go语言极简一本通")
	SwitchInterface(99)
	SwitchInterface(&b)
}
