package main

import (
	"fmt"
)

// 常量，初始化时必须赋值，且一旦赋值，就不可变
const(
	A = iota // 从0 开始依次换行递增
	B
	C
	D,E = iota,iota // 写在同一行，值相同，并且都是从上面顺延下来
	F = iota
)


func test1(){
	fmt.Println(A,B,C,D,E,F)
	// 0 1 2 3 3 4
}


func main(){
	test1()
}