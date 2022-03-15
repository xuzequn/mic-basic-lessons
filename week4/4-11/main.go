package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"mic-basic-lessons/week4/4-8/model"
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

	//单点更新 db.Model().Update()
	//  UPDATE `users` SET `age`=18,`updated_at`='2022-03-14 22:52:42.645' WHERE name='欢喜'
	// ErrMissingWhereClause,更新一定要带着where 条件，不然gorm 不给执行。
	//db.Model(&model.User{}).Where("name=?", "欢喜").Update("age", 18)

	//u1 := model.User{}
	//db.First(&u1, 7)
	//fmt.Println(u1.Name)
	//db.Model(&u1).Update("email", "abc@qq.com")

	//u2 := model.User{}
	//db.First(&u2, 9)
	//// UPDATE `users` SET `age`=18,`updated_at`='2022-03-14 22:59:02.374' WHERE name='Go语言极简一本通' AND `id` = 9
	//db.Model(&u2).Where("name=?", "Go语言极简一本通").Update("age", 18)

	// 多点更新 db.Model().Updates()
	//u3 := model.User{}
	//db.First(&u3, 9)
	//
	// UPDATE `users` SET `name`='吃货我来了',`age`=17,`updated_at`='2022-03-14 23:04:43.62' WHERE `id` = 9
	//db.Model(&u3).Updates(model.User{Name: "吃货我来了", Age: 17})

	//用 map
	//u4 := model.User{}
	//db.First(&u4, 9)
	//  UPDATE `users` SET `age`=18,`name`='面向加薪学习',`updated_at`='2022-03-14 23:07:13.698' WHERE `id` = 9
	//db.Model(&u4).Updates(map[string]interface{}{
	//	"name": "面向加薪学习",
	//	"age":  18,
	//})

	// Select  更新时，Select 了一个值，updates 两个值，但是只更新select 选择的的值
	//u5 := model.User{}
	//db.First(&u5, 9)
	//db.Model(&u5).Select("name").Updates(map[string]interface{}{
	//	"name": "吃货",
	//	"age":  17,
	//})

	// Omit忽略，更新时会忽略Omit 选择的值
	//u6 := model.User{}
	//db.First(&u6, 9)
	// UPDATE `users` SET `age`=17,`updated_at`='2022-03-14 23:12:04.639' WHERE `id` = 9
	//db.Model(&u6).Omit("name").Updates(map[string]interface{}{
	//	"name": "吃货",
	//	"age":  17,
	//})

	u7 := model.User{}
	db.First(&u7, 9)
	// UPDATE `users` SET `age`=0,`name`='吃货',`updated_at`='2022-03-14 23:14:57.458' WHERE `id` = 9
	db.Model(&u7).Select("name", "age").Updates(map[string]interface{}{
		"name": "吃货",
		"age":  0,
	})

}
