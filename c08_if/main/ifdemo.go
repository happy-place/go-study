package main

import (
	"fmt"
	"go_study/c08_if/model"
)

func main() {
	var a int
	fmt.Printf("请输入整数：")
	fmt.Scanln(&a)
	if a > 10 {
		fmt.Printf("%v > 10\n", a)
	}

	if a > 5 {
		fmt.Printf("%v > 5\n", a)
	} else {
		fmt.Printf("%v <= 5\n", a)
	}

	if a > 5 {
		fmt.Printf("%v > 5\n", a)
	} else if a == 5 {
		fmt.Printf("%v = 5\n", a)
	} else {
		fmt.Printf("%v < 5\n", a)
	}

	model.Test1(20,31)
	model.Test2(20,19)
	model.Test3(2,13)
	model.Test4(200)

	fmt.Printf("请输入小明分数：")
	fmt.Scanf("%v",&a)
	model.Test5(a)

	fmt.Printf("请输入二元一次方程a*x^2 + b*x + c = 0 的系数项：")
	var b int
	var c int
	fmt.Scanf("%v %v %v",&a,&b,&c)
	model.Test6(a,b,c)

	model.Test7()

}
