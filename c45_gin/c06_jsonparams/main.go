package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

/**
	form 参数绑定到 struct
	curl -XPOST http://127.0.0.1:8080/user -d 'name=xiaoming&date=2019-01-01'
 */
type requestUser struct {
	Name string `form:"name"`
	Uid  int    `form:"uid"`
	Age  int    `form:"age"`
	Date time.Time `form:"date" time_format:"2006-01-02"` // 2006-01-02T15:04:05Z07:00
}
// 不设置解析格式，或报错：parsing time "2019-01-01" as "2006-01-02T15:04:05Z07:00": cannot parse "" as "T"{xiaoming 0 0 0001-01-01 00:00:00 +0000 UTC}

func main() {
	r := gin.Default()
	r.POST("/user", userAction)
	r.Run()
}

func userAction(c *gin.Context) {
	var requestUser requestUser

	//这里是根据请求的content-type进行解析(可以解析json)
	if err := c.ShouldBind(&requestUser); err != nil {
		c.String(http.StatusOK, err.Error())
		c.Abort()
	}

	c.String(http.StatusOK, "%v", requestUser)
}