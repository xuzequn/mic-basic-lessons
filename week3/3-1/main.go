package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func hello (c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"msg":"面向对象学习-从0到Go语言微服务架构师",
	})
}

func main()  {

	r := gin.Default() // 返回gin engine
	r.GET("/", hello) // 配置get 方法路由
	r.Run(":9090") //8080
	
}
