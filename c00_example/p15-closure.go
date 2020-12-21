package main

import "fmt"

// 闭包，返回值为另一个匿名函数，变量i被藏匿在函数中
func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main(){
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	nextInts := intSeq()
	fmt.Println(nextInts())
}

/**
1
2
3
1

 */


