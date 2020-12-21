package main

import (
	"fmt"
	"os"
)

type point struct {
	x,y int
}

func main(){
	p := point{1,2}
	fmt.Printf("%v\n",p) // {1 2}
	fmt.Printf("%+v\n",p) // {x:1 y:2}
	fmt.Printf("%#v\n",p) // newdemo.point{x:1, y:2}
	fmt.Printf("%T\n",p) // newdemo.point
	fmt.Printf("%t\n",true) // true
	fmt.Printf("%d\n",123) // 123
	fmt.Printf("%b\n",14) // 二进制 1110
	fmt.Printf("%c\n",33) // ! （整数对应unicode编码）
	fmt.Printf("%x\n",456) // 1c8 （16进制）
	fmt.Printf("%f\n",78.9) // 78.900000 （浮点，，默认保持6为小数）
	fmt.Printf("%e\n",123400000.0) // 1.234000e+08 科学计数法
	fmt.Printf("%E\n",123400000.0) // 1.234000E+08
	fmt.Printf("%s\n","a \"string\"") // a "string" 代引号输出
	fmt.Printf("%q\n","string") // 整体引起来
	fmt.Printf("%x\n","hex this") // 6865782074686973（base-16 编码）
	fmt.Printf("%p\n",&p) // 0xc00001e090 (指针)
	fmt.Printf("|%6d|%6d|\n",12,345) // |    12|   345| （整数6个字符长度右对齐）
	fmt.Printf("|%6.2f|%6.2f|\n",1.2,3.45) // |  1.20|  3.45| (浮点6位长度，右对齐，小数点2位)
	fmt.Printf("|%-6.2f|%-6.2f|\n",1.2,3.45) // |1.20  |3.45  | (浮点6位长度，左对齐)
	fmt.Printf("|%6s|%6s|\n","foo","b") // |   foo|     b| （浮点6位长度右对齐）
	fmt.Printf("|%-6s|%-6s|\n","foo","b") // |foo   |b     | (浮点6位长度，左对齐)
	s := fmt.Sprintf("a %s","string") // 格式化组合字符串
	fmt.Println(s)

	// 标准错误流中输出
	fmt.Fprintf(os.Stderr,"an %s\n","error")
}
