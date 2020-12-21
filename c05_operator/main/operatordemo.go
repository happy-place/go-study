package main

import (
	"fmt"
	"go_study/c05_operator/model"
)

func main(){
	// 整除
	fmt.Println(10/4) // 2

	var n1 float32 = 10 / 4 // 2
	fmt.Println(n1)

	// 标准除：分子、分母任意一个为浮点，即可
	var n2 float64 = 10.0 /4 // 2.5
	fmt.Println(n2)

	var n3 = 10 / 4.0
	fmt.Println(n3) // 2.5

	// 取余，符号与被除数一致
	fmt.Println(10 % 3) // 1
	fmt.Println(-10 % 3) // -1
	fmt.Println(10 % -3) // 1
	fmt.Println(-10 % -3) // -1

	// 自运算，只有 i++、i--，没有 --i、++i，且不能赋值给其他变量
	var i = 1
	i++
	fmt.Println(i)
	i--
	fmt.Println(i)

	// 运算符练习
	fmt.Println(model.GetTemperature(120.0))
	fmt.Println(model.GetWeeksAndDays(72))

	//
	var aa = 1
	var bb = 2
	fmt.Println(aa==bb)
	fmt.Println(aa!=bb)
	fmt.Println(aa>bb)
	fmt.Println(aa>=bb)
	fmt.Println(aa<bb)
	fmt.Println(aa<=bb)

	flag := aa > bb
	fmt.Printf("flag=%v, flag type:%T\n",flag,flag) // flag=false, flag type:bool

	// 逻辑运算, 短路与 短路或 非
	fmt.Println(true && false) // false
	fmt.Println(true || false) // true
	fmt.Println(! false) // true

	// 计算后赋值 -= *= /= %=
	a := 1
	a += 20
	fmt.Println(a)

	// 位运算 <<= >>= &= |= ^= |=
	a = 1
	a <<= 1 // 按位右移
	fmt.Println(a) // 2

	// 2 1
	fmt.Println(model.Swap(1,2))

	// & 取地址 * 取指针
	prt := &a
	fmt.Printf("prt: %v, prt value: %v, prt type: %T\n",prt,*prt,prt)
	// prt: 0xc00001e0b8, prt value: 2, prt type: *int

	fmt.Println("Max:",model.Max(1,2,3))
}

// go 设计理念是 一件事情有且仅有一种实现方式，因此没有三元运算
