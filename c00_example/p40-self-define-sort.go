package main

import (
	"fmt"
	"sort"
)

// 声明自定义类型 ByLength，实质是 []string 类型
type ByLength []string

// 声明 Len()、Swap()、Less() 函数，这三个函数是 sort.Sort 的接口组成
func (s ByLength) Len() int {
	// 数组长度
	return len(s)
}

func (s ByLength) Swap(i,j int) {
	// 交换数组下标元素
	s[i],s[j] = s[j],s[i]
}

func (s ByLength) Less(i,j int) bool{
	// 比较下标元素长短
	return len(s[i]) < len(s[j])
}

func main(){
	fruits := []string{"peach","banana","kiwi"}
	// ByLength(fruits) 直接将数组对象封装成为 ByLength对象
	// 然后由于 ByLength已经实现了sort.Sort的接口，所以可以直接排序
	sort.Sort(ByLength(fruits))
	fmt.Println(fruits)
}

/**
[kiwi peach banana]
 */