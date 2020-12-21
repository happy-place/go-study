package main

import (
	"errors"
	"fmt"
)

// a只能在函数体内被引用
func test1() {
	var a int = 1
	fmt.Println(a)
}

// 基本类型 - 传递值
func test2(a int) {
	a = 3
	fmt.Println(a)
}

func test3(a *int) {
	fmt.Println(*a) // 取地址 0
	*a = 3          // 地址赋值，然后设置值
}

// 报错 golang 不支持方法重载
//func test3(a ... int){
//
//}

// 形参列表中存在可变参数时，可变参数必须放在最后
func test4(a ...int) {
	// 可变参数
	fmt.Println(a) // [3]
}

// 多返回值
func test5(a int, b int) (int, int) {
	return a + b, a - b
}

// 抛出异常
func test6(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("被除数不能为0")
	}
	return a / b, nil
}

func test66(a int, b int) {
	res, err := test6(a, b)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

// golang 中运行向一个函数作为参数，传入另外一个函数
func any(arr []int, filter func(int) bool) bool {
	for e := range arr {
		if filter(e) {
			return true
		}
	}
	return false
}

// 先定义函数接口，然后传入
type filterFunc func(int) bool
func all(arr []int, filter filterFunc) bool {
	for e := range(arr){
		if !filter(e){
			return false
		}
	}
	return true
}

// golang 中运行对函数返回值进行命名，命名过程相当于在函数作用域内声明了sum,sub两个变量
func test7(a int,b int)(sum int,sub int){
	sum = a + b
	sub = a - b
	return
}

// 闭包，返回值时另外一个函数，但函数部分条件被限定了
func test8(a int) func(b int) int {
	return func(b int)int {
		return a + b
	}
}

// 形参列表参数类型相同，可以简写
func test9(n1,n2 float32) float32{
	fmt.Printf("n1 type: %T\n",n1)
	return n1 + n2
}

func main() {
	a := 0
	test2(a)
	fmt.Println(a) // 0

	test3(&a)      //
	fmt.Println(a) // 3

	test4(a)

	n1, n2 := test5(1, 2)
	fmt.Println(n1, n2)

	test66(1, 2)
	//test66(1,0) // panic: 被除数不能为0

	// golang 中函数也是一种数据类型
	fmt.Printf("%T\n", test66) // func(int, int)

	b := all([]int{-1,0,1},func(a int) bool{
		return a > 0
	})
	fmt.Println(b)

	// 允许只接收一个参数
	n1,_ = test7(1,2)
	fmt.Println(n1)

	// 调用闭包函数
	f2 := test8(10)
	n1 = f2(2)
	fmt.Println(n1)

	n3 := test9(1,2)
	fmt.Println(n3)


}
