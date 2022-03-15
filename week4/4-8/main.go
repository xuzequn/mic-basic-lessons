package main

import (
	"fmt"
	"gorm.io/gorm"
	"mic-basic-lessons/week4/4-8/model"
)

var db *gorm.DB

func main() {

	users := []model.User{{
		Name: "欢喜",
	}, {
		Name: "面向加薪学习",
	}, {
		Name: "Go语言极简一本通",
	}, {
		Name: "Go语言微服务核心22讲",
	}, {
		Name: "从O到Go语言微服务架构师",
	},
	}
	// type one
	//db.Create(users)

	// type two 每次添加两条
	//db.CreateInBatches(users, 2)

	// type three
	// 值添加了一个值。create 传参是一个 map
	db.Model(&model.User{}).Create([]map[string]interface{}{
		{
			"Name": "欢喜",
		}, {
			"Name": "面向加薪学习",
		}, {
			"Name": "Go语言极简一本通",
		}, {
			"Name": "Go语言微服务核心22讲",
		}, {
			"Name": "从O到Go语言微服务架构师",
		},
	})

	for _, user := range users {
		fmt.Println(user.ID)
	}
}
