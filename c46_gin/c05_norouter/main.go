package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.NoRoute(NoResponse)
	r.GET("/user",GetHandler)
	r.Run(":8080")
}

func GetHandler(c *gin.Context) {
	c.Get("get handler")
}

func NoResponse(c *gin.Context) {
	c.JSON(http.StatusOK,gin.H{
		"status":404,
		"error":"404, page not exists!",
	})
}
