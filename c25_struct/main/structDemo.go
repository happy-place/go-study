package main

import (
	"fmt"
	"encoding/json"
)

/**
	1)结构体时候一种自定义数据类型
	2)声明结构体时，数据空间已经存在了，未赋值时，取默认值，
		type 结构体名称 struct {
			变量名 类型 // 基本类型、引用类型、数组
			变量名 类型
			变量名 类型
		}
	3)slice map 指针默认值是nil
	4) 结构体相互赋值，是拷贝关系，而非引用关系


 */
type Cat struct {
	// 没有赋值时，赋默认值
	Name string // ""
	Age int // 0
	Color string // ""
}

type Person struct {
	Name string
	Age int
	score float64
	Ptr *int
	Slice []int
	Hobby map[string]string
}

type Monster struct {
	Name string
	Age int
}

func create() {
	// 方式1
	var c1 Cat
	fmt.Println(c1) // { 0 }
	c1.Age = 11
	c1.Name = "tom"
	c1.Color = "yellow"
	fmt.Println(c1) //{tom 11 yellow}

	// 方式2：
	c2 := Cat{
		"jerry",
		9,
		"red",
	}
	fmt.Println(c2)

	// 方式3
	c3 := Cat{
		Name:"jerry",
		Age:9,
		Color:"red",
	}
	// 显示key打印
	fmt.Printf("%+v",c3) // {Name:jerry Age:9 Color:red}

	// 方式4：
	var c4 *Cat = new(Cat)
	c4.Age = 12 // 实质是 (*c4).Age = 12，编译器自动进行转换
	(*c4).Name = "lily"
	(*c4).Color = "black"
	fmt.Println(c4.Color)
	fmt.Println(*c4) // c4 是地址

	// 方式5
	var c5 *Cat = &Cat{}
	//var c6 *Cat = &Cat{"scott",21,"white"}

	c5.Color = "white"
	(*c5).Age = 21
	c5.Name = "scott"
	fmt.Println(*c5)

}

func create2(){
	var p1 Person
	fmt.Println(p1) // { 0 0 <nil> [] map[]}

	if p1.Ptr == nil {
		fmt.Println("p1.score == nil")
	}

	if p1.Slice == nil {
		fmt.Println("p1.slice == nil")
	}

	if p1.Hobby == nil {
		fmt.Println("p1.Hobby == nil")
	}

	n := 10
	p1.Ptr = &n

	p1.Slice = make([]int,10)
	p1.Slice = append(p1.Slice,1,2)

	p1.Hobby = make(map[string]string)
	p1.Hobby["a"] = "1"
	p1.Hobby["b"] = "2"
	fmt.Println(p1) // { 0 0 0xc0000b4020 [0 0 0 0 0 0 0 0 0 0 1 2] map[a:1 b:2]}

}

func test1(){
	m1 := Monster{
		"tom",
		12,
	}
	m2 := m1  // 赋值，实质是将m1拷贝给m2，获得两个完全独立对象
	// %p 打印结构体地址
	fmt.Printf("m1: %+v, %p\n",m1,&m1) // m1: {Name:tom Age:12}, 0xc0000a6020
	fmt.Printf("m1: %+v, %p\n",m2,&m2) // m1: {Name:tom Age:12}, 0xc0000a6040

	m2.Age = 22
	fmt.Printf("m1: %+v, %p\n",m1,&m1) // m1: {Name:tom Age:12}, 0xc0000a6020
	fmt.Printf("m1: %+v, %p\n",m2,&m2) // m1: {Name:tom Age:22}, 0xc0000a6040
}


func test3(){
	c1 := Cat{"m1", 12, "white",}
	c2 := c1  // 值拷贝
	c3 := &c2 // 地址传递
	fmt.Printf("%v %p\n",c1,&c1) // {m1 12 white} 0xc000060180
	fmt.Printf("%v %p\n",c2,&c2) // {m1 12 white} 0xc0000601b0
	fmt.Printf("%v %p\n",*c3,c3) // &{m1 12 white} 0xc00000e028

	/**
	{m1 12 white} 0xc000060180
	{m1 12 white} 0xc0000601b0
	{m1 12 white} 0xc0000601b0
	*/

	c2.Color = "black"
	fmt.Printf("%v %p\n",c1,&c1)
	fmt.Printf("%v %p\n",c2,&c2)
	fmt.Printf("%v %p\n",*c3,c3)
	/**
	{m1 12 white} 0xc000060180
	{m1 12 black} 0xc0000601b0
	{m1 12 black} 0xc0000601b0
	 */
}

/**
	结构体使用的内存是连续的
 */

type Point struct {
	x int
	y int
}

type Rect struct {
	leftUp,rightDown Point
}

type Rect2 struct {
	leftUp,rightDown *Point
}

func test4(){
	// 8个字节
	r1 := Rect{Point{1,2}, Point{3,4},}
	fmt.Printf("r1.leftUp.x地址: %p\nr1.leftUp.y地址: %p\nr1.rightDown.x地址: %p\nr1.rightDown.y地址: %p\n",
		&r1.leftUp.x,&r1.leftUp.y,&r1.rightDown.x,&r1.rightDown.y)
	/**
	结构体中变量连续存储 int 默认是64位，刚好栈8个字节
	r1.leftUp.x地址: 0xc000016100
	r1.leftUp.y地址: 0xc000016108

	r1.rightDown.x地址: 0xc000016110
	r1.rightDown.y地址: 0xc000016118
	 */

	r2 := Rect2{&Point{10,30}, &Point{20,40},}
	fmt.Printf("r2.leftUp分配地址: %p, r2.rightDown分配地址: %p\n",&(r2.leftUp),&(r2.rightDown))
	fmt.Printf("r2.leftUp存储地址: %p\nr2.rightDown存储地址: %p\n", r2.leftUp,r2.rightDown)
	/**
		r2.leftUp分配地址: 0xc000010200, r2.rightDown分配地址: 0xc000010208
		r2.leftUp存储地址: 0xc00001e090
		r2.rightDown存储地址: 0xc00001e0a0
	 */
}

type Cat2 struct {
	// 给结构体加上标签，可以在序列化时，保证字段首字母小写
	Name string `json:"name"`
	Age int `json:"age"`
	Color string `json:"color"`
}

func test5(){
	c1 := Cat{"tom",12,"black"}
	var c2 Cat2
	//c2 = c1
	c2 = Cat2(c1) // 一个结构体变量可以强转为另一个结构体前提是，结构体变量名称、类型、顺序必须完全一致。
	fmt.Println(c2)
}

// 基于一个已经存在结构体重新定义新结构体，golang认为这是两个不同类型，不能直接赋值，但可以强转
type Cat3 Cat

func test6(){
	c1 := Cat{"tom",12,"black"}
	var c2 Cat3
	//c2 = c1
	c2 = Cat3(c1) // 一个结构体变量可以强转为另一个结构体前提是，结构体变量名称、类型、顺序必须完全一致。
	fmt.Println(c2)
}

func testMarshal(){
	c1 := Cat{"tom",12,"black"}
	jString, err := json.Marshal(c1) // jString 为[]byte 字节数组
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jString))

	c2 := Cat2{"tom",12,"black"}
	jString,_ = json.Marshal(c2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jString))
}

// 将结构体 与 方法绑定，此方法只能被对应绑定的结构体变量调用，且属于值传递
func (c Cat) printName(){
	c.Name = "black"
	fmt.Printf("obj: %+v, ptr: %p\n",c,&c)
	// obj: {Name:black Age:21 Color:white}, ptr: 0xc0000601b0
}

func testPrintName(){
	c := Cat{"tom",21,"white"}
	c.printName() // black
	fmt.Printf("obj: %+v, ptr: %p\n",c,&c)
	// obj: {Name:tom Age:21 Color:white}, ptr: 0xc000060180
}

// 拷贝的是地址
func (c *Cat) printName2(){
	c.Color = "black"
	fmt.Printf("obj: %+v, ptr: %p\n",*c,c)
}

func testPrintName2(){
	c := Cat{"tom",21,"white"}
	c.printName2() // black
	fmt.Printf("obj: %+v, ptr: %p\n",c,&c)
}

func test7(){
	c1 := Cat{"tom", 12, "white",} // {tom 12 white},ptr: 0xc000090180 结构体变量
	c2 := &Cat{"tom", 12, "white",} // {tom 12 white},ptr: 0xc0000901b0 结构体指针
	fmt.Printf("%v,ptr: %p\n",c1,&c1)
	fmt.Printf("%v,ptr: %p\n",*c2,c2)
}

func main(){
	//create()
	//create2()
	//test1()
	//test3()
	//test4()
	//test5()
	//testMarshal()
	//testPrintName()
	//testPrintName2()
	test7()


}