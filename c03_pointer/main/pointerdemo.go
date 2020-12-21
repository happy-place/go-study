package main

import "fmt"

// 指针： 基本数据类型 int float bool string array struct
func main(){
	var i int = 20
	// i= 20 i的地址: 0xc00010c008
	fmt.Println("i=",i,"i的地址:",&i)

	var ptr *int = &i
	// ptr: 0xc00001e090, ptr value:20, ptr type: *int
	fmt.Printf("ptr: %v, ptr value:%v, ptr type: %T\n",ptr,*ptr,ptr)

	*ptr = 10 // i= 10 通过指针直接修改了原值
	fmt.Println("i=",i)
}
// 基本数据类型：int float bool string array struct
// 引用类型：指针 切片 map chan interface

/**
基本数据类型：int float bool string array struct
	变量直接存储值，内存通常在栈中分配
引用类型：指针 切片 map chan interface
	变量存储一个地址，此地址对应的内存空间才是真正存储数据区域，引用存储在栈，
	数据存储在堆，栈指向堆。
	当失去地址引用时，堆上的存储空间就变为垃圾。
 */
