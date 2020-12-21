package main

import "fmt"

func main(){
	// 变量类型一旦确认，就不能修改
	// 声明变量 并 初始化赋值，go会自动推断类型
	var a = "initial"
	fmt.Println(a)

	var b,c int = 1,2
	fmt.Println(b,c)

	var d = true
	fmt.Println(d)

	// 只声明变量，未赋值时，按默认值处理
	var e int
	fmt.Println(e)

	// var f string = "short" 的简写，声明变量及类型，同时进行初始化
	f := "short"
	fmt.Println(f)
}
