package main

import "fmt"

func main(){

	// 声明数组：int 默认值为0
	var a [5]int
	fmt.Println("emp:",a)

	// 修改元素
	a[4] = 100
	fmt.Println("set:",a)
	fmt.Println("get:",a[4])

	// 数组长度
	fmt.Println("len:",len(a))

	b := [5]int{1,2,3,4,5}
	fmt.Println("dcl:",b)

	// 二维数组
	var twoD [2][3]int
	for i:=0;i<2;i++{
		for j:=0;j<3;j++{
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d:",twoD)
}
/**
emp: [0 0 0 0 0]
set: [0 0 0 0 100]
get: 100
len: 5
dcl: [1 2 3 4 5]
2d: [[0 1 2] [1 2 3]]
 */
