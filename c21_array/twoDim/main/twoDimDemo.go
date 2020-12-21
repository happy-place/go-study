package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

/**
	二维数组
 */
func mutilArr(){
	var arr [4][6]int

	arr[0] = [6]int{0,0,0,0,0,0}
	arr[1] = [6]int{0,0,1,0,0,0}
	arr[2] = [6]int{0,2,0,3,0,0}
	arr[3] = [6]int{0,0,0,0,0,0}

	for i:=0;i<len(arr);i++{
		for j:=0;j<len(arr[i]);j++{
			fmt.Printf("%v ",arr[i][j])
		}
		fmt.Println()
	}

	fmt.Println(arr)

	for _, subArr:=range arr{
		for _,e := range subArr{
			fmt.Printf("%v ",e)
		}
		fmt.Println()
	}

	arr2 := [2][3]int{{1,2,3},{4,5,6}}
	fmt.Println(arr2) // [[1 2 3] [4 5 6]]
}

func avgScore(){
	var scores [3][5]int
	for i:=0;i<3;i++{
		for j:=0;j<5;j++{
			var score int
			fmt.Printf("输入第%v班第%v位学员成绩：",i+1,j+1)
			fmt.Scanf("%v",&score)
			scores[i][j] = score
		}
	}

	totalSum := 0
	for i,arr := range(scores) {
		sum := 0
		for _,e := range arr {
			totalSum += e
			sum += e
		}
		fmt.Printf("第%v班平均成绩：%v\n",i+1,sum /len(arr))
	}
	fmt.Printf("所有学员平均成绩：%v\n",totalSum / (len(scores) * len(scores[0])))
}

func test1(){
	const length = 10
	var nums [length]int
	for i:=0;i<length;i++{
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(100) + 1
		nums[i] = n
	}
	fmt.Println(nums)

	for i:=0;i<length/2;i++{
		nums[i],nums[length-1-i] = nums[length-1-i],nums[i]
	}
	fmt.Println(nums)
	var maxAndMin [2][2]int
	// max
	maxAndMin[0] = [2]int{nums[0],0}
	// main
	maxAndMin[1] = [2]int{nums[0],0}
	exist := false
	for i,e := range(nums){
		if e > maxAndMin[0][0]{
			maxAndMin[0][0] = e
			maxAndMin[0][1] = i
		}
		if e < maxAndMin[1][0] {
			maxAndMin[1][0] = e
			maxAndMin[1][1] = i
		}
		if !exist && e == 55 {
			exist = true
		}
	}
	fmt.Printf("最大值:%v 对应下标:%v\n",maxAndMin[0][0],maxAndMin[0][1])
	fmt.Printf("最小值:%v 对应下标:%v\n",maxAndMin[1][0],maxAndMin[1][1])
	fmt.Printf("是否存在55 %v",exist)
}

/**
 往已经存在升序数组中插入一个元素
 */
func test2(){
	arr := [5]int{1,2,4,5,6}
	var arr2 [6]int
	target := 3
	inserted := false
	for i:=0;i<len(arr2);i++ {
		if !inserted{
			if arr[i]<target {
				arr2[i] = arr[i]
			}else{
				arr2[i] = target
				inserted = true
			}
		} else {
			arr2[i] = arr[i-1]
		}
	}
	fmt.Println(arr2) // [1 2 3 4 5 6]
}

/**
3行4列数组，周边清零
 */
func test3(){
	var arr [3][4]int
	for i,ar := range arr {
		for j,_ := range ar {
			var num int
			fmt.Printf("请输入第%v行第%v个元素：",i+1,j+1)
			fmt.Scanf("%v",&num)
			arr[i][j] = num
		}
	}

	for _,ar := range arr{
		fmt.Println(ar)
	}

	for i,ar := range arr {
		for j,_ := range ar {
			if i * j * (i-2) * (j-3) == 0{
				arr[i][j] = 0
			}
		}
	}

	for _,ar := range arr{
		fmt.Println(ar)
	}
}

func test4(){
	var arr [4][4]int
	for i,ar := range arr {
		for j,_ := range ar {
			var num int
			fmt.Printf("请输入第%v行第%v个元素：",i+1,j+1)
			fmt.Scanf("%v",&num)
			arr[i][j] = num
		}
	}
	fmt.Println(arr)

	for i:=0;i<len(arr)/2;i++{
		arr[i],arr[len(arr)-1-i] = arr[len(arr)-1-i],arr[i]
	}
	fmt.Println(arr)

}

func test5(){
	num := [5]int{1,3,5,7,9}
	fmt.Println(num)
	for i:=0;i<len(num)/2;i++{
		num[i],num[len(num)-1-i] = num[len(num)-1-i],num[i]
	}
	fmt.Println(num)
}

func test6(){
	arr := [10]string{"a","b","aa","BB","Ac","dd","ee","ef","aa","er"}
	existed := false
	target := "aa"
	for i,e := range arr {
		if e == target {
			existed  =true
			fmt.Printf("第%v个元素存在%v\n",i,target)
		}
	}
	if !existed {
		fmt.Printf("不存在%v\n",target)
	}
}

func test7(){
	const length = 10
	var arr [length]int
	for i:=0;i<length;i++{
		rand.Seed(time.Now().UnixNano())
		arr[i] = rand.Intn(100) + 1
	}
	fmt.Println(arr)
	for i:=0;i<len(arr);i++{
		for j:=0;j<len(arr)-1-i;j++{
			if arr[j]>arr[j+1]{
				arr[j],arr[j+1] = arr[j+1],arr[j]
			}
		}
	}
	fmt.Println(arr)
	var find func (arr [length]int,leftIndex int,rightIndex int,target int)

	find = func (arr [length]int,leftIndex int,rightIndex int,target int){
		middle := (leftIndex+rightIndex)/2
		if leftIndex<=rightIndex{
			if target < arr[middle] {
				find(arr,leftIndex,middle-1,target)
			}else if target > arr[middle]{
				find(arr,middle+1,rightIndex,target)
			}else {
				fmt.Printf("arr[%v] is %v\n",middle,target)
			}
		}else{
			fmt.Printf("%v is not existed in arr\n",target)
		}
	}

	find(arr,0,len(arr)-1,90)
}

func test8(){
	arr :=[]int{3,4,6,1,6,7,8,9,12}
	res := [2][2]int{{arr[0],0},{arr[0],0}}
	for i:=1;i<len(arr);i++{
		if arr[i]>res[0][0]{
			res[0][0] = arr[i]
			res[0][1] = i
		}
		if arr[i]<res[1][0]{
			res[1][0] = arr[i]
			res[1][1] = i
		}
	}
	fmt.Printf("最大值：%v,下标为：%v\n",res[0][0],res[0][1])
	fmt.Printf("最小值：%v,下标为：%v\n",res[1][0],res[1][1])
}

func test9(){
	arr := []int{3,4,6,1,6,7,8,9,12}

	sum := 0
	for _,e := range arr {
		sum += e
	}
	avg :=  sum/(len(arr)*1.0)
	fmt.Printf("平均分：%v\n",avg)
	gtAvgCnt := 0
	eqAvgCnt := 0
	ltAvgCnt := 0
	for _,e := range arr {
		if e>avg {
			gtAvgCnt ++
		}else if e==avg{
			eqAvgCnt ++
		}else{
			ltAvgCnt ++
		}
	}

	fmt.Printf("超过平均分人数：%v\n",gtAvgCnt)
	fmt.Printf("等于平均分人数：%v\n",eqAvgCnt)
	fmt.Printf("小于平均分人数：%v\n",ltAvgCnt)
}

func test10(){
	var arr [8]float64
	maxAndMin := [2][2]float64{{-1,0},{11,0}}
	var sum = 0.0
	for i:=0;i<len(arr);i++{
		fmt.Printf("请输入第%v为评委打分：",i+1)
		var score float64
		fmt.Scanf("%v",&score)
		arr[i] = score
		sum += score
		if score > maxAndMin[0][0] {
			maxAndMin[0][0] = score
			maxAndMin[0][1] = float64(i)
		}
		if score < maxAndMin[1][0]{
			maxAndMin[1][0] = score
			maxAndMin[1][1] = float64(i)
		}
	}

	avg := (sum - maxAndMin[0][0] - maxAndMin[1][0])/(float64(len(arr)) - 2)

	fmt.Println("评委打分：",arr)
	fmt.Printf("最高分: %v,由第%.0f位评委打出\n",maxAndMin[0][0]+1,maxAndMin[0][1])
	fmt.Printf("最低分: %v,由第%.0f位评委打出\n",maxAndMin[1][0]+1,maxAndMin[1][1])
	fmt.Printf("平均分：%.2f\n",avg)

	goodAndBad := [2][2]float64{{0,0},{0,0}}
	for i,e := range arr {
		delta := math.Abs(e-avg)
		if goodAndBad[0][0] > delta {
			goodAndBad[0][0] = delta
			goodAndBad[0][1] = float64(i)
		}
		if goodAndBad[1][0] < delta {
			goodAndBad[1][0] = delta
			goodAndBad[1][1] = float64(i)
		}
	}
	fmt.Printf("最佳评委是第%.0f位，与均值差距为：%.2f\n",goodAndBad[0][0]+1,goodAndBad[0][1])
	fmt.Printf("最差评委是第%.0f位，与均值差距为：%.2f\n",goodAndBad[1][0]+1,goodAndBad[1][1])

}


func main(){
	//mutilArr()
	//avgScore()
	//test1()
	//test2()
	//test3()
	//test4()
	//test5()
	//test6()
	//test7()
	//test8()
	//test9()
	test10()
}
