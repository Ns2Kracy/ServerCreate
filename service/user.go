package service

import (
	"ServerCreate/model"
	"github.com/kataras/iris/v12"
	"xorm.io/xorm"
)

type UserService struct {
	db *xorm.Engine
}

func (us *UserService) CheckNameAndPassword(userName, userPassword string) (model.Users, bool) {
	var user model.Users

	us.db.Table("users").Where(" user_name = ? and user_password = ? ", userName, userPassword).Get(&user)

	return user, user.Id != 0

}

func (us *UserService) CheckUserName(userName string) (model.Users, error) {
	var user model.Users
	if _, err := us.db.Table("users").Where(" user_name = ? ", userName).Get(&user); err != nil {
		return model.Users{}, err
	}
	return user, nil
}
func (us *UserService) CheckUserPassword(userPassword string) (model.Users, error) {
	var user model.Users
	if _, err := us.db.Table("users").Where(" user_password = ? ", userPassword).Get(&user); err != nil {
		return model.Users{}, err
	}
	return user, nil
}

func (us *UserService) CheckUserId(userId int64) (model.Users, error) {
	var user model.Users
	if _, err := us.db.ID(userId).Get(&user); err != nil {
		return model.Users{}, err
	}
	return user, nil
}

func (us *UserService) UserInserter(user model.Users) bool {
	_, err := us.db.Table("users").Insert(&user)
	if err != nil {
		iris.New().Logger().Info(err.Error())
	}
	return err == nil
}

func (us *UserService) UserUpdater(user model.Users) error {
	if _, err := us.db.Table("users").Update(&user); err != nil {
		return err
	}
	return nil
}
