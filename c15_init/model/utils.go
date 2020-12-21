package model

import "fmt"

var Age int
var Name string

func init(){
	// 引入过程初始化
	fmt.Println("utils.go init")
	Age = 10
	Name = "tom"
}
