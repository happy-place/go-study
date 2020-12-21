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
	form 参数绑定到 struct
	curl -XPOST http://127.0.0.1:8080/user -d 'name=xiaoming&date=2019-01-01'
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
	// 方案1：r := gin.Default() 默认都输出到日志文件
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	gin.DefaultErrorWriter = io.MultiWriter(f)

	r := gin.Default() // 默认 日志 和 panic 异常输出控制台, 通过上述设置 正常访问记录日志，异常直接返回页面

	r.POST("/user", userAction)

	//绑定
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("checkUidAvailable", checkUidAvailable)
	}

	r.Run()


}