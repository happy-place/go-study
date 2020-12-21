package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetHandler(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"msg":"get handler",
	})
}
