package main

import (
	"fmt"
)
/**
defer 延迟执行
1) 先按正常代码顺序执行，碰到多个defer时，将defet语句压入defer;
2) 同一个方法正常语句执行完毕之后，再按栈顺序执行defer语句(先入后出)；
3) defer语句编译时，会直接拷贝一份引用的变量副本，基本数据类型直接拷贝值（不能感知后面的变化），引用类型拷贝地址（可以感知后面的变化）

值传递和引用传递
1）实质传递的都是变量副本，变量为基本数据类型时，传递的是值的拷贝，变量为引用类型时，传递的是引用的拷贝。
2）引用拷贝效率更高，值拷贝数据量比较大时，效率低；

值类型：基本数据类型int系列、float系列、byte系列（char）、complex系列、bool、string
引用类型：指针pointer、数组array、结构体struct、切片slice、哈希map、管道chan、interface

值传递：直接在栈上分配存储空间；
引用传递：分配栈上分配空间存储引用地址，堆上分配空间真正存储数据。



 */
func sum(a int,b int,c []int) int {
	defer fmt.Println("4: ",a)
	fmt.Printf("a=%v, b=%v\n",a,b)
	defer fmt.Println("3: ",c)
	a++
	b++
	c[0] = 10
	res := a + b
	fmt.Println("2: ",res)
	return res
}

func main(){
	res := sum(1,2,[]int{3,4,5})
	fmt.Println("5: ",res)
}

/**
a=1, b=2
2:  5
3:  2
4:  1
5:  5
 */

