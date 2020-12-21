package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// 关闭彩色控制台
	gin.DisableConsoleColor()
	//gin.SetMode(gin.DebugMode)
	//gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.TestMode)

	router := gin.Default()


	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "get",
		})
	})
	router.POST("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "post",
		})
	})
	router.PUT("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "put",
		})
	})
	router.DELETE("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "delete",
		})
	})
	router.PATCH("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "patch",
		})
	})
	router.HEAD("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "head",
		})
	})
	router.OPTIONS("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "options",
		})
	})

	group := router.Group("/group") // 默认为 /
	group.GET("test",func(c *gin.Context) {
		c.JSON(200, gin.H{ // H 是hash简写
			"message": "group get",
		})
	})

	router.Run(":8080")
}
