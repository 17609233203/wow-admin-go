package models

import (
	"gorm.io/gorm"
)

type BaseSysUser struct {
	gorm.Model
	DepartmentId int    `gorm:"comment:'部门ID'"`
	Name         string `gorm:"comment:'姓名'"`
	UserName     string `gorm:"comment:'用户名'"`
	PassWord     string `gorm:"comment:'密码'"`
	PassWordV    int    `gorm:"comment:'密码版本，改完密码使原来密码失效'"`
	NickName     string `gorm:"comment:'昵称'"`
	HeadImg      string `gorm:"comment:'头像'"`
	Phone        string `gorm:"comment:'手机号'"`
	Email        string `gorm:"comment:'邮箱'"`
	status       int    `gorm:"comment:'状态'"`
	remark       string `gorm:"comment:'备注'"`
}

func (table *BaseSysUser) BaseSysUser() string {
	return "base_sys_user"
}
