package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"mic-basic-lessons/week3/3-4/model"
	"net/http"
)

func main() {
	r := gin.Default()
	productGroup := r.Group("/product")
	{
		productGroup.GET("/detail", defaultHandler)
		productGroup.POST("/add", addHandler)
		productGroup.POST("/add/json", addJsonHandler)
	}
	r.Run()
}

func addJsonHandler(c *gin.Context) {
	var p model.Product
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "解析参数错误",
		})
		return
	}
	//id := rand.Intn(10000)
	c.JSON(http.StatusOK, gin.H{
		"id":   p.ID,
		"name": p.Name,
	})

}

func addHandler(c *gin.Context) {
	id := rand.Intn(10000)
	name := c.DefaultPostForm("name", "fromDefaultName")
	c.JSON(http.StatusOK, gin.H{
		"id":   id,
		"name": name,
	})
}

func defaultHandler(c *gin.Context) {
	id := c.DefaultQuery("id", "0")
	name := c.DefaultQuery("name", "defaultName")
	c.JSON(http.StatusOK, gin.H{
		"id":   id,
		"name": name,
	})

}
