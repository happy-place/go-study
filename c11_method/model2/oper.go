package model

import (
	"fmt"
	"errors"
)

func Operation(a float64 ,b float64,opt byte)(float64,error){
	var res float64
	switch opt{
	case '+':
		res = a + b
	case '-':
		res = a - b
	case '*':
		res = a * b
	case '/':
		res = a / b
	default:
		return -1,errors.New(fmt.Sprintf("不支持操作符:%c\n",opt))
	}
	//fmt.Printf("%v %c %v = %v\n",a,opt,b,res)
	return res,nil
}
