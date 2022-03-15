package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"mic-basic-lessons/week4/4-3/model"
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
	//err = db.AutoMigrate(&model.Product{})
	//if err != nil {
	//	panic(err)
	//}
}

func main()  {
	//db.Create(&model.Product{Code: "D42", Price: 200})

	var p model.Product
	// SELECT * FROM `products` WHERE `products`.`id` = 1 AND `products`.`deleted_at` IS NULL ORDER BY `products`.`id` LIMIT 1
	//db.First(&p,1)
	//fmt.Println(p.Code)

    // SELECT * FROM `products` WHERE code='D42' AND `products`.`deleted_at` IS NULL ORDER BY `products`.`id` LIMIT 1
	//db.First(&p, "code=?","D42")
	//fmt.Println(p.Price)

	// UPDATE `products` SET `price`=300,`updated_at`='2022-03-11 20:13:30.912' WHERE `products`.`deleted_at` IS NULL AND `id` = 1
	//db.Model(&p).Update("price",300)
	//db.First(&p, "code=?","D42")
	//fmt.Println(p.Price)

    // SELECT * FROM `products` WHERE `products`.`id` = 1 AND `products`.`deleted_at` IS NULL ORDER BY `products`.`id` LIMIT 1
	//db.First(&p, 1)
    //  UPDATE `products` SET `updated_at`='2022-03-11 21:20:30.769',`code`='FF42',`price`=600 WHERE `products`.`deleted_at` IS NULL AND `id` = 1
	//db.Model(&p).Updates(model.Product{
	//	Code: "FF42",
	//	Price: 600,
	//})

	//SELECT * FROM `products` WHERE `products`.`id` = 1 AND `products`.`deleted_at` IS NULL AND `products`.`id` = 1 ORDER BY `products`.`id` LIMIT 1
	//db.First(&p,1)
	//fmt.Println(p.Code)
	//fmt.Println(p.Price)

	// 零值 string "" int 0
	// 数据被保护了
	//db.First(&p, 1)
	//UPDATE `products` SET `updated_at`='2022-03-11 21:25:01.587' WHERE `products`.`deleted_at` IS NULL AND `id` = 1
	//db.Model(&p).Updates(model.Product{
	//	Code: "",
	//	Price: 0,
	//})

	//db.First(&p,1)
	////UPDATE `products` SET `code`='FF42',`price`=0,`updated_at`='2022-03-11 21:27:52.853' WHERE `products`.`deleted_at` IS NULL AND `id` = 1
	//db.Model(&p).Updates(map[string]interface{}{
	//	"Code": "FF42",
	//	"Price": 0,
	//})
	db.First(&p,1)
	fmt.Println(p.Code)
	fmt.Println(p.Price)

	//delete 软删除 flag status， 0 或1 ，当删除的时候就将flag update 为0 ，然后查询的时候选择 flag = 1 。
	// 商城后台是有权限看到任何记录的。 delete_at 填充值不为null
	// UPDATE `products` SET `deleted_at`='2022-03-11 21:36:49.742' WHERE `products`.`id` = 1 AND `products`.`id` = 1 AND `products`.`deleted_at` IS NULL
	db.Delete(&p,1)





}
