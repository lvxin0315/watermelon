package controller

import "github.com/gin-gonic/gin"

//基础控制器
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
