package controller

import (
	"ServerCreate/middleware"
	"ServerCreate/model"
	"github.com/kataras/iris/v12"
)

// UpdateStatus 更新当前的用户状态
func UpdateStatus(ctx iris.Context) {
	var status model.UserStatus

	//校验登录，
	session := middleware.Sess.Start(ctx)
	session.Get("Id")

	//更新数据
	err := UserStatusService.StatusUpdater(status)
	if err != nil {
		ctx.JSON(iris.Map{
			"Status": false,
			"Code":   200,
			"Msg":    "信息更新失败",
			"Data":   nil,
		})
	}
	ctx.JSON(iris.Map{
		"Status": true,
		"Code":   200,
		"Msg":    "更新成功",
		"Data":   status,
	})
}

// StatusDetail 获取当前的用户，显示当前用户的状态
func StatusDetail(ctx iris.Context) {
	statusByte := middleware.Sess.Start(ctx).Get("Id")

	if statusByte == nil {
		ctx.JSON(iris.Map{
			"Status": false,
			"Code":   200,
			"Msg":    "未登录",
			"Data":   nil,
		})
		return
	}

	statusId := ctx.Values().GetIntDefault("Id", 0)

	status, err := UserStatusService.CheckStatusId(uint(statusId))
	if err != nil {
		ctx.JSON(iris.Map{
			"Status": false,
			"Code":   200,
			"Msg":    "用户不存在",
			"Data":   nil,
		})

	}
	ctx.JSON(iris.Map{
		"Status": true,
		"Code":   200,
		"Msg":    "信息",
		"Data":   status,
	})
}
