package models

import "gorm.io/gorm"

type BaseSysDepartment struct {
	gorm.Model
	Name     string `gorm:"comment:'部门名称'"`
	ParentId uint   `gorm:"comment:'上级部门Id'"`
	OrderBy  int    `gorm:"comment:'排序'"`
}

func (table *BaseSysDepartment) BaseSysDepartment() string {
	return "base_sys_department"
}
