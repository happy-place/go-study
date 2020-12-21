package main

import (
	"fmt"
)

/**
	切片：动态数组，切片是数组的引用，遵循引用传递机制
	遍历元素，下标访问、求切片长度 与数组一致
	切片长度可以变化

	slice 是引用类型，包含三个成员变量，首递增，length长度, capacity 容量
	类似结构：
	type slice struct{
		ptr *[x]int
		len int
		cap int
	}

	切片实质是对数组的引用，因此修改切片元素，会导致底层数组元素改变

 */

func create1(){
	// 基于是数组创建切片，可同时基于数组或切片操作元素，切片ptr指向数组开始切片下标
	s1 := [5]int{1,2,3,4,5}
	s2 := s1[2:4] // [2,4)
	s2 = append(s2,6)
	fmt.Printf("len: %v, cap: %v, %v\n",len(s2),cap(s2),s2) // len: 3, cap: 3, [3 4 6]
}

func create2(){
	// 直接定义切片
	var s1 []int = []int{1,2,3}
	s1 = append(s1, 2)
	s1 = append(s1, 3,4,5) // 追加多个，本质上是拷贝到新数组了
	s2 := []int{6,7,8}
	s1 = append(s1,s2...) // 将切片s2追加到s1上
	fmt.Printf("%v %T\n",s1,s1) // [1 2 3 2] []int

}

func create3(){
	// make([]type,len,[cap]) len <= cap
	s1 := make([]string,3) // 填入的是长度，容量选填，切片底层数组交给make维护，切片ptr指向数组首地址
	s1[0] = "a"
	s1[1] = "b"
	s1[2] = "c"
	s1 = append(s1, "d") // len: 4, cap: 6, [a b c d] cap通常是初始化容量的两倍，随着元素增加，容量分批次递增
	fmt.Printf("len: %v, cap: %v, %v\n",len(s1),cap(s1),s1)

	s2 := make([]int,3,8)
	fmt.Printf("len: %v, cap: %v, %v\n",len(s2),cap(s2),s2) // len: 3, cap: 8, [0 0 0]

}

func forLoop(){
	s1 := make([]int,3,6)
	for i:=0;i<len(s1);i++{
		fmt.Printf("%v,",s1[i]) // 0,0,0,0
	}
	fmt.Println()

	for _,e := range s1{
		fmt.Printf("%v,",e) // 0,0,0,0
	}
	fmt.Println()
}

func slice(){
	// 切片可以继续切片
	s1 := []int{1,2,3,4,5,6,7}
	s2 := s1[2:5]
	s3 := s1[2:]
	s4 := s1[:5]
	s5 := s1[:]
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)
}

// append 之后，slice 底层映射的数组被隐藏，无法直接操作；且此时切片与原先数组无关系
func sliceAppend(){
	var s1 []int = []int{1,2,3}
	s1 = append(s1, 2)
	s1 = append(s1, 3,4,5) // 追加多个，本质上是拷贝到新数组了
	s2 := []int{6,7,8}
	s1 = append(s1,s2...) // 将切片s2追加到s1上
	fmt.Printf("%v %T\n",s1,s1) // [1 2 3 2] []int
}

func sliceCopy(){
	s1 := []int{1,2,3}
	s2 := make([]int,3,6)
	copy(s1,s2) // 将s2拷贝到s1, 深拷贝，s1 与 s2 无关，拷贝者与被拷贝者必须都是切片
	/**
		1.拷贝者与被拷贝者必须都是切片;
		2.深拷贝，s1 与 s2 无关
	 */
	fmt.Println(s2)
	s2[1] = 100
	fmt.Println(s1)
	fmt.Println(s2)
}

func copy2(){
	s1 := []int{1,2,3,4,5}
	s2 := make([]int,1)
	s2[0] = 100
	fmt.Println(s2)
	copy(s1,s2) // copy(dst,src)
	fmt.Println(s1) // [100 2 3 4 5] 小往大拷贝：有几个就覆盖几个

	s1[0] = 200
	copy(s2,s1)
	fmt.Println(s2) // [200] 大往小拷贝：能拷贝几个就拷贝几个
}

func test(s []int){
	s[0] = 100
}

func testRefer(){
	arr := [5]int{1,2,3,4,5}
	s1 := arr[:]
	test(s1) // 传递引用
	fmt.Println(s1) // [100 2 3 4 5]

}

/**
	1）string 本质上是byte数组，因此可以进行切片处理
	2）数据结构类似于
	type string struct{
		ptr *[5]byte
		len int
	}
	3）string元素不可修改 str[0] = 'a' 报错,如果必须修改，只能先将string转为byte数组，然后修改byte[]数组部分元素，最后将byte数组转换string
 */
func stringSlice(){
	str := "hello go"
	s1 := str[2:5]
	fmt.Println(s1) // llo

	// 修改纯ascii编码字符串
	b1 := []byte(str)
	b1[0] = 'z'
	str = string(b1)
	fmt.Println(str) // zello go

	// 修改包含中文字符串
	str1 := "hello 中国"
	b2 := []rune(str1) // 兼容utf-8编码
	b2[6] = '美'
	str2 := string(b2)
	fmt.Println(str2) // hello 美国

}

// 递归
func fbn1(n uint64)[]uint64{
	if n==0 {
		return []uint64{1}
	}else if n == 1{
		return []uint64{1,1}
	}
	s1 := fbn1(n-1)
	s2 := fbn1(n-2)
	s1 = append(s1,s1[len(s1)-1] + s2[len(s2)-1])
	return s1
}

// 递推
func fbn2(n uint64)[]uint64{
	s := make([]uint64,n)
	s[0] = 1
	s[1] = 1
	for i:=uint64(2);i<n;i++{
		s[i] = s[i-1] + s[i-2]
 	}
 	return s
}


func main(){
	//create1()
	//create2()
	//create3()
	//forLoop()
	//slice()
	//sliceAppend()
	//sliceCopy()
	//copy2()
	//testRefer()
	//stringSlice()
	s := fbn1(7)
	fmt.Println(s)

	s1 := fbn2(7)
	fmt.Println(s1)
}




