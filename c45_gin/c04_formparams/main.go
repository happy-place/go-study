package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
/**
	默认参数
 */
func main() {
	r := gin.Default()

	// "age": c.Param("age"), /:age 路径传参
	// "name": c.Query("name"), ?k=v 普通键值对传参
	r.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name": c.Query("name"),
			"sex":  c.DefaultQuery("sex", "女"), // 要么传入，要么使用默认的
		})
	})

	// Content-Type: application/x-www-form-urlencoded 表单参数
	r.POST("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name": c.PostForm("name"),
			"sex":  c.DefaultPostForm("sex", "女"), // 要么传入，要么使用默认的
		})
	})

	// 指定端口启动
	r.Run("0.0.0.0:8081")
}