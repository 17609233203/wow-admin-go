package models

import "gorm.io/gorm"

type BaseSysDepartmentRole struct {
	gorm.Model
	DepartmentId uint `gorm:"comment:'部门id'"`
	RoleId       uint `gorm:"comment:'权限id'"`
}

func (table *BaseSysDepartmentRole) BaseSysDepartmentRole() string {
	return "base_sys_department_role"
}
