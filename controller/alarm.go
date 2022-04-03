package controller

import (
	"ServerCreate/middleware"
	"github.com/kataras/iris/v12"
)

func Alarm(ctx iris.Context) {
	//TODO:调用平台API实现短信报警，需判断要报警的情况
	session := middleware.Sess.Start(ctx)
	session.Get("Id")

}
