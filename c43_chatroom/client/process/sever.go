package process

import (
	"fmt"
	"os"
)

// 显示登陆之后界面
func ShowMenu(){
	fmt.Println("------恭喜xxx登陆成功------")
	fmt.Println("------1.显示在线用户列表------")
	fmt.Println("------2.发送消息------")
	fmt.Println("------3.信息列表------")
	fmt.Println("------4.退出系统------")
	fmt.Printf("请选择[1~4]:")
	var key string
	fmt.Scanf("%s",&key)
	switch key {
	case "1":
		fmt.Println("显示在线用户列表")
	case "2":
		fmt.Println("发送消息")
	case "3":
		fmt.Println("信息列表")
	case "4":
		fmt.Println("你选择退出了系统...")
		os.Exit(0)
	default:
		fmt.Println("你输入选项不正确")
	}

}