package base

import (
	"ginchat/models"
	"ginchat/utils"
	"gorm.io/gorm"
)

func GetUserInfo(user *models.BaseSysUser) *gorm.DB {
	return utils.DB.First(&user)
}
func CreateUser(user *models.BaseSysUser) *gorm.DB {
	return utils.DB.Create(&user)
}
