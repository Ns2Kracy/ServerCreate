package model

// Users 用户基本信息
type Users struct {
	Model `xorm:"extends"`
	//UserId       uint       `xorm:"pk" json:"id"`
	UserName     string `xorm:"varchar(255)" json:"userName"`
	UserPassword string `xorm:"varchar(255)" json:"userPassword"`
}
