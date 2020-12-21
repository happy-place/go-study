package model

import (
	"fmt"
	"strings"
)

func FamilyAccontMenu(){
	space := "\t\t"
	history := strings.ReplaceAll("收支space余额space用途\n","space",space)
	template := strings.ReplaceAll("%s%.2fspace%.2fspace%s\n","space",space)
	initLen := len(history)
	var selection int
	var money float64 = 0.0
	var note string = ""
	var balance float64 = 0.0
	var toExit bool = false
	for ;!toExit;{
		fmt.Printf("家庭收支管理系统\n\t【1】查看明细\n\t【2】存款\n\t【3】取款\n\t【4】退出\n请输入选项：")
		fmt.Scanf("%v",&selection)
		switch selection {
		case 1:
			if len(history) == initLen{
				fmt.Println("无收支记录，去存笔款吧！")
			}else{
				fmt.Println(history)
			}
		case 2:
			fmt.Printf("请输入存款金额：")
			fmt.Scanf("%v",&money)
			fmt.Printf("请输入收入来源：")
			fmt.Scanf("%s",&note)
			balance += money
			history += fmt.Sprintf(template,"+",money,balance,note)
		case 3:
			fmt.Printf("请输入取款金额：")
			fmt.Scanf("%v",&money)
			fmt.Printf("请输入取款用途：")
			fmt.Scanf("%s",&note)
			if money > balance {
				fmt.Println("余额不足！")
			}else{
				balance -= money
				history += fmt.Sprintf(template,"-",money,balance,note)
			}
		case 4:
			var signal string
			fmt.Println("你确定要退出吗（y/n）？")
			fmt.Scanf("%s",&signal)
			toExit = signal == "y"
		default:
			fmt.Println("无效操作~")
		}
	}
}