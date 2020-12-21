package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"io"
	"math/rand"
	"net/http"
	"os"
)

/**
	日志 与 panic 分开处理
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

func main() {
	// 日志和异常
	// 方案1：r := gin.Default() 默认都输出到控制台
	// 方案2：r.Use(gin.Logger(), gin.Recovery()) 正常调用记录到日志文件，异常抛出到上层
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	gin.DefaultErrorWriter = io.MultiWriter(f)

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery()) // 正常日志计入gin.log,异常，直接抛出到控制台

	r.POST("/user", userAction)

	//绑定
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("checkUidAvailable", checkUidAvailable)
	}

	r.Run()


}