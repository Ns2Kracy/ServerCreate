package middleware

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

var Sess = sessions.New(sessions.Config{Cookie: "sessions"})

func Auth(ctx iris.Context) {
	//检查登录状态
	session := Sess.Start(ctx)
	userId := session.GetIntDefault("id", 1)
	ctx.Values().Set("id", userId)
	ctx.ViewData("id", userId)

	ctx.Next()
}

func ManageAuth(ctx iris.Context) {
	//后台管理强制要求登录
	Auth(ctx)
	adminId := ctx.Values().GetIntDefault("id", 0)
	if adminId == 0 {
		ctx.JSON(iris.Map{
			"Status": false,
			"Code":   200,
			"Msg":    "未登录，请登录",
			"Data":   nil,
		})
		return
	}

	ctx.Next()
}
