package main

import (
	"fmt"
	"flag"
)

// go build p60-flag.go --word=happy -numb=30 apple=34
func main(){
	// 解析输入参数，一下为默认值，返回指针
	wordPtr := flag.String("word","foo","a string")
	numbPtr := flag.Int("numb",42,"an int")
	boolPtr := flag.Bool("fork",false,"a bool")

	// 直接解析到指定变量引用中
	var svar string
	flag.StringVar(&svar,"svar","bar","a string var")

	// 开始解析
	flag.Parse()

	fmt.Println("word:",*wordPtr)
	fmt.Println("numb:",*numbPtr)
	fmt.Println("bool:",*boolPtr)
	fmt.Println("svar:",svar)
	fmt.Println("tail:",flag.Args()) // 剩余参数
}
