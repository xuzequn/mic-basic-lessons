package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mic-basic-lessons/week3/3-6/model"
	"net/http"
)

func main() {
	r := gin.Default()
	r.POST("login", loginHandler)
	r.Run()
}

//c.MustBindWith()
func loginHandler(c *gin.Context) {
	var login model.Accountlogin
	err := c.ShouldBind(&login)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "输入错误",
		})
		return
	}
	//c.MustBindWith()
	c.JSON(http.StatusOK, gin.H{
		"msg": fmt.Sprintf("%s-%s", login.AccountName, login.Password),
	})

}
