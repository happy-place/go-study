package main

import (
	"fmt"
)

/**
从上到下依次匹配直到匹配上为止，不需要break
switch 表达式 {
	case 表达式1,表达式2,...: // 或的关系
		语句块1
	case 表达式1,表达式2,...:
		语句块2
	.....
	default:
		兜底语句块
}

case 语句是常量时，不能重复（编译报错），变量则无限制
default 不是必须的
*/

func week() {
	var num byte
	fmt.Printf("请输入a~g直接任意字符：")
	fmt.Scanf("%c", &num) // 字符串使用%c接收
	switch num {
	case 'a':
		fmt.Println("周一")
	case 'b':
		fmt.Println("周二")
	case 'c':
		fmt.Println("周三")
	case 'd':
		fmt.Println("周四")
	case 'e':
		fmt.Println("周五")
	case 'f':
		fmt.Println("周六")
	default:
		fmt.Println("周天")
	}
}

func multi(){
	var a int32 = 10
	//var b int64 = 29
	switch a {
	//case b: int64 与 int32 类型不一致不能匹配
	//	fmt.Println("of")
	case 10,20,30: // a 同时符合三者之一，即匹配，但是三个值不能重复
		fmt.Println("matched")
	default:
		fmt.Println("no matched")
	}
}

// switch 可以不写表达式，case 后面可以跟值 或 布尔表达式
func test1(){
	var a int = 10
	switch  {
	case a == 10:
		fmt.Printf("a == %v\n",a)
	case a> 20 && a < 20:
		fmt.Printf("a == %v\n",a)
	default:
		fmt.Printf("default\n")
	}
}

func test2(){
	// switch 进行赋值，然后后面引用,此时a的作用范围仅限于switch语句块
	switch a := 10; {
	case a == 10:
		fmt.Printf("a == %v\n",a)
	case a> 20 && a < 20:
		fmt.Printf("a == %v\n",a)
	default:
		fmt.Printf("default\n")
	}
	//fmt.Println(a) 无法引用到
}

func test3(){
	var num = 10
	switch num {
	case 10:
		fmt.Println("10")
		fallthrough // 默认只穿透一层
	case 20:
		fmt.Println("20")
	case 30:
		fmt.Println("30")
	default:
		fmt.Println("default")
	}
	/**
	10
	20
	*/
}

func typeMatch(x interface{}){
	// switch 用于类型匹配
	switch i:=x.(type) {
	case nil:
		fmt.Printf("x类型：%T\n",i)
	case int:
		fmt.Printf("x类型：%T\n",i)
	case float64:
		fmt.Printf("x类型：%T\n",i)
	case func(int) float64:
		fmt.Printf("x类型：%T\n",i)
	case bool,string:
		fmt.Printf("x类型：%T\n",i)
	default:
		fmt.Println("未知类型")
	}
}

func t(a int) float64{
	return 1.0
}

func test4(){
	var a byte
	fmt.Printf("请输入小写字符：")
	fmt.Scanf("%c",&a)
	switch a {
	case 'a','b','c','d','e':
		fmt.Println(string(a - 32))
	default:
		fmt.Println("other")
	}
}

func test5(){
	var score byte
	fmt.Printf("请输入成绩：")
	fmt.Scanf("%v",&score)
	switch {
	case score > 100 || score < 0:
		fmt.Println("输入分数需要在0~100")
	case score >= 60:
		fmt.Println("成绩合格")
	default:
		fmt.Println("成绩不合格")
	}
}

func season(){
	var month byte
	fmt.Printf("输入月份：")
	fmt.Scanf("%v",&month)
	switch month {
	case 3,4,5:
		fmt.Println("spring")
	case 6,7,8:
		fmt.Println("summer")
	case 9,10,11:
		fmt.Println("autumn")
	default:
		fmt.Println("winter")
	}
}

func launch(){
	var weekday string
	fmt.Printf("输入星期：")
	fmt.Scanf("%s",&weekday)
	switch weekday{
	case "星期一":
		fmt.Println("干煸豆角")
	case "星期二":
		fmt.Println("醋溜土豆")
	case "星期三":
		fmt.Println("红烧狮子头")
	case "星期四":
		fmt.Println("油炸花生米")
	case "星期五":
		fmt.Println("蒜蓉扇贝")
	case "星期六":
		fmt.Println("东北乱炖")
	case "星期日","星期天":
		fmt.Println("大盘鸡")
	}
}


func main(){
	//week()
	//test3()
	//typeMatch(1)
	//typeMatch(23.0)
	//typeMatch("a")
	//typeMatch(t)

	//test4()
	//test5()
	//season()
	launch()
}

