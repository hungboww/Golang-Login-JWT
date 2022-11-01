package controllers

import "github.com/gin-gonic/gin"

func GetInforUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "Is Login",
	})
}
