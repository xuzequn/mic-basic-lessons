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

func init(){
	var err error
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值 > 1s
			LogLevel:      logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,   // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:      true,         // 禁用彩色打印
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

func main()  {
	// 面向对象的查询方式
	//u1 := model.User{}
	// SELECT * FROM `users` WHERE name='欢喜' ORDER BY `users`.`id` LIMIT 1
	//db.Where("name=?","欢喜").First(&u1)
	//fmt.Println(u1.ID)

	//u2 := model.User{}
	////  SELECT * FROM `users` WHERE `users`.`name` = '欢喜' ORDER BY `users`.`id` LIMIT 1
	//db.Where(model.User{Name: "欢喜"}).First(&u2)
	//fmt.Println(u2.ID)

	// 不等于的查询条件
	//var user []model.User
	//db.Where("name <> ?","欢喜").Find(&user)
	//
	//for _,item:= range user{
	//	fmt.Println( item.ID)
	//}

	// 结构作为查询参数
	//var userlist []model.User
	//db.Where(map[string]interface{}{"name":"欢喜"}).Find(&userlist)
	//for _,item := range userlist{
	//	fmt.Println(item.ID)
	//}

	// 结构作为查询参数，gorm 只查询非零值的字段条件，通过map去避免
	var userlist2 []model.User
	db.Where(map[string]interface{}{
		"name":"欢喜",
		"age": 0,
	}).Find(&userlist2)
	for _,item := range userlist2{
		fmt.Println(item.ID)
	}


}
