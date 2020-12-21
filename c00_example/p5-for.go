package main

import "fmt"

func main(){
	// 自增逻辑放在代码块内部
	i := 1
	for i<=3 {
		fmt.Println(i)
		i++
	}

	// 自增逻辑放在for语句上
	for j:=7;j<=9;j++ {
		fmt.Println(j)
	}

	// break 跳出循环
	for {
		fmt.Println("loop")
		break
	}

	// continue 跳过后面代码，继续向下执行
	for n:=0;n<=5;n++{
		if n % 2 ==0 {
			continue
		}
		fmt.Println(n)
	}

}

