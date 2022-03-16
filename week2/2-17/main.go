package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	//依赖管理
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
	})
	r.Run()
}
