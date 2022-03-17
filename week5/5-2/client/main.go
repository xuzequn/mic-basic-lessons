package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Result struct {
	Msg string `json:"msg"`
}

func GetUrl(url string) (string, error) {
	r, err := http.Get(url)
	defer r.Body.Close()
	if err != nil {
		panic(err)
	}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", err
	}
	var result Result
	err = json.Unmarshal(b, &result)
	if err != nil {
		return "", err
	}
	return result.Msg, nil
}

func main() {
	result, err := GetUrl("http://127.0.0.1:8080/hello?lesson=从0到Go语言微服务架构师")
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
