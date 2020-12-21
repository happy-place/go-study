package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"io"
	"math/rand"
	"net/http"
	"os"
)

/**
	授信IP访问中间件
 */

// 更多验证规则参考 https://godoc.org/gopkg.in/go-playground/validator.v8
type requestUser struct {
	Name string `form:"name" binding:"required,len=10"`
	Uid  int    `form:"uid" binding:"checkUidAvailable"`
	Age  int    `form:"age" binding:"required,gt=10"`
}

// Structure
func checkUidAvailable(field validator.FieldLevel) bool {
	if fieldValue, ok := field.Field().Interface().(int); ok {
		//随机概率
		n := rand.Intn(2)
		if n == fieldValue {
			return true
		}
	}

	return false
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

func ipAuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		ipList := []string{ // 授信ip
			"127.0.0.1",
		}

		flag := false
		clientIp := c.ClientIP()
		fmt.Println(clientIp)
		for _, host := range ipList {
			if clientIp == host {
				flag = true
				break
			}
		}

		if !flag {
			c.String(http.StatusOK, "error")
			c.Abort()
		}
	}
}

func main() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	gin.DefaultErrorWriter = io.MultiWriter(f)

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery(), ipAuthMiddleWare()) // 日志 异常恢复 授信ip过滤

	r.POST("/user", userAction)

	//绑定
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("checkUidAvailable", checkUidAvailable)
	}

	r.Run()


}