package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"mic-basic-lessons/week4/4-13/model"
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

}

func main() {
	//var e1 model.Employer
	//db.First(&e1)
	//fmt.Println(e1.Name)
	//fmt.Println(e1.CompanyID)
	//fmt.Println(e1.Company.Name)

	//SELECT * FROM `companies` WHERE `companies`.`id` = 1 AND `companies`.`deleted_at` IS NULL
	//SELECT * FROM `employers` WHERE `employers`.`deleted_at` IS NULL ORDER BY `employers`.`id` LIMIT 1
	//db.Preload("Company").First(&e1)
	//fmt.Println(e1.Name)
	//fmt.Println(e1.CompanyID)
	//fmt.Println(e1.Company.Name)

	var e2 model.Employer
	// SELECT `employers`.`id`,`employers`.`created_at`,`employers`.`updated_at`,`employers`.`deleted_at`,`employers`.`name`,`employers`.`company_id`,`Company`.`id` AS `Company__id`,`Company`.`created_at` AS `Company__created_at`,`Company`.`updated_at` AS `Company__updated_at`,`Company`.`deleted_at` AS `Company__deleted_at`,`Company`.`name` AS `Company__name` FROM `employers` LEFT JOIN `companies` `Company` ON `employers`.`company_id` = `Company`.`id` WHERE `employers`.`deleted_at` IS NULL ORDER BY `employers`.`id` LIMIT 1s
	db.Joins("Company").First(&e2, 4)
	fmt.Println(e2.Name)
	fmt.Println(e2.CompanyID)
	fmt.Println(e2.Company.Name)

}
