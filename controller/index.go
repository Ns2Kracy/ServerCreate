package controller

import (
	"ServerCreate/datasource"
	"ServerCreate/service"
)

var (
	db                = datasource.NewEngine()
	UserService       = service.NewUserService(db)
	UserStatusService = service.NewUserStatusService(db)
)
