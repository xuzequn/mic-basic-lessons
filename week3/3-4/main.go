package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mic-basic-lessons/week3/3-4/model"
	"net/http"
)

func main()  {
	r := gin.Default()
	// 商品
	productGroup := r.Group("/product")
	{
		productGroup.GET("/productlist", productList)
		//productGroup.GET("/product/1", productDetail1)
		//productGroup.GET("/product/2", productDetail2)
		productGroup.GET("/:id/:name", productDetail3)
		productGroup.GET("/file/*all", fileHeader) // 处理url 后面所有的请求
	}

	r.Run()
}

func productDetail3(c *gin.Context) {
	var p model.Product
	err := c.ShouldBindUri(&p)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":"输入有误",
		})
	}else {
		c.JSON(http.StatusOK, gin.H{
			"msg":fmt.Sprintf("%d-%s",p.ID,p.Name),
		})
	}
}

func fileHeader(c *gin.Context) {
	all:=c.Param("all")
	c.JSON(http.StatusOK,gin.H{
		"msg":all,
	})
}

func productDetail(c *gin.Context) {
	id := c.Param("id")
	name := c.Param("name")
	if id == ""|| name==""{
		c.JSON(http.StatusBadRequest,gin.H{
			"msg":"参数请求错误",
		})
	} else{
		c.JSON(http.StatusOK,gin.H{
			"msg":id+"-"+name,
		})
	}
}


func productList(c *gin.Context) {

}

