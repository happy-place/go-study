package model

import "fmt"

func swap(arr []int,i int,j int){
	arr[i],arr[j] = arr[j],arr[i]
}

/**
	冒泡法排序
	1）外层每次循环找到一个最大值，并移动到当前搜索区间末尾；
	2）随着极值逐渐被找出，每次外层循环对应的内层循环，搜索区间逐渐收敛
 */
func BubbleSort(arr []int){
	fmt.Println(arr)
	for i:=0;i<len(arr);i++{
		for j:=0;j<len(arr) - i -1;j++{
			if arr[j]>arr[j+1] {
				swap(arr,j,j+1)
			}
		}
	}
	fmt.Println(arr)
}
