package main

import (
	"github.com/gin-gonic/gin"
	"mic-basic-lessons/week3/3-3/model"
	"net/http"
)

func main()  {
	r := gin.Default()
	// 账户
	accountGroup := r.Group("/account")
	{
		accountGroup.GET("/list", accountList)
		accountGroup.POST("/add",addAccount)
	}
	// 商品
	productGroup := r.Group("/product")
	{
		productGroup.GET("/productlist", productList)
	}
	// 订单
	// 。。。
	r.Run()
}

func productList(context *gin.Context) {
	
}

func addAccount(context *gin.Context) {
	
}

func accountList(c *gin.Context) {
	var accountList []model.Account
	a1 := model.Account{
		No: 1,
		Name: "老王",
	}
	accountList = append(accountList, a1)
	a2 := model.Account{
		No: 2,
		Name: "老张",
	}
	accountList = append(accountList, a2)
	c.JSON(http.StatusOK,gin.H{
		"accountList": accountList,
	})
}
