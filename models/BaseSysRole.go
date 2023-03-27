package models

import "gorm.io/gorm"

type BaseSysRole struct {
	gorm.Model
	UserId    uint   `gorm:"comment:'用户id'"`
	Name      string `gorm:"comment:'角色名称'"`
	Label     string `gorm:"comment:'角色标签'"`
	Remark    string `gorm:"comment:'备注'"`
	Relevance int    `gorm:"comment:'数据权限是否关联上下级'"`
}

func (table *BaseSysRole) BaseSysRole() string {
	return "base_sys_role"
}
