package main

import (
	"fmt"
	"os"
	"flag"
)

// 直接根据传参顺序解析
func defaultParse(){
	// var Args []string
	args := os.Args
	//inputs := args[1:]
	// 第0个参数：/private/var/folders/n1/mp08c98j4pjd62f8l1r_lkz00000gn/T/___go_build_ArgsDemo_go
	for i,arg := range args {
		fmt.Printf("第%d个参数：%v\n",i,arg)
	}
}

// 根据-标签解析
// go run xxx.go -u jack -p error -age 33
func flagParse(){
	var user string
	var pwd string
	var age int
	flag.StringVar(&user,"u","tom","用户名")
	flag.StringVar(&pwd,"p","hanks","密码")
	flag.IntVar(&age,"age",23,"年龄")
	flag.Parse()
	fmt.Printf("user: %s，pwd：%s，age：%d\n",user,pwd,age)
}

func main(){
	//defaultParse()
	flagParse()
}
