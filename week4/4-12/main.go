package main

import (
	"database/sql"
	"fmt"
	"log"
	"mic-basic-lessons/week4/4-5/model"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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

}

func AddProduct() {
	p1 := model.Product{
		Code: sql.NullString{
			String: "Go语言极简一本通",
			Valid:  true,
		},
		Price: 99,
	}

	p2 := model.Product{
		Code: sql.NullString{
			String: "Go语言微服务架构22讲",
			Valid:  true,
		},
		Price: 99,
	}

	var productList []model.Product

	productList = append(productList, p1)
	productList = append(productList, p2)

	db.Create(productList)
}

func main() {
	// delete
	// soft delete

	var p model.Product
	db.First(&p, 4)
	//UPDATE `products` SET `deleted_at`='2022-03-16 23:24:01.977' WHERE `products`.`id` = 4 AND `products`.`deleted_at` IS NULL
	db.Delete(&p)

	// UPDATE `products` SET `deleted_at`='2022-03-16 23:25:48.603' WHERE price=777 AND `products`.`deleted_at` IS NULL
	// UPDATE `products` SET `deleted_at`='2022-03-16 23:31:12.532' WHERE price=9999 AND `products`.`deleted_at` IS NULL
	db.Where("price=?", 9999).Delete(&model.Product{})

}
