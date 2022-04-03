package service

import "xorm.io/xorm"

func NewUserStatusService(db *xorm.Engine) UserStatusService {
	return UserStatusService{
		db: db,
	}
}

func NewUserService(db *xorm.Engine) UserService {
	return UserService{
		db: db,
	}
}
