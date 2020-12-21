package main

import "fmt"

// go 没有字符类型，使用字节 byte 表示
// go 字符串，实质就是字节数组
// go 使用的是utf-8编码，utf-8编码包含ascii，因此使用时没有乱码困扰
// go 中字符本质上是一个整数，数值就是对应unicode编码，因此可直接进行算数运算
func main(){
	var c1 byte = 'c'
	var c2 byte = '0'

	// 直接输出字符时，实质输出的是对应字符的ascii码值
	fmt.Println("c1=",c1,"c2=",c2) // c1= 99 c2= 48
	// 输出字符
	fmt.Printf("c1=%c,c2=%c\n",c1,c2)

	// byte 保存汉字溢出了
	//var c3 byte = '北'
	//fmt.Printf("c3=%c\n",c3) // constant 21271 overflows byte

	// 换个更大类型 就可以保存了
	var c4 int = '北'
	fmt.Printf("c4=%c\n",c4)

	// go 使用的是utf-8编码，utf-8编码包含ascii
	// go 中字符本质上是一个整数，数值就是对应unicode编码，因此可直接进行算数运算

	var c5 = 22269
	fmt.Printf("c5=%c\n",c5) // c5=国

	var c6 = '中' + 40
	fmt.Printf("c6=%c\n",c6) // c6=乕

	// 存储: 字符 -> 对应码值 -> 二进制 -> 存储
	// 读取: 二进制 -> 码值 -> 字符 -> 读取


	 var a byte
	fmt.Println("a=",a) // a= 0

}
