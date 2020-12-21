package main

import (
	"fmt"
	"os"
)

func main(){
	// 下标为0位置是脚本名称
	argsWithProg := os.Args
	// 从下标为1开始就是全部参数
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]

	fmt.Println(argsWithProg) // ['p59-args.go',a b c d e]
	fmt.Println(argsWithoutProg) // [a b c d e]
	fmt.Println(arg) // c
}
