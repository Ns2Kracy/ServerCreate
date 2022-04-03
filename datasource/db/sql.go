package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"xorm.io/xorm"
)

var db = NewEngine()

//定时重复执行SQL脚本
func main() {
	sql1()
	go sql2()
}

func sql1() {
	for {
		db.ImportFile("./JSON.sql")
		fmt.Println("执行")
		time.Sleep(5 * time.Second)
	}
}
func sql2() {
	for {
		db.ImportFile("./JSON.sql")
		fmt.Println("执行")
		time.Sleep(1 * time.Second)
	}
}
func NewEngine() *xorm.Engine {
	db, _ := xorm.NewEngine("mysql", "root:nk20021001@tcp(139.224.19.236:3306)/servercreate?charset=utf8")
	return db
}
