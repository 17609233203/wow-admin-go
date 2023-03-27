package base

import (
	"ginchat/models"
	"ginchat/service/base"
	"ginchat/utils"
	"github.com/gin-gonic/gin"
)

type ILoginReq struct {
	UserName string `form:"userName" binding:"required"`
	PassWord string `form:"passWord" binding:"required"`
}

// Login
// @Tags	登录模块
// @Summary	用户登录
// @Param	userName rawData string false "用户名"
// @Param	passWord rawData string false "密码"
// @Success 200	{string} json{"code","message"}
// @router	/api/admin/base/open/login [post]
func Login(c *gin.Context) {
	req := &ILoginReq{}
	if c.ShouldBind(&req) != nil {
		utils.ServerError(c, "获取参数失败")
	} else {
		req.PassWord = utils.Encryption(req.PassWord)
		user := models.BaseSysUser{UserName: req.UserName, PassWord: req.PassWord}
		base.GetUserInfo(&user)
		token, err := utils.GenerateToken(user.ID)
		if err != nil {
			utils.ServerError(c, "生成token失败")
		} else {
			utils.Success(c, gin.H{
				"token": token,
			})
		}

	}
}
func CreateUser(c *gin.Context) {
	user := models.BaseSysUser{
		UserName: "admin",
		PassWord: "e10adc3949ba59abbe56e057f20f883e",
		NickName: "超级管理员",
		Name:     "吴迪",
		Phone:    "17609233203",
		Email:    "1103733590@qq.com",
	}
	base.CreateUser(&user)
	utils.Success(c, gin.H{
		"data": "成功",
	})
}
