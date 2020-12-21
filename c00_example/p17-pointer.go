package main

import "fmt"

// 普通函数，传递值
func zeroval(ival int){
	ival = 0
}

// 传递指针，可以直接修改值
func zeroprt(iprt *int){
	*iprt = 0 // 指定地址 赋值
}

func main(){
	i := 1
	fmt.Println("initial:",i)

	zeroval(i)
	fmt.Println("zeroval:",i)

	zeroprt(&i) // 取地址
	fmt.Println("zeroprt:",i)

	fmt.Println("pointer:",&i)
}

/**
initial: 1
zeroval: 1
zeroprt: 0
pointer: 0xc0000b4008
 */