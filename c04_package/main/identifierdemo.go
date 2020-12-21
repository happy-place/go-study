package main

import (
	"fmt"
	"go_study/c04_package/model" // ~/.bash_porfile 中 GOPATH 定位到了 src
)

// Go中没有public 、private 概念，首字母大写代表public，首字母小写代表private
// 跨包引用时，直接 包名.变量名|函数名 即可
func main(){
	var a,b int = 1,2
	sum := model.GetSum(a,b)
	fmt.Println("sum=",sum)
	fmt.Println("Word:",model.Word)
}
