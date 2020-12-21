package model

import (
	"fmt"
	"strings"
)

type FamilyAccount struct {
	space     string // "\t\t"
	header 	  string // "家庭收支管理系统\n\t【1】查看明细\n\t【2】存款\n\t【3】取款\n\t【4】退出\n请输入选项："
	history   string // := strings.ReplaceAll("收支space余额space用途\n","space",space)
	template  string // := strings.ReplaceAll("%s%.2fspace%.2fspace%s\n","space",space)
	initLen   int    // := len(history)
	balance   float64 // = 0.0
	toExit    bool    // = false
}

func NewAccount() *FamilyAccount {
	header := "家庭收支管理系统\n\t【1】查看明细\n\t【2】存款\n\t【3】取款\n\t【4】退出\n请输入选项："
	space := "\t\t"
	history := strings.ReplaceAll("收支space余额space用途\n", "space", space)
	template := strings.ReplaceAll("%s%.2fspace%.2fspace%s\n", "space", space)
	initLen := len(history)
	fc := &FamilyAccount{
		header: header,
		space:    space,
		history:  history,
		template: template,
		initLen:  initLen,
		balance:  0.0,
		toExit:   false,
	}
	return fc
}

func (fc *FamilyAccount) show() {
	if len(fc.history) == fc.initLen {
		fmt.Println("无收支记录，去存笔款吧！")
	} else {
		fmt.Println(fc.history)
	}
}

func (fc *FamilyAccount) deposit() {
	var money float64 = 0.0
	var note string = ""
	fmt.Printf("请输入存款金额：")
	fmt.Scanf("%v", &money)
	fmt.Printf("请输入收入来源：")
	fmt.Scanf("%s", &note)
	fc.balance += money
	fc.history += fmt.Sprintf(fc.template, "+", money, fc.balance, note)
}

func (fc *FamilyAccount) withdraw() {
	var money float64 = 0.0
	var note string = ""
	fmt.Printf("请输入取款金额：")
	fmt.Scanf("%v", &money)
	fmt.Printf("请输入取款用途：")
	fmt.Scanf("%s", &note)
	if money > fc.balance {
		fmt.Println("余额不足！")
	} else {
		fc.balance -= money
		fc.history += fmt.Sprintf(fc.template, "-", money, fc.balance, note)
	}
}

func (fc *FamilyAccount) exit() {
	var signal string
	fmt.Println("你确定要退出吗（y/n）？")
	fmt.Scanf("%s", &signal)
	fc.toExit = signal == "y"
}

func (fc *FamilyAccount) Meau() {
	var selection int
	for ;!fc.toExit;{
		fmt.Printf(fc.header)
		fmt.Scanf("%v",&selection)
		switch selection {
		case 1:
			fc.show()
		case 2:
			fc.deposit()
		case 3:
			fc.withdraw()
		case 4:
			fc.exit()
		default:
			fmt.Println("无效操作~")
		}
	}
}
