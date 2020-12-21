package main

import "fmt"

// 递归调用
func test1(n int) int {
	if n==0{
		return 1
	}
	return test1(n-1)*n
}

func fabic(n int) int {
	if n==1 || n==2 {
		return 1
	}
	return fabic(n-1) + fabic(n-2)
}

func peach(day int) int{
	if day > 10 {
		fmt.Println("输入天数不能超过10")
		return 0
	}
	if day == 10 {
		return 1
	}
	return (peach(day+1)+1) * 2
}

func main(){
	n := 7
	res := test1(n)
	fmt.Printf("%v! = %v\n",n,res)

	fmt.Println(fabic(7))

	fmt.Println(peach(1)) // 1534
}



