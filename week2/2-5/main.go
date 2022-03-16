package main

import "fmt"

type Run interface {
	Running()
}

type Swim interface {
	Swimming()
}

type Sport interface {
	Run
	Swim
}

func GoSports(s Sport) {
	s.Running()
	s.Swimming()
}

type Boy struct {
	Name string
}

func (b *Boy) Running() {
	fmt.Println(b.Name + "在跑步。。。。")
}

func (b *Boy) Swimming() {
	fmt.Println(b.Name + "在游泳。。。。")
}

func main() {
	b := Boy{Name: "欢喜"}
	GoSports(&b)
}
