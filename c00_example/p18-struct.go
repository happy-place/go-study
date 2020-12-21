package main

import "fmt"

type person struct {
	name string
	age int
}

func main(){
	// 直接初始化
	fmt.Println(person{"Bob",20})
	fmt.Println(person{name:"Alice",age:30})
	// 没有赋值的按默认值处理
	fmt.Println(person{name:"Fred"})
	// 取结构体地址
	fmt.Println(&person{name:"Ann",age:40})
	s := person{name:"Sean",age:50}
	fmt.Println(s.name) // 对象.属性 访问
	sp := &s // 取结构体地址，本质上还是对象
	fmt.Println(sp.age) // 通过结构体指针访问属性

	sp.age = 51
	fmt.Println(sp.age)
}

/**
{Bob 20}
{Alice 30}
{Fred 0}
&{Ann 40}
Sean
50
51
 */