package main

import (
	"fmt"
	"go_study/c11_method/model"
	m2 "go_study/c11_method/model2" // 给包起别名
)

// 自定义函数、系统函数

func main(){
	v,err := model.Operation(1,2,'^')
	if err !=nil{
		panic(err)
	}
	fmt.Println(v)

	//
	v,err = m2.Operation(1,2,'^')
}

//