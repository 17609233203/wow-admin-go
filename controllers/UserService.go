package controllers

//import (
//	"ginchat/models"
//	"github.com/gin-gonic/gin"
//)
//
//// User
//// @Tags 用户列表
//// @Summary	查看用户
//// @Success 200 {string} json{"code","data"}
//// @router /user/getUserList [get]
//func User(c *gin.Context) {
//	data := models.GetUserList()
//	c.JSON(200, gin.H{
//		"data": data,
//	})
//}
//
//type Ireq struct {
//	Name       string `json:"name"`
//	PassWord   string `json:"passWord"`
//	RePassWord string `json:"rePassWord"`
//}
//
//// CreateUser
//// @Tags	用户模块
//// @Summary	新增用户
//// @Param	name rawData string false "用户名"
//// @Param	passWord rawData string false "密码"
//// @Param	rePassWord rawData string false "确认密码"
//// @Success 200	{string} json{"code","message"}
//// @router	/user/createUser [post]
//func CreateUser(c *gin.Context) {
//	user := models.UserBasic{}
//	req := &Ireq{}
//
//	if err := c.BindJSON(&req); err != nil {
//		c.JSON(504, gin.H{
//			"message": "出错了",
//			"err":     err,
//		})
//	}
//
//	if req.PassWord != req.RePassWord {
//		c.JSON(-1, gin.H{
//			"message": "两次密码不一致",
//		})
//	}
//	user.PassWord = req.PassWord
//	//models.CreateUser(user)
//	c.JSON(200, req)
//}
