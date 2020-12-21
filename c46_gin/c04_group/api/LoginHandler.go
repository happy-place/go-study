package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginHandler(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"msg":"login handler",
	})
}

