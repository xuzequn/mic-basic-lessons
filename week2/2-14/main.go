package main

import (
	"fmt"
	"sync"
)

func main()  {
	// sync.waitGroup Add(n) Done() -1 Wait() n>0， 内部有一个计数器

	s:= []string {"葱烧海参", "烧鹅","生蚝","鲍鱼","茶甜鸭","干炒牛河"}
	var wg sync.WaitGroup
	for _, item := range (s){
		wg.Add(1)
		go SayFoodName(item, &wg)
	}
	wg.Wait()
	fmt.Println("您点的菜上齐了")


}

func SayFoodName(name string, wg *sync.WaitGroup) {
	fmt.Println("您点的菜："+ name)
	wg.Done()
	
}
