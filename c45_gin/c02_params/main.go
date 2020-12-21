package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
/**
	请求参数
 */
func main() {
	r := gin.Default()

	// get 方式请求 http://localhost:8081/getUser/123
	r.GET("/getUser/:uid", func(c *gin.Context) {
		// http.StatusOK 枚举常量，等效于 200
		c.JSON(http.StatusOK, gin.H{
			"uid": c.Param("uid"),
		})
	})
	// 指定端口启动
	r.Run("0.0.0.0:8081")
}