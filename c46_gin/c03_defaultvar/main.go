package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/user",GetHandler)
	r.Run(":8080")
}

func GetHandler(c *gin.Context) {
	name := c.DefaultQuery("name", "Guest") // 参数存在默认值
	pass := c.Query("pass")
	id, existed := c.GetQuery("id") // 通过error判断是否存在
	if !existed {
		name = "the key 'id' not existed"
	}
	c.Data(http.StatusOK,"text/plain",[]byte(fmt.Sprintf("get success! %s %s %s",name,pass,id)))
}
