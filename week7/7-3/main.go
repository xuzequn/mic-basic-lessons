package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

//func main() {
//	//debug info warn error fatal 逐渐升级
//	r := gin.Default()
//	//pro, _ := zap.NewProduction()
//	//defer pro.Sync()
//	//sugar := pro.Sugar()
//	zap.S()
//	port := 9090
//	zap.S().Infof("服务器启动端口是%d", port)
//	r.GET("/", func(context *gin.Context) {
//		context.JSON(http.StatusOK, gin.H{
//			"msg": "ok",
//		})
//	})
//	err := r.Run(fmt.Sprintf(":%d", port))
//	if err != nil {
//		zap.S().Panic("服务器启动失败", err.Error())
//	}
//}

func main() {
	r := gin.Default()
	port := 9090
	pro, _ := zap.NewProduction()
	zap.ReplaceGlobals(pro)
	s := zap.S()
	defer s.Sync()
	r.GET("/", func(context *gin.Context) {
		s.Info("调用请求成功")
		context.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
	})
	err := r.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		zap.S().Panic("服务器启动失败", err.Error())
	}
}
