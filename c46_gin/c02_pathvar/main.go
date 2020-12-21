package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HinHandler(c *gin.Context){
	name := c.Param("name")
	pass := c.Param("pass")
	c.String(http.StatusOK,"你好%s，你的密码为%s",name,pass)
}

func main(){
	r := gin.Default()
	r.GET("/user/:name/:pass",HinHandler)
	r.Run(":8080")
}
