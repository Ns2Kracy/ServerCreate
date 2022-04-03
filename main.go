package main

import (
	"ServerCreate/route"

	"github.com/kataras/iris/v12"
)

func main() {
	app := newApp()

	err := app.Run(iris.Addr(":8080"), iris.WithOptimizations)

	if err != nil {
		return
	}
}

func newApp() *iris.Application {
	app := iris.New()
	//设置日志级别  开发阶段为debug
	app.Logger().SetLevel("debug")
	route.InitRoute(app)
	return app

}
