package controller

import (
	"ServerCreate/middleware"
	"ServerCreate/model"
	"github.com/kataras/iris/v12"
)

// Report TODO:生成当前状态的健康报表
func Report(ctx iris.Context) {
	//获取当前用户
	session := middleware.Sess.Start(ctx)
	session.Get("id")
	statusId := ctx.Values().GetIntDefault("id", 0)

	status, err := UserStatusService.CheckStatusId(uint(statusId))

	if err != nil {
		ctx.JSON(iris.Map{
			"Status": false,
			"Code":   200,
			"Msg":    "不存在",
			"Data":   nil,
		})
		return
	}

	ctx.JSON(iris.Map{
		"Status": true,
		"Code":   200,
		"Msg":    "健康报表",
		"Data":   status,
	})
}

/*
 * 心跳相关
 */

func ReportMaxHeartRate(ctx iris.Context) {
	session := middleware.Sess.Start(ctx)
	session.Get("id")
	var status model.UserStatus
	MaxHeartRate := UserStatusService.MaxHeartRateReport(status.HeartRate)

	ctx.JSON(iris.Map{
		"Status": true,
		"Code":   200,
		"Msg":    "最大心跳",
		"Data":   MaxHeartRate,
	})
	iris.New().Logger().Info(MaxHeartRate)
}

func ReportMinHeartRate(ctx iris.Context) {
	session := middleware.Sess.Start(ctx)
	session.Get("id")
	var status model.UserStatus
	MinHeartRate := UserStatusService.MaxHeartRateReport(status.HeartRate)

	ctx.JSON(iris.Map{
		"Status": true,
		"Code":   200,
		"Msg":    "最小心跳",
		"Data":   MinHeartRate,
	})
	iris.New().Logger().Info(MinHeartRate)

}

func ReportAverageHeartRate(ctx iris.Context) {
	session := middleware.Sess.Start(ctx)
	session.Get("id")
	var status model.UserStatus
	AverageHeartRate := UserStatusService.AverageHeartReport(status.HeartRate)
	ctx.JSON(iris.Map{
		"Status": true,
		"Code":   200,
		"Msg":    "平均心跳",
		"Data":   AverageHeartRate,
	})
	iris.New().Logger().Info(AverageHeartRate)

}

/*
 * 体温相关
 */

func ReportMaxTemperature(ctx iris.Context) {
	session := middleware.Sess.Start(ctx)
	session.Get("id")

	var status model.UserStatus
	MaxTemperature := UserStatusService.MaxTemperatureReport(status.Temperature)
	ctx.JSON(iris.Map{
		"Status": true,
		"Code":   200,
		"Msg":    "最大体温",
		"Data":   MaxTemperature,
	})
	iris.New().Logger().Info(MaxTemperature)

}

func ReportMinTemperature(ctx iris.Context) {
	session := middleware.Sess.Start(ctx)
	session.Get("id")

	var status model.UserStatus
	MinTemperature := UserStatusService.MinTemperatureReport(status.Temperature)
	ctx.JSON(iris.Map{
		"Status": true,
		"Code":   200,
		"Msg":    "最小体温",
		"Data":   MinTemperature,
	})
	iris.New().Logger().Info(&MinTemperature)

}

func ReportAverageTemperature(ctx iris.Context) {
	session := middleware.Sess.Start(ctx)
	session.Get("id")

	var status model.UserStatus
	AverageTemperature := UserStatusService.AverageTemperatureReport(status.Temperature)
	ctx.JSON(iris.Map{
		"Status": true,
		"Code":   200,
		"Msg":    "最大体温",
		"Data":   AverageTemperature,
	})
	iris.New().Logger().Info(AverageTemperature)

}
