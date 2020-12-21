package model

import (
	"fmt"
)

type account struct {
	accNo   string
	balance float64
	pwd     string
}

func NewAcc(accNo string,balance float64,pwd string) *account {
	if len(accNo) < 6 || len(accNo) > 10 {
		fmt.Println("账号长度需保持在6~10位之间")
		return nil
	}
	if len(pwd) !=6 {
		fmt.Println("密码长度需保持为6位")
		return nil
	}
	if balance <= 20 {
		fmt.Println("余额需大于20")
		return nil
	}
	return &account{accNo,balance,pwd}
}

func (acc *account) SetAccNo(accNo string) {
	length := len(accNo)
	if length >= 6 && length <= 10 {
		acc.accNo = accNo
	} else {
		fmt.Println("账号长度需保持在6~10位之间")
	}
}

func (acc *account) GetAccNo() string {
	return acc.accNo
}

func (acc *account) SetBalance(balance float64) {
	if balance > 20 {
		acc.balance = balance
	} else {
		fmt.Println("余额需大于20")
	}
}

func (acc *account) GetBalance() float64 {
	return acc.balance
}

func (acc *account) SetPwd(pwd string) {
	length := len(pwd)
	if length == 6 {
		acc.pwd = pwd
	}else{
		fmt.Println("密码长度需保持为6位")
	}
}

func (acc *account)GetPwd()string{
	return acc.pwd
}
