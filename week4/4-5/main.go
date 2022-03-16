package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"mic-basic-lessons/week4/4-5/model"
	"os"
	"time"
)

var db *gorm.DB

func init() {
	var err error
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值 > 1s
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 禁用彩色打印
		},
	)
	dsn := "root:123456@tcp(127.0.0.1:3306)/orm_test?charset=utf8mb4&parseTime=True&loc=Local"
	// 全局模式
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("连接成功...")
	//err = db.AutoMigrate(&model.Product{})
	//if err != nil {
	//	panic(err)
	//}
	db.AutoMigrate(&model.Food{})
}

func main() {
	//db.Create(&model.Product{
	//	Code: sql.NullString{String: "D666", Valid: true},
	//	Price: 8888,
	//})

	//var p model.Product
	//db.First(&p, 3)
	//fmt.Println(p.Code)
	//fmt.Println(p.Price)
	//
	//db.Model(&p).Updates(model.Product{
	//	Code: sql.NullString{"",true},
	//	Price: 9999,
	//})

}
