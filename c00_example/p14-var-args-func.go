package main

import "fmt"

// 可变参数
func sum(nums ... int){
	fmt.Print(nums," ")
	total :=0
	for _,num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main(){
	sum(1,2)
	sum(1,2,3)
	nums := []int{1,2,3,4}
	sum(nums...) // 数组按可变参数入参
}

/**
[1 2] 3
[1 2 3] 6
[1 2 3 4] 10
 */