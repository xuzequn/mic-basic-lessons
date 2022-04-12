package main

import (
	"go.uber.org/zap"
	"time"
)

func Newlogger() (*zap.Logger, error) {
	config := zap.NewProductionConfig()
	config.OutputPaths = append(config.OutputPaths, "./7-2.log")
	return config.Build() // Logger, err
}

func main() {
	logger, err := Newlogger()
	if err != nil {
		panic(err)
	}
	sugar := logger.Sugar()
	defer sugar.Sync()
	e := "出错了"
	sugar.Infof("error:%s", e)
	sugar.Info("又出错了",
		"err", e,
		"time", time.Now().Unix(),
	)

}
