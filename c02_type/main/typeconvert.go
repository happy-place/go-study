package main

import (
	"fmt"
	_ "unsafe" // 没被使用的包，又不想删除，前面使用'_'
	"strconv"
)

// go中不能进行自动类型转换，需要显示转换
func main(){
	var a int32 =22
	//var b float32 = a // 不同类型不能直接赋值
	var b =float32(a)
	var c = uint32(a)

	var d int64 = int64(a) // 低精度 -> 高精度
	fmt.Printf("a=%v, b=%v, c=%v, d=%v\n",a,b,c,d) // a= 22 b= 22 c= 22

	// 被强转对象本质没有改变
	fmt.Printf("a=%T\n",a) // a=int32

	// 高精度转低精度，不会报错，但结果可能不是想要的值
	var e int64 = 999999
	var f int8 = int8(e)
	fmt.Println("e=",e,"f=",f) // e= 999999 f= 63

	// 编译测试
	// 组1: 编译不通过
	//var n1 int32 = 12
	//var n2 int64
	//var n3 int8
	//n2 = n1 + 20  // <<< int64 = int32 + int32
	//n3 = n1 + 20  // <<< int8 = int32 + int32

	// 组2：
	//var n1 int32 = 12
	//var n3 int8
	//var n4 int8
	//n4 = int8(n1) + 127 // 此处 127 被当成 int8，n4空间不够，溢出了
	//n3 = int8(n1) + 128 // 此处 128 按 int32 处理，编译不通过

	// 基本数据类型与string直接转换
	// 方案1：fmt.Sprintf("%v")

	var num1 int = 99
	var num2 float64 = 23.456
	var num3 bool = true
	var myChar byte = 'h'
	var str string
	str = fmt.Sprintf("%d",num1)
	fmt.Printf("str=%q,str type: %T\n",str,str) // num1="99",num1 type: string

	str = fmt.Sprintf("%f",num2)
	fmt.Printf("str=%q,str type: %T\n",str,str) // num1="23.456000",num1 type: string

	str = fmt.Sprintf("%t",num3)
	fmt.Printf("str=%q,str type: %T\n",str,str) // num1="true",num1 type: string

	str = fmt.Sprintf("%c",myChar)
	fmt.Printf("str=%q,str type: %T\n",str,str) // str="h",str type: string

	// 方式2： strconv
	str = strconv.FormatInt(99,10)
	fmt.Printf("str=%q,str type: %T\n",str,str) // str="99",str type: string

	// 'f'（-ddd.dddd)，10: 小数位保留10位，64：float64
	str = strconv.FormatFloat(num2,'f',10,64)
	fmt.Printf("str=%q,str type: %T\n",str,str) // str="23.4560000000",str type: string
	/**
	bitSize表示f的来源类型（32：float32、64：float64），会据此进行舍入。
	fmt表示格式：'f'（-ddd.dddd）、'b'（-ddddp±ddd，指数为二进制）、'e'（-d.dddde±dd，十进制指数）、'E'（-d.ddddE±dd，十进制指数）、'g'（指数很大时用'e'格式，否则'f'格式）、'G'（指数很大时用'E'格式，否则'f'格式）。
	prec控制精度（排除指数部分）：对'f'、'e'、'E'，它表示小数点后的数字个数；对'g'、'G'，它控制总的数字个数。如果prec 为-1，则代表使用最少数量的、但又必需的数字来表示f。
	 */

	str = strconv.FormatBool(num3)
	fmt.Printf("str=%q,str type: %T\n",str,str) // str="true",str type: string

}
