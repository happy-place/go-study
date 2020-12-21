package main

import (
	"fmt"
	"strconv"
)

func main(){
	// 1.234 浮点
	f,_ := strconv.ParseFloat("1.234",64)
	fmt.Println(f)

	// 123 十进制，base=0 表示自动推断进制
	i,_ := strconv.ParseInt("123",0,64)
	fmt.Println(i)

	// 456，自动识别十六进制，然后转换为十进制整数
	d,_ := strconv.ParseInt("0x1c8",0,64)
	fmt.Println(d)

	// 789 10进制，无符整形
	u,_ := strconv.ParseUint("789",0,64)
	fmt.Println(u)

	// 135 10进制整形转换
	k,_ := strconv.Atoi("135")
	fmt.Println(k)

	// 解析错误报异常：strconv.Atoi: parsing "wat": invalid syntax
	_,e := strconv.Atoi("wat")
	fmt.Println(e)
}
