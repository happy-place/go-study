package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main(){
	fmt.Println(s)

	// 常量类型赋值一旦确认，就不能修改了
	const n = 500000000
	const d = 3e20/n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))

}


