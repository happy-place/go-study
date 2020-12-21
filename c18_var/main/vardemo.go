package main

import "fmt"

var age = 10 // 脚本内部全局变量
var Name = "tom" // 整个包内全局变量

// 实质是 var family string; family = "beijing" 两条语句的缩写，第二条语句是赋值语句，不能出现在函数体外
//family := "beijing"

func test1(){
	age := 21 // 函数内部定义的具备变量
	fmt.Println("age=",age)
}

func test2(){
	age = 30 // 函数内部修改了全局变量
	fmt.Println("age=",age)
}

func retation(arr [][]int){
	fmt.Println(arr)
	for i:=0;i<len(arr);i++{
		for j:=0;j<i;j++{ // 以对角线为周反转元素
			arr[j][i],arr[i][j] = arr[i][j],arr[j][i]
		}
	}
	fmt.Println(arr)
}

func main(){
	fmt.Println("age=",age)
	test1()
	test2()
	fmt.Println("age=",age)

	retation([][]int{{1,2,3},{4,5,6},{7,8,9}})
}


/**
age= 10
age= 21
age= 30
age= 30
 */