package main

import "fmt"
import "time"

func f(from string){
	for i:=0;i<3;i++{
		fmt.Println(from,":",i)
		time.Sleep(1000)
	}
}

func main(){
	// 直接调用普通函数
	f("direct")

	// 在协程中异步调用
	go f("goroutine")

	// 声明匿名函数同时，在协程中发起异步调用
	go func(msg string){
		fmt.Println(msg)
	}("going")

	// 定义字符串变量接收命令行传入数据
	var input string
	fmt.Scanln(&input)
	fmt.Println("input:",input)
	fmt.Println("done")
}