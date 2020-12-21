package main

import (
	"fmt"
	"go_study/c15_init/model"
)

var name string = getName()

/**
 * 1) golang 默认每个脚本（包含不可执行脚本）包含一个init()函数，此函数在main()之前被执行
 * init 中可以执行一些初始化操作，自动编译执行；
 * 2) 文件中包含全局变量定义，init()和main()时，执行顺序是 全局变量定义 > init() > newdemo()
   3）不可执行脚本中的init() 对全局变量初始化时，引包过程就执行了init()

 */

func getName() string{
	fmt.Println("getName") // 1
	return "tom"
}

func init(){
	fmt.Printf("init %s\n",name) // 2
}

func main(){
	fmt.Println("hello go") // 3

	fmt.Println(model.Age)

}
/**
init
hello go
 */