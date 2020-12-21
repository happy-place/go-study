package main

import (
	"os"
	"strings"
	"fmt"
)

func main(){
	// 设置环境变量
	os.Setenv("FOO","1")
	// 取出环境变量
	fmt.Println("FOO:",os.Getenv("FOO")) // FOO: 1
	fmt.Println("BAR:",os.Getenv("BAR")) // BAR:

	fmt.Println()

	// 遍历环境
	for _,e := range os.Environ(){
		pair := strings.Split(e,"=") // 字符串拆分
		fmt.Println(pair[0],"->",pair[1])
	}
}
