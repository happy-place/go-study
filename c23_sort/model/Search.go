package model

import "fmt"

/**
	对于无序数组，只能挨个遍历查找
 */
func SearchByCmp(arr []string,target string) bool{
	for _,arr := range arr {
		if arr == target{
			return true
		}
	}
	return false
}

/**
	递归法实现二分查找：
	1）二分查找前提是有序数组；
	2) 找出本次查找的中点，判断目标值与中点大小，从而确定需要搜索哪一半空间，进行下次递归调用；
	3）找到条件时，中点恰好等于目标点
	4）判断不存在条件时，左右游标交叉（右边小于左边）
 */
func BinarySearch(arr []int,leftIndex int,rightIndex int,target int) {
	middle := (rightIndex + leftIndex) / 2
	if leftIndex<=rightIndex {
		if target > arr[middle] {
			BinarySearch(arr,middle+1,rightIndex,target)
		}else if target < arr[middle]{
			BinarySearch(arr,leftIndex,middle-1,target)
		}else {
			fmt.Printf("%v index is %v\n",target,middle)
		}
	}else{
		fmt.Printf("%v is not existed\n",target)
	}
}