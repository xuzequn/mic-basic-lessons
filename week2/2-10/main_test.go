package main

import (
	"testing"
)

func TestShowWithTable(t *testing.T){
	tests := []struct{
		name string
		want string
	}{
		{
			name: "面向加薪学习-从0到Go语言微服务架构师",
			want: "面向加薪学习-从0到Go语言微服务架构师",
		},
		{
			name: "面向加薪学习-从0到Go语言微服务架构师",
			want: "面向加薪学习-从0到Go语言微服务架构师2",
		},
	}
	for _,item := range(tests){
		t.Run(item.name, func(t *testing.T){
			if got:=Show();got != item.want{
				t.Errorf("Show() = %v, want = %v", got,item.want)
			}
		})
		}
}


func TestShow(t *testing.T) {
	result := Show()
	want := "面向加薪学习-从0到Go语言微服务架构师"
	if result == want{
		t.Logf("Show()= %v, want= %v", result, want)
	} else {
		t.Errorf("Show()= %v, want= %v", result, want)
	}
}

