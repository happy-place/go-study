package main

import (
	"fmt"
	"sort"
)

func main(){
	str := []string{"c","a","b"}
	sort.Strings(str) // 对string进行排序
	fmt.Println(str)

	ints := []int{7,2,4}
	sort.Ints(ints) // 对int进行排序
	fmt.Println(ints)

	//判断是否排好序
	alreadySorted := sort.IntsAreSorted(ints)
	fmt.Println(alreadySorted)
}

/**
[a b c]
[2 4 7]
true
 */