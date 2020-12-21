package main

import (
	"github.com/gin-gonic/gin"
	"go-study/c46_gin/c04_group/api"
)

func main(){
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// 基于组进行访问
	adminGroup := r.Group("admin")
	{
		adminGroup.GET("ping",api.HelloPage)
		adminGroup.GET("welcome",api.LoginHandler)
		adminGroup.GET("getquery",api.GetHandler)
	}

	r.Run(":8080")
}
