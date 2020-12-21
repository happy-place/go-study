package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HelloPage(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"msg":"hello",
	})
}
