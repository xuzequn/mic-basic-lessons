package main

import (
	"bufio"
	"fmt"
	"os"
)

func Cook(){

	// 栈， 先进后出， 在panic 之前执行， 在return 之后执行， 执行完之后 退出
	defer fmt.Println("吃点好的开饭")
	defer fmt.Println("播放音乐")
	fmt.Println("卖菜")
	fmt.Println("卖肉")
	panic("停水了")
	fmt.Println("做饭")
}

func WriteMenu(fileName string, foods []string){
	curDir,err := os.Getwd()
	// f 实现了 writer 接口
	path := curDir+fileName
	f ,err:= os.Create(path)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	w := bufio.NewWriter(f)
	defer  w.Flush()
	for _,item := range foods{
		fmt.Fprintln(w, item)
	}

}

func main()  {

	//defer
	//Cook()
	s:=[]string{"葱烧海参","烧鹅","炭烤生蚝","茶甜鸭","干炒牛河"}
	WriteMenu("/week2/2-8/menu.txt",s)
	
}
