package model

import (
	"fmt"
	"math"
)

func Test1(a int32,b int32){
	if a + b >= 50 {
		fmt.Println("hello world ",)
	}
}

func Test2(a float64,b float64) {
	if a > 10.0 && b < 20.0 {
		fmt.Printf("%v + %v = %v",a,b,a+b)
	}
}

func Test3(a int,b int) {
	var c = a + b
	if c % 3 == 0 && c % 5 ==0 {
		fmt.Println("能被上3整除，又能被5整除")
	}
}

func Test4(year int){
	if (year % 4==0 && year %100 !=0)|| (year % 400 == 0) {
		fmt.Printf("%v是闰年",year)
	} else {
		fmt.Printf("%v是平年",year)
	}
}

func Test5(score int){

	if score == 100 {
		fmt.Println("奖励一辆BMW")
	}else if score >80 && score <= 90 {
		fmt.Println("奖励一台iphone7 plus")
	}else if score >60 && score <= 80 {
		fmt.Println("奖励一个iPads")
	}else{
		fmt.Println("奖励一顿板子")
	}
}

func Test6(a int,b int,c int) {
	// a*x^2 + b*x + c = 0
	dd := float64(b * b - 4 * a * c)
	if dd >= 0 {
		x1 := (float64(-b) + math.Sqrt(dd)) / float64(2*a)
		x2 := (float64(-b) - math.Sqrt(dd))/ float64(2*a)
		fmt.Printf("%v*x^2 + %v*x + %v = 0 方程的两个解是: x1=%f, x2=%f\n", a,b,c,x1,x2)
	}else if dd == 0 {
		x1 := float64(-b) / float64(2*a)
		fmt.Printf("%v*x^2 + %v*x + %v = 0 方程的只有一个解是: x1=%f\n", a,b,c,x1)
	} else {
		fmt.Println("方程无解")
	}

}

func Test7(){
	var speed int
	var gender string
	fmt.Printf("输入速度：")
	fmt.Scanf("%v",&speed)
	if speed < 8 {
		fmt.Printf("输入性别：")
		fmt.Scanf("%s",&gender)
		if gender == "男" {
			fmt.Println("进入男子决赛组")
		}else{
			fmt.Println("进入女子决赛组")
		}
	}
}

func Test8(month byte,age byte) {
	if month >=4 && month <= 10 {
		if age >= 60 {
			fmt.Printf("1/3价票：%v\n",60/3)
		} else if age >= 18 && age < 60{
			fmt.Printf("全价票：%v\n",60)
		} else{
			fmt.Printf("半价票：%v\n",30)
		}
	}else {
		if age >= 18 && age < 60{
			fmt.Printf("票价：%v\n",40)
		} else{
			fmt.Printf("票价：%v\n",20)
		}
	}
}





