package main

import "fmt"

// 逐个字段录入
func scanln(){
	var name string
	var age byte
	var sal float32
	var isPass bool
	fmt.Print("请输入姓名：")
	fmt.Scanln(&name)
	fmt.Print("请输入年龄：")
	fmt.Scanln(&age)
	fmt.Print("请输入薪水：")
	fmt.Scanln(&sal)
	fmt.Print("是否通过考试：")
	fmt.Scanln(&isPass)
	fmt.Printf("姓名\t年龄\t薪水\t考试\n%s\t%v\t%f\t%v\n",name,age,sal,isPass)
}

// 一次性录入然后解析
func scanf(){
	var name string
	var age byte
	var sal float32
	var isPass bool
	fmt.Println("请输入'姓名 年龄 薪水 考试通过与否'")
	fmt.Scanf("%s %d %f %t",&name,&age,&sal,&isPass)
	fmt.Printf("姓名\t年龄\t薪水\t考试\n%s\t%v\t%f\t%v\n",name,age,sal,isPass)
}

func main(){
	//scanln()
	scanf()
}

/**
姓名    年龄    薪水    考试
tom     12      2300.100098     true
 */