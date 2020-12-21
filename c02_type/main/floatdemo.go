package main

import "fmt"

// go 浮点有固定取值范围和长度，不受OS影响
func main(){
	// 默认浮点为 floet64
	var n1 = 1.23
	fmt.Printf("n1 is %T\n",n1) // n1=float64
	var n2 = .23 // 浮点必须有小数点
	fmt.Println("n2=",n2) // n2= 0.23

	// 浮点 = 符号位 + 整数位
	var n3 float32 = 1.234678009 // 009 丢失
	var n4 float64 = 1.234678009 // 默认是 float64
	fmt.Println("n3=",n3,"n4=",n4) // n3= 1.234678 n4= 1.234678009

	var n5 = 2.3456e3
	var n6 = 2.3456E10
	var n7 = 2.3456E-2
	fmt.Println("n5=",n5,"n6=",n6,"n7=",n7) // n5= 2345.6 n6= 2.3456e+10 n7= 0.023456

}
