package router

import (
	"ginchat/controllers"
	"ginchat/controllers/admin/base"
	"ginchat/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("index", controllers.Index)

	r.POST("/api/admin/base/open/login", base.Login)
	return r
}
