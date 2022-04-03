package controller

import (
	"ServerCreate/middleware"
	"ServerCreate/model"
	"github.com/kataras/iris/v12"
)

const ID = "id"

// 注册
func SignUp(ctx iris.Context) {
	iris.New().Logger().Info(" 用户注册 ")

	//定义请求体
	var userRegister model.Request

	//解码请求体
	if err := ctx.ReadJSON(&userRegister); err != nil {
		ctx.JSON(iris.Map{
			"Status": false,
			"Code":   200,
			"Msg":    "参数校验错误",
			"Data":   nil,
		})
		return
	}

	//判断当前注册的用户名是否存在
	exist, _ := UserService.CheckUserName(userRegister.Name)

	if exist.UserName != "" {
		ctx.JSON(iris.Map{
			"Status": false,
			"Code":   200,
			"Msg":    "用户已存在",
			"Data":   nil,
		})
		return
	}
	iris.New().Logger().Info(userRegister)

	//使用请求的数据建立一个新的用户
	newUser := model.Users{
		UserName:     userRegister.Name,
		UserPassword: userRegister.Password,
	}

	//将注册获取的数据存入数据库以供登录使用
	success := UserService.UserInserter(newUser)
	if !success {
		ctx.JSON(iris.Map{
			"Status": false,
			"Code":   200,
			"Msg":    "注册失败",
			"Data":   nil,
		})
		return
	}

	ctx.JSON(iris.Map{
		"Status": true,
		"Code":   200,
		"Msg":    "注册成功",
		"Data":   nil,
	})
}

// 登录
func SignIn(ctx iris.Context) {
	iris.New().Logger().Info(" 用户登录 ")

	//定义请求体
	var req model.Request

	//解码请求体
	if err := ctx.ReadJSON(&req); err != nil {
		ctx.JSON(iris.Map{
			"Status": false,
			"Code":   200,
			"Msg":    "参数校验错误",
			"Data":   nil,
		})
		return
	}

	//判断请求体是否为空
	if req.Name == "" || req.Password == "" {
		ctx.JSON(iris.Map{
			"Status": false,
			"Code":   200,
			"Msg":    "用户名或密码为空",
			"Data":   nil,
		})
		return
	}

	//请求体log
	iris.New().Logger().Info(req)

	//检查数据库中是否匹配
	user, exist := UserService.CheckNameAndPassword(req.Name, req.Password)
	if !exist {
		ctx.JSON(iris.Map{
			"Status": false,
			"Code":   200,
			"Msg":    "用户名或密码错误",
			"Data":   nil,
		})
		return
	}

	//启用session
	session := middleware.Sess.Start(ctx)
	session.Set("id", int(user.Id))
	ctx.JSON(iris.Map{
		"Status": true,
		"Code":   200,
		"Msg":    "登录成功",
		"Data":   nil,
	})
}

func SignOut(ctx iris.Context) {
	session := middleware.Sess.Start(ctx)
	//删除session
	session.Delete("id")
	ctx.JSON(iris.Map{
		"Status": true,
		"Code":   200,
		"Msg":    "登出",
		"Data":   nil,
	})
}

func UserDetail(ctx iris.Context) {

	session := middleware.Sess.Start(ctx)
	Id, _ := session.GetInt("id")
	userId := uint(ctx.Values().GetIntDefault("id", Id))
	user, err := UserService.CheckUserId(int64(userId))

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
		"Data":   user,
	})
}
