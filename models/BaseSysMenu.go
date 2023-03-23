package models

import (
	"gorm.io/gorm"
)

type BaseSysMenu struct {
	gorm.Model
	ParentId uint   `gorm:"comment:'父id'"`
	Name     string `gorm:"comment:'菜单名称'"`
	Router   string `gorm:"comment:'路由'"`
	Perms    string `gorm:"comment:'权限标识'"`
	Types    int    `gorm:"comment:'类型：0目录1菜单2按钮'"`
	Icon     string `gorm:"comment:'icon'"`
	OrderNum int    `gorm:"comment:'排序'"`
	ViewPath string `gorm:"comment:'视图地址'"`
	IsShow   int    `gorm:"comment:'是否显示'"`
}

func (table *BaseSysMenu) BaseSysMenu() string {
	return "base_sys_menu"
}
