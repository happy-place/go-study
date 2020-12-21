package main

import (
	"fmt"
	"strings"
)

var f2 func(int,int)int

// 闭包，一个函数与其执行环境组成一个整体(避免反复传环境、同时可以实现环境递增效果)
func clouser(a int) func (int) int {
	// 以下部分才是闭包，返回的是匿名函数，但函数执行环境依赖于外部传参
	return func(b int) int {
		return a+b
	}
}

func addUpper(a int) func(x int) {
	var str string
	return func(x int){
		a += x
		str += "$"
		fmt.Printf("a=%v, str=%s\n",a,str)
	}
}

func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if strings.HasSuffix(name,suffix){
			// 扩展名一致 直接返回
			return name
		}else if strings.Contains(name,"."){
			// 扩展名不同，予以替换
			index := strings.LastIndex(name,".")
			temp := strings.Split(name,".")
			oldSuffix := "." + temp[len(temp)-1]
			return strings.Replace(name,oldSuffix,suffix,index)
		}else{
			// 无扩展名，直接添加
			return fmt.Sprintf("%s%s",name,suffix)
		}
	}
}



func main(){
	// 匿名函数直接调用，且只能调用一次
	res := func (a int,b int) int {
		return a + b
	}(1,2)
	fmt.Println("res=",res)

	// 匿名函数赋值给变量可以重复使用
	f1 := func (a int,b int) int {
		return a+b
	}
	res = f1(1,2)
	fmt.Println("res=",res)

	// 匿名函数赋值给全局变量，就相当于一个全局定义函数
	f2 = func (a int,b int) int {
		return a+b
	}

	f3 := clouser(10)
	res = f3(2)
	fmt.Println("res=",res)

	f4 := addUpper(1)
	f4(1)
	f4(2)
	f4(3)
	/**
	a=2, str=$
	a=4, str=$$
	a=7, str=$$$
	*/

	f5 := makeSuffix(".jpg")
	fmt.Println(f5("a.svg"))
	fmt.Println(f5("a.jpg"))
	fmt.Println(f5("a"))
}
