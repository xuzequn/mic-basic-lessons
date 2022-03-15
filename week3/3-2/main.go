package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func hello (c *gin.Context){
	panic("程序错了")
	c.JSON(http.StatusOK,gin.H{
		"msg":"面向对象学习-从0到Go语言微服务架构师",
	})
}

func main()  {

	//r := gin.Default() // 返回gin engine, 使用 loger 和 recover 中间件 ，use 使用中间件
	r:=gin.New()// 没有使用中间件， 直接创建一个gin engine
	r.GET("/", hello) // 配置get 方法路由
	r.Run(":9090") //8080

}
