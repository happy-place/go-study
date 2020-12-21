package main

import (
	"fmt"
	"go_study/c07_binary/model"
)

func main(){
	// 十进制
	var n1 int = 19
	fmt.Println(n1)

	// 二进制 1001172 不能直接表示，但可以打印输出
	fmt.Printf("%b",n1)

	// 八进制
	var n2 int = 0110
	fmt.Println(n2) // 72

	// 以 0x 或0X开头的称为十六进制
	var n3 int = 0x110
	fmt.Println(n3) // 272

	var n4 int = 0X110
	fmt.Println(n4) // 272

	// 计算机使用补码参与运算
	// 一个字节有8位二进制，首位为符号位，正数首位为0，负数为1
	// 正数的 原码 反码 补码 都相同
	// 负数原码：首位(符号位)为1，其余是多少就算多少
	// 负数反码：原码首位不变(符号位)，其余位取反
	// 负数补码：反码 + 1

	var a,b = 2,3
	model.PrintBin(a,b)

	var c = a & b
	fmt.Printf(" %b\n%c%b\n=%b=%v\n",a,'&',b,c,c)

	c = a | b
	fmt.Printf(" %b\n%c%b\n=%b=%v\n",a,'|',b,c,c)

	c = a ^ b
	fmt.Printf(" %b\n%c%b\n=%b=%v\n",a,'|',b,c,c)

	a = -a
	c = a ^ b
	fmt.Printf(" %b\n%c%b\n=%b=%v\n",a,'|',b,c,c)



}
