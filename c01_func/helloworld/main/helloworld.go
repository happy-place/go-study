package main

import "fmt"

func main(){
	fmt.Println("hello Golaod")
}

/**
1.安装SDK

2.配置环境变量
export GOROOT=/usr/local/go
export GOBIN=$GOROOT/bin
export PATH=$PATH:$GOBIN
// 当前项目路径
export GOPATH=$HOME/go_proj

3.测试
go version

4.源码包 $GOBIN

5.编写、编译、运行
编写：写.go文件
编译: go build xx.go 生成二进制可执行文件，会自动将所有依赖汇总到一个文件（可执行文件比源码大很多）
运行：
	方案1：./xx 直接操作编译好的可执行文件
	方案2：go run xx.go 操作源码，运行

6.格式化
	gofmt xx.go 格式化打印出来
	gofmt -w xx.go 格式化然后写回源文件

7.语法要求
1）源码必须是.go后缀结尾文件；
2）必须以包package为单位组织文件；
3）import 对象为包名 import "fmt" 或 包内脚本 import "math/rand"，且凡是import对象必须被使用，否则报错；
4）newdemo() 方法时程序入口，凡是创建变量必须使用；
5）不使用分号作为结束符；
6）多条语句禁止写在同一行
7）循环体、分支结构一般不使用（），但语句块仍然使用{},且左大括号，不能单独起行
 func newdemo(){

 }
 */



