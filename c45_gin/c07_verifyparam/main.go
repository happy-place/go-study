package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"math/rand"
	"net/http"
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
		fmt.Println(n)
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
	r := gin.Default()

	r.POST("/user", userAction)

	//绑定
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("checkUidAvailable", checkUidAvailable)
	}

	r.Run()
}