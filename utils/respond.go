package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 请求成功
func Success[T any](c *gin.Context, data T) {
	c.JSON(http.StatusOK, gin.H{
		"message": "成功",
		"data":    data,
	})
}

// 服务器内部错误
func ServerError(c *gin.Context, data string) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": data,
	})
}
