package models

import (
	"gorm.io/gorm"
)

type BaseSysUser struct {
	gorm.Model
	DepartmentId int
	Name         string
	UserName     string
	PassWord     string
	PassWordV    int
	NickName     string
	HeadImg      string
	Phone        string
	Email        string
	status       int
	remark       string
}

func (table *BaseSysUser) TableName() string {
	return "base_sys_user"
}
