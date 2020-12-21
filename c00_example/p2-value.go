package main

import "fmt"

func main(){
	fmt.Println("go"+"lang")
	fmt.Println("1+1=",1+1)
	fmt.Println("7.0/3.0=",7.0/3) // 分子、分母任意取浮点就能实现标准除法，否则按按取整处理
	fmt.Println("7/3.0=",7/3.0) // 分子、分母任意取浮点就能实现标准除法，否则按按取整处理
	fmt.Println("7.0/3=",7.0/3) // 分子、分母任意取浮点就能实现标准除法，否则按按取整处理
	fmt.Println("7/3=",7/3) // 分子、分母任意取浮点就能实现标准除法，否则按按取整处理
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
