package main

import "fmt"
import "unsafe"

func main(){
	var a int = 9
	fmt.Println(a)

	// 长度不够
	//var b int8 = 129
	//fmt.Println(b) // constant 129 overflows int8

	var b int8 = 127
	fmt.Println(b) //  constant 256 overflows uint8

	//var c uint8 = 256
	//fmt.Println(c)

	var c uint8 = 255
	fmt.Println(c)

	var d int = 8900
	fmt.Println(d)
	var e uint = 1

	var f byte = 255

	fmt.Println(e,f)

	var n1 = 100
	//n1类型 int,占用字节数:8
	fmt.Printf("n1类型 %T,占用字节数:%d\n",n1,unsafe.Sizeof(n1))

	// 给 int 取别名，myInt尽管能当int使用，但golang依旧认为myInt与int是不同类型
	type myInt int
	var m myInt = 40
	var num2 int
	//num2 = m // 不同类型不能直接赋值
	 num2 = int(m)

	fmt.Printf("m=%v,num2=%v, m type is %T\n",m,num2,m) // m=40, m type is newdemo.myInt



}

/**
数据类型
基本数据类型
	数值型
		整形：int(32为为4个字节，64位为8个字节) int8（-128~127） int16(-2^15~2^15-1) int32(-2^31~2^31-1) int64(-2^63~2^63-1)
			uint(32为为4个字节，64位为8个字节) uint8(0~2^8-1) uint16(0~2^16-1) uint32(0~2^32-1) uint64(0~2^64-1)
			byte(1个字节)
		浮点型 float32 float64，complex32 complex64
	字符型：没有专门字符型，使用byte保存单个字母字符
	布尔型：bool
	字符串：官方将字符型归到基本数据类型。
派生/复杂数据类型
	指针 pointer
	数组 array
	结构体 struct => class
	管道 chan
	函数 func
	切片 slice
	接口 interface
	集合 map


bit 计算机中最小存储单位，byte计算机中基本存储单元，1byte=8bit
*/