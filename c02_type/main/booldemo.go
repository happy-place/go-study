package main

import "fmt"
import "unsafe"

// 布尔类型：true false，占1个字节
// 适用于逻辑运算
func main(){
	var b = false
	fmt.Println("b=",b) // b= false
	fmt.Printf("size of bool: %d",unsafe.Sizeof(b)) // size of bool: 1
}
