package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"mic-basic-lessons/week4/4-15/model"
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
	//db.AutoMigrate(&model.CreditCard{})

}

func main() {

	// has many
	var e1 model.Employer
	db.First(&e1)
	//for i := 0; i < 5; i++ {
	//	id := uuid.New()
	//	c := model.CreditCard{
	//		Number:     id.String(),
	//		EmployerID: e1.ID,
	//	}
	//	db.Create(&c)
	//
	//}

	//type one
	// SELECT * FROM `credit_cards` WHERE `credit_cards`.`employer_id` = 1 AND `credit_cards`.`deleted_at` IS NULL
	// SELECT * FROM `employers` WHERE `employers`.`deleted_at` IS NULL AND `employers`.`id` = 1 ORDER BY `employers`.`id` LIMIT 1
	//db.Preload("CreditCards").First(&e1)
	//for _, card := range e1.CreditCards {
	//	fmt.Println(card.Number)
	//}

	//type two
	//SELECT * FROM `credit_cards` WHERE `credit_cards`.`employer_id` = 1 AND `credit_cards`.`deleted_at` IS NULL
	var creditCards []model.CreditCard
	db.Model(&e1).Association("CreditCards").Find(&creditCards)
	for _, card := range creditCards {
		fmt.Println(card.Number)
	}
}
