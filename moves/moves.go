package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	// 监听请求
	//app := iris.New()
	err := http.ListenAndServe("http://10.102.117.253:8888/", nil)
	if err != nil {
		log.Println("http server listen :", err)
	}
}

func getJsonTest(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal("parse form error ", err)
	}
	// 初始化请求变量结构
	formData := make(map[string]interface{})
	// 调用json包的解析，解析请求body
	json.NewDecoder(r.Body).Decode(&formData)
	for key, value := range formData {
		log.Println("key:", key, " => value :", value)
	}
}
