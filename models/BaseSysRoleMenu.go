package models

import "gorm.io/gorm"

type BaseSysRoleMenu struct {
	gorm.Model
	RoleId uint `gorm:"comment:'角色id'"`
	MenuId uint `gorm:"comment:'菜单id'"`
}

func (table *BaseSysRoleMenu) BaseSysRoleMenu() string {
	return "base_sys_role_menu"
}
