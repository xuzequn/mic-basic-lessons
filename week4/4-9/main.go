package main

import (
	"fmt"
	"github.com/pkg/errors"
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
	//u1 := model.User{}
	//r := db.First(&u1)
	//fmt.Println(u1.ID)
	//fmt.Println(r.Error)
	//fmt.Println(r.RowsAffected)
	//b := errors.Is(r.Error,gorm.ErrRecordNotFound)
	//if b {
	//	fmt.Println("查无此人")
	//} else {
	//	fmt.Println("他是",u1.Name)
	//}

	// take
	// SELECT * FROM `users` LIMIT 1
	//u2 := model.User{}
	//db.Take(&u2)
	//fmt.Println(u2.ID)

	// primarykey
	//u3 := model.User{}
	//r2 := db.First(&u3, 8)
	//fmt.Println(u3.Name)
	//fmt.Println(r2)

	//u3 := model.User{}
	//// SELECT * FROM `users` WHERE `users`.`id` IN (8,9,10) ORDER BY `users`.`id` LIMIT 1
	//r2 := db.First(&u3, 8,9,10)
	//fmt.Println(u3.ID)
	//fmt.Println(r2)

	//var users []model.User
	//// SELECT * FROM `users` WHERE `users`.`id` IN (7,8,9)
	//r3 := db.Find(&users, []int{7,8,9})
	//fmt.Println(r3)
	//for _,user := range users{
	//	fmt.Println(user.Name)
	//}

	// 主键写错成字符串还能查询吗？
	u3 := model.User{}
	r2 := db.First(&u3, "golang")
	fmt.Println(u3.Name)
	fmt.Println(r2.RowsAffected)

	b := errors.Is(r2.Error, gorm.ErrRecordNotFound)
	if b {
		fmt.Println("查无此人")
	} else {
		fmt.Println("他是", u3.Name)
	}

}
