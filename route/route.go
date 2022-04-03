package route

import (
	"ServerCreate/controller"
	"ServerCreate/middleware"
	"github.com/kataras/iris/v12"
)

//TODO:完成路由配置
func InitRoute(app *iris.Application) {
	manage := app.Party("/")
	{
		//登录
		manage.Post("/signin", controller.SignIn)
		//注册
		manage.Post("/signup", controller.SignUp)

		user := manage.Party("/user")
		{
			//用户基础信息
			user.Get("/detail", controller.UserDetail)
			//登出
			user.Post("/signout", controller.SignOut)
		}
		status := manage.Party("/user/status", middleware.ManageAuth)
		{
			//驾驶状态信息
			status.Get("/detail", controller.StatusDetail)
			//更新信息
			status.Post("/update", controller.UpdateStatus)
			//报警
			status.Post("/alarm", controller.Alarm)
		}
		report := manage.Party("/user/status/report", middleware.ManageAuth)
		{
			//HeartRate
			report.Get("/max/heartrate", controller.ReportMaxHeartRate)
			report.Get("/min/heartrate", controller.ReportMinHeartRate)
			report.Get("/average/heartrate", controller.ReportAverageHeartRate)

			//Temperature
			report.Get("/max/temperature", controller.ReportMaxTemperature)
			report.Get("/min/temperature", controller.ReportMinTemperature)
			report.Get("/average/temperature", controller.ReportAverageTemperature)
		}
	}
}
