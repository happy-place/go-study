package main

import "fmt"

// 递归： 收敛条件+对自调用
func fact(n int) int {
	if n==0{
		return 1
	}
	return n * fact(n-1)
}

func main(){
	fmt.Println(fact(7))
}
// 5040
