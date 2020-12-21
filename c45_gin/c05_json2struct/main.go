package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)
/**
	json 参数
	curl -XPOST http://127.0.0.1:8080/user?name=1 -d {"name":1}
 */
func main() {
	r := gin.Default()

	r.POST("/user", func(c *gin.Context) {
		bodyBytes, error := ioutil.ReadAll(c.Request.Body)
		if error != nil {
			c.String(http.StatusOK, error.Error())
			c.Abort()
		}
		c.String(http.StatusOK, string(bodyBytes))
	})

	// 指定端口启动
	r.Run("0.0.0.0:8081")
}