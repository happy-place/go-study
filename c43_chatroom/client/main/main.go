package main

import (
	"fmt"
	"go-study/c43_chatroom/client/process"
)

var userId int
var userPwd string

func menu(){
	var key string
	//var loop = true
	for;true;{
		fmt.Println("------------------ 欢迎登录多人聊天室系统 ------------------ ")
		fmt.Println("\t\t1 登录聊天室系统")
		fmt.Println("\t\t2 注册用户")
		fmt.Println("\t\t3 退出系统")
		fmt.Printf("\t\t请选择[1~3]：")
		fmt.Scanf("%s",&key)
		switch key {
		case "1":
			fmt.Println("登录~")
			fmt.Printf("请输入用户ID：")
			fmt.Scanf("%d",&userId)
			fmt.Printf("请输入密码：")
			fmt.Scanf("%s",&userPwd)
			up := &process.UserProcess{}
			err := up.Login(userId, userPwd)
			if err != nil {
				fmt.Println("登录失败")
			}else{
				fmt.Println("登录成功")
			}
			//loop = false
		case "2":
			fmt.Println("注册用户~")
			//loop = false
		case "3":
			fmt.Println("退出系统!")
			//loop = false
		default:
			fmt.Println("无效选择，请重新输入！")
		}
	}

}


func main(){
	menu()
}