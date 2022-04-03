package model

import (
	"time"
	"xorm.io/xorm"
)

// Model 通用
type Model struct {
	Id          uint      `json:"id" xorm:"pk autoincr"`
	CreatedTime time.Time `json:"created_time" xorm:"datetime created"`
	UpdatedTime time.Time `json:"updated_time" xorm:"datetime updated"`
	//Users       Users      `xorm:"extends"`
	//UserStatus  UserStatus `xorm:"extends"'`
	//删除字段不包含在json中
	DeletedAt xorm.AfterDeleteProcessor `json:"deleted_at" xorm:"-"`
}
