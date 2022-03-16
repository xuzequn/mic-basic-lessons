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
	h.GoToDinner(dest)
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

func main() {
	b := Boy{Name: "小明"}
	g := Girl{Name: "小兰"}
	p := Pet{Name: "大黄"}
	Happy(&b, &g, "野炊吧")
	Happy(&b, &p, "在家吃肉肉")

}
