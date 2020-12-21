package main

import "fmt"
import "unsafe"
func main(){
	var s1 = "武汉 900 abcde"
	fmt.Println(s1) // 武汉 900 abcde
	fmt.Printf("size of string: %d\n",unsafe.Sizeof(s1)) //size of string: 16

	// 字符串可重新赋值
	s1 = "abc"
	fmt.Println(s1)

	// 报错：但字符串元素，不能修改
	//s1[0] = '3'

	// 自动识别特殊转义字符，并予以高亮显示
	s2 := "abc\tdef"
	fmt.Println(s2) // abc     def

	// 多行字符串使用反引号
	s3 := `
package newdemo

import "fmt"
import "unsafe"
func newdemo(){
	var s1 = "武汉 900 abcde"
	fmt.Println(s1)
}`
	fmt.Println(s3)

	// 字符串拼接
	s4 := "hello " + "world"
	s4 += "wa!"
	fmt.Println(s4) // hello worldwa!

	// 长串拼接，换行时，最后需要保持为"+"
	s5 := "hello " + "hello " + "hello " + "hello " +
		"hello " + "hello " + "hello " + "hello " +
		"hello "
	fmt.Println(s5)

	var a int
	var b float32
	var c float64
	var d bool
	var e byte
	var f string
	fmt.Println("a=",a,"b=",b,"c=",c,"d=",d,"e=",e,"f=",f)
	// a= 0 b= 0 c= 0 d= false e= 0 f=
}

/**
基本数据类型默认值
int -> 0
float -> 0
byte -> 0
string -> ""
bool -> false
 */