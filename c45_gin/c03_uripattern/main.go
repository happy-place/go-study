package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
/**
	uri 前缀匹配
 */
func main() {
	r := gin.Default()

	r.GET("/user/*a", func(c *gin.Context) {
		c.String(http.StatusOK, "hello user")
	})

	// 指定端口启动
	r.Run("0.0.0.0:8081")
}