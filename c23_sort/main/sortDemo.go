package main

import (
	"fmt"
	"go_study/c23_sort/model"
)

func bubbleSortTest(){
	arr := []int{3,1,5,6,3,2,11,12}
	model.BubbleSort(arr)
}

func findByCmp(){
	sources := []string{"apple","orange","grape"}
	var target string
	for{
		fmt.Printf("请输入目标：")
		fmt.Scanf("%s",&target)
		existed := model.SearchByCmp(sources,target)
		fmt.Printf("%s existed in sources %v\n",target,existed)
	}
}

func binarySearch(){
	arr := []int{1,4,5,7,10,12,54}
	var target int
	for{
		fmt.Printf("请输入目标：")
		fmt.Scanf("%v",&target)
		model.BinarySearch(arr,0,len(arr)-1,target)
	}
}

func main(){
	//bubbleSortTest()
	//findByCmp()
	binarySearch()
}
