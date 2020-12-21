package main

import "fmt"

// 单返回值
func getSum(a int,b int) int{
	return a + b
}
// 多返回值
func getSumAndSub(a int,b int)(int ,int) {
	return a + b, a - b
}

// 声明全局变量
var a = 100
var b = 200

// 一次性声明多个全局变量
var (
	a1 = 100
	b1 = 200
)

// 变量 = 变量名 + 值 + 数据类型：一旦声明，类型就不能改变
// 常量：一旦声明，值就不能变
func main(){
	// 指定类型初始化
	var a int = 1
	// 隐式推断类型
	var b = 2

	// 声明变量同时，初始化 等效于 var sum,sub int = getSumAndSub(a,b)
	sum,sub := getSumAndSub(a,b)
	fmt.Println("sum=",sum,",sub=",sub)
	sum,_ = getSumAndSub(a,b) // sum 已经声明过，直接使用 = 赋值
	fmt.Println("sum=",sum)
}


