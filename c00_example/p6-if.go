package main

import "fmt"

func main(){
	// if-else if- else 分支，go没有三目运算
	// 判断条件带 () 也是支持的
	if (7 % 2 ==0) {
		fmt.Println("7 is even")
	}else{
		fmt.Println("7 is odd")
	}

	if 8 % 4 ==0 {
		fmt.Println("8 is divisible by 4")
	}

	// if 语句中可以写赋值语句
	if num := 9; num <0 {
		fmt.Println(num," is negative")
	}else if num < 10 {
		fmt.Println(num," has 1 digit")
	}else {
		fmt.Println(num," has multiple digits")
	}


}


