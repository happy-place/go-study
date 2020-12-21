package main

import "fmt"

// 多返回值函数
func vals() (int,int){
	return 3,7
}

func main(){
	a,b := vals()
	fmt.Println(a,b)

	// 只接受第二个返回值
	_,c :=vals()
	fmt.Println(c)
}

/**
3 7
7
 */