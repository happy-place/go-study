package main

import "fmt"

func main(){
	// 切片 slice，与 数组非常相似，但又具微小区别
	s := make([]string,3)
	fmt.Println(s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set:",s)
	fmt.Println("get:",s[2])
	fmt.Println("len:",len(s))

	// 追加
	s = append(s,"d")
	s = append(s,"e","f")
	fmt.Println("append:",s)

	// 复制
	c := make([]string,len(s))
	copy(c,s)
	fmt.Println("copy:",c)

	// 切片
	l := s[2:5]
	fmt.Println("sl1",l)

	l = s[:5]
	fmt.Println("sl2",l)

	l = s[2:]
	fmt.Println("sl3",l)

	t := []string{"g","h","i"}
	fmt.Println("dcl:",t)

	// 不定长二位数组
	twoD := make([][]int ,3)
	for i:=0;i<3;i++{
		innerLen := i +1
		twoD[i] = make([]int,innerLen)
		for j :=0;j<innerLen;j++{
			twoD[i][j] = i + j
		}
	}

	fmt.Println(twoD)

}

/**
[  ]
set: [a b c]
get: c
len: 3
append: [a b c d e f]
copy: [a b c d e f]
sl1 [c d e]
sl2 [a b c d e]
sl3 [c d e f]
dcl: [g h i]
[[0] [1 2] [2 3 4]]
 */