package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Cook(){
	now := time.Now()
	fmt.Println("开始做饭")
	time.Sleep(3*time.Second)
	fmt.Println("饭做完了")
	expried:= time.Now().Sub(now)
	fmt.Println(expried)
}
func Buy(){
	fmt.Println( "去买菜")
	time.Sleep(2*time.Second)
	fmt.Println("买菜回来")
}

func Cook2(){
	fmt.Println("开始做饭")
	time.Sleep(3*time.Second)
	fmt.Println("饭做完了")

}

func Eat(){
	fmt.Println("吃饭")
	time.Sleep(1*time.Second)
	fmt.Println("饭吃完了")
}

func Wash(){
	fmt.Println("开始洗碗")
	time.Sleep(1*time.Second)
	fmt.Println("洗碗结束")

}

// type one
func CustomMiddleWare(c *gin.Context){
	now:=time.Now()
	// 权限鉴定
	// if 通过 { c.Next() } else { 要么返回要么报错， 这是业务相关处理 }
	c.Next()
	expried:=time.Now().Sub(now)
	fmt.Println(expried)

}

//  type two
func CustomMiddleWare2() gin.HandlerFunc {
	return func(c *gin.Context){
		now:=time.Now()
		c.Next()
		expried:=time.Now().Sub(now)
		fmt.Printf("耗时：%v \n",expried)
	}
}

func main()  {
	//Cook()
	r:= gin.Default()
	//r.Use(CustomMiddleWare)
	r.Use(CustomMiddleWare2())
	r.GET("/", func(c *gin.Context) {
		Buy()
		Cook2()
		Eat()
		Wash()
	})
	r.Run()
}
