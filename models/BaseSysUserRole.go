package models

import "gorm.io/gorm"

type BaseSysUserRole struct {
	gorm.Model
	UserId uint `gorm:"comment:'用户id'"`
	RoleId uint `gorm:"comment:'权限id'"`
}

func (table *BaseSysUserRole) BaseSysUserRole() string {
	return "base_sys_user_role"
}
