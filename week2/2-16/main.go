package main

import (
	"fmt"
	"sync"
	"time"
)

type Goods struct {
	v map[string]int
	m sync.Mutex // 互斥锁
	// sync.RWMutex
}

func (g *Goods) Inc(key string, num int) {
	g.m.Lock()
	defer g.m.Unlock()
	fmt.Printf("库存数量增加，已加锁\n")
	g.v[key]++
	fmt.Printf("库存数量增加完毕，解锁\n")
}

func (g *Goods) Value(key string) int {
	g.m.Lock()
	defer g.m.Unlock()
	fmt.Println("库存值上锁\n")
	return g.v[key]
}

func main() {
	mutext := sync.Mutex{}
	g := Goods{
		v: make(map[string]int),
		m: mutext,
	}
	for i := 0; i < 10; i++ {
		go g.Inc("榴莲", i)

	}
	time.Sleep(1 * time.Second)
	fmt.Println(g.Value("榴莲"))
}
