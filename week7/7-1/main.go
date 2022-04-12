package main

import (
	"go.uber.org/zap"
	"time"
)

//zap

func main() {
	//zap.NewDevelopment()
	pro, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	// 最后日志落盘
	defer pro.Sync()
	// 写日志的方法
	sugar := pro.Sugar()
	e := "出错了"
	sugar.Infof("error:%s", e)
	sugar.Info("又出错了",
		"err",
		e,
		"time:", time.Now().Unix(),
	)
	// zap.Newproduction.logger 是 suger 的4～ 10倍， suger 使用了 golang的反射
	pro.Info("loger info test...",
		zap.String("name err", e),
		zap.Int("price err", 123),
	)
}
