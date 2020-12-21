package model

import "fmt"

type Account struct {
	AccountNo string
	Pwd string
	Balance float64
}

func (acc *Account) Deposite(money float64,pwd string){
	if pwd!= acc.Pwd {
		fmt.Println("输入密码不正确")
		return
	}
	if money <= 0 {
		fmt.Println("输入金额不正确")
		return
	}
	acc.Balance += money
	fmt.Println("存库成功")
}

func (acc *Account)WithDraw(money float64,pwd string){
	if pwd !=acc.Pwd{
		fmt.Println("输入密码不正确")
		return
	}
	if money<=0 {
		fmt.Println("输入金额不正确")
		return
	}
	if money>acc.Balance {
		fmt.Println("余额不足")
		return
	}
	acc.Balance -= money
}

func (acc *Account) Query(pwd string){
	if pwd !=acc.Pwd{
		fmt.Println("输入密码不正确")
		return
	}
	fmt.Printf("当前余额为：%.2f\n",acc.Balance)
}
