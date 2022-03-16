package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func CustomMiddleWare(c *gin.Context) {
	now := time.Now()
	name, ok := c.Get("name")
	if !ok {
		//return
		c.Abort()
	}
	fmt.Println(name)
	c.Next()
	expried := time.Now().Sub(now)
	fmt.Println(expried)

}

func main() {
	//Cook()
	r := gin.Default()
	//r.Use(CustomMiddleWare)
	r.Use(CustomMiddleWare)
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
	})
	r.Run()
}
