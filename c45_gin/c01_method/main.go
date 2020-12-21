package main

import (
	"github.com/gin-gonic/gin"
)

/**
	请求方法
 */
func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "get",
		})
	})

	r.PUT("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "put",
		})
	})

	r.POST("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "post",
		})
	})

	r.DELETE("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "delete",
		})
	})

	r.Any("/any", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "any",
		})
	})

	r.Handle("POST","/handler",func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "post handler",
		})
	})

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run()
}