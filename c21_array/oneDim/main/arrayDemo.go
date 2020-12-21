package main

import (
	"fmt"
	"math/rand"
	"time"
)
/**
数组：用于批量存储相同类型值 ，有长度限制吗，长度是数组类型一部分
 */

func test1(){
	var hens [6]float64

	hens[0] = 3.0
	hens[1] = 5.0
	hens[2] = 1.0
	hens[3] = 3.4
	hens[4] = 2.0
	hens[5] = 50.0

	totalWeight := 0.0
	for _,n := range hens{
		totalWeight += n
	}
	fmt.Printf("avg weight: %.2f\n",totalWeight/float64(len(hens))) // avg weight: 10.73
}

func test2(){
	var scores [5]float64
	for i:=0;i<len(scores);i++{
		fmt.Printf("请输入第%v个学员成绩；",i+1)
		fmt.Scanf("%f",&scores[i])
	}
	fmt.Println(scores)
}

func createArray(){
	// 初始化数组
	var arr1 [3]int = [3]int{1,2,3}
	fmt.Printf("%v, %T\n",arr1,arr1)
	var arr2 = [3]int{1,2,3}
	fmt.Printf("%v, %T\n",arr2,arr2)
	var arr3 = [...]int{1,2,3}
	fmt.Printf("%v, %T\n",arr3,arr3)
	var arr4 = []int{1,2,3} // 切片
	fmt.Printf("%v, %T\n",arr4,arr4)
	var arr5 = []int{1:1,0:2,2:3} // kv 结构，k表示数组下标 // 切片
	fmt.Printf("%v, %T\n",arr5,arr5)

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)
	fmt.Println(arr5)
}

func arrArgs1(arr [3]int){
	// 接收的是[3]int 数组长度明确，因此是将数组副本传入，是值传递，内部修改数组不影响外部数组
	arr[0] = arr[0] * -1
}

func valuePass(){
	arr := [3]int{1,2,3}
	arrArgs1(arr)
	fmt.Println(arr) // [1 2 3] arrArgs1 接收的是[3]int 数组长度明确，因此是将数组副本传入，是值传递
}

func arrArgs2(arr []int){
	// 接收的是长度不明确数组，传入的是数组引用，内部值改变，会影响外部
	arr[0] = arr[0] * -1
}

func referPass(){
	arr := []int{1,2,3}
	arrArgs2(arr)
	fmt.Println(arr) // [-1 2 3] 接收的是长度不明确数组，传入的是数组引用，内部值改变，会影响外部
}

func arrPointerArgs3(arr *[3]int){
	// 接收的是指针，因此肯定是引用传递
	(*arr)[0] = (*arr)[0] * -1
}

func referPass2(){
	arr := [3]int{1,2,3}
	arrPointerArgs3(&arr)
	fmt.Println(arr) // [-1 2 3] 接收的是长度不明确数组，传入的是数组引用，内部值改变，会影响外部
}

func loopChar(){
	a := int('A')
	var arr [26]byte
	for i:=0;i<26;i++{
		arr[i] = byte(a+i)
		fmt.Printf("%c,",a+i) // A,B,C,D,E,F,G,H,I,J,K,L,M,N,O,P,Q,R,S,T,U,V,W,X,Y,Z,
	}
	fmt.Println()
	for _,e := range(arr){
		fmt.Printf("%c,",e)
	}
	fmt.Println()
}

func max(){
	a := []int {10,2,6,11,12,31,2}
	maxValue := a[0]
	i := 0
	for ;i<len(a);i++{
		if a[i] > maxValue{
			maxValue = a[i]
		}
	}
	fmt.Printf("max-data: %v, max-data-index: %v\n",maxValue,i)
	// max-data: 31, max-data-index: 7 最大值和其下标
}

func sumAndAvg(arr []int){
	sum := 0
	for _,e := range arr {
		sum += e
	}
	avg := sum / (len(arr) * 1.0)
	fmt.Printf("sum: %v, avg: %v, avg-type: %T\n",sum,avg,avg)
	// sum: 6, avg: 2, avg-type: int
}

func reverseArr(){
	var arr [5]int
	for i:=0;i<5;i++{
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(100)
		arr[i] = n
	}
	fmt.Println(arr)

	//for i:=len(arr)-1;i>=0;i--{
	//	fmt.Printf("%v,",arr[i])
	//}
	//fmt.Println()

	length := len(arr)
	for i:=0;i<length/2;i++{
		arr[i],arr[length-1-i] = arr[length-1-i],arr[i]
	}
	fmt.Println(arr)


}

func main(){
	//test1()
	//test2()
	//createArray()
	//valuePass()
	//referPass()
	//referPass2()
	//loopChar()
	//max()
	//sumAndAvg([]int{1,2,3})
	reverseArr()
}
