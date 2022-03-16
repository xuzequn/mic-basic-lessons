package main

import "fmt"

// interface 契约

func main() {
	//GoToDinner("xxx")

	var w Where = LaoZhang{}
	w.WhereDiner()

	var g GoToHappy = XiaoHong{}
	g.GoToDinner("")

	w = LaoWang{}
	w.WhereDiner()
	g = XiaoLi{}
	g.GoToDinner("XXX")

}

type Where interface {
	WhereDiner()
}

type GoToHappy interface {
	GoToDinner(dest string)
}

type LaoWang struct {
}

type LaoZhang struct {
}

type XiaoLi struct {
}

type XiaoHong struct {
}

func (LaoWang) WhereDiner() {
	fmt.Println("周末去吃大餐吗")
	fmt.Println("法餐？日料？火锅？烧烤？")
}

func (LaoZhang) WhereDiner() {
	fmt.Println("周末去吃大餐吗")
	fmt.Println("找老王他们来家里做客吗？")
}

func (XiaoLi) GoToDinner(dest string) {
	fmt.Println("女朋友：吃火锅" + dest + "火锅店")
	fmt.Println("周末，就去" + dest + "火锅店，愉快的玩耍啦")
}

func (XiaoHong) GoToDinner(dest string) {
	if dest == "" {
		fmt.Println("好的, 来家里吃吧。")
	} else {
		fmt.Println("女朋友：吃火锅" + dest + "火锅店")
		fmt.Println("去" + dest + "火锅店，愉快的玩耍啦")
	}

}

func GoToDinner(dest string) {

}
