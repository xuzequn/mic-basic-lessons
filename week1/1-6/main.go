package main

import (
	"fmt"
	"net/http"
)

func GoToHappy(sunSine bool){

	if sunSine {
		fmt.Println("出去嗨皮，出去浪")
	} else {
		fmt.Println("从0到Go语言架构师-路线图")
	}

}

func VisitUrl(url string) (int, error) {
	res, err := http.Get(url)
	defer res.Body.Close()
	if err != nil {
		return 0, err
	} else {
		return res.StatusCode,nil
	}
}

func SwitchShow(url string){
	if code, err := VisitUrl(url); err != nil {
		fmt.Println(err)
	} else {
		switch code {
		case 200:
			fmt.Println("\n请求成功")
			break // failthough
		case 404:
			fmt.Println("\n网址错误")
		default:
			panic("未知错误")
		}
	}
}

func SwitchShow2(url string){
	if code, err := VisitUrl(url); err != nil {
		fmt.Println(err)
	} else {
		switch {
		case code==200:
			fmt.Println("\n请求成功")
			break // failthough
		case code==404:
			fmt.Println("\n网址错误")
		default:
			panic("未知错误")
		}
	}
}

func main(){
	//VisitUrl("http://www.baidu.com")
	SwitchShow("http://www.baidu.com")
}
