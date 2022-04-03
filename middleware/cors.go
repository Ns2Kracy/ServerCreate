package middleware

import (
	"github.com/kataras/iris/v12"
)

//CrsAuth 设置跨域
func CrsAuth(ctx iris.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	if ctx.Request().Method == "OPTIONS" {
		ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Api, Accept, Authorization, Version, Token")
		ctx.StatusCode(204)
		return
	}
	ctx.Next()
}
