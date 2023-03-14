package controllers

import (
	"github.com/gin-gonic/gin"
)

// Index
// @Tags 首页
// @Success 200 {string} helowindex
// @router /index [get]
func Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "helow index",
	})
}
