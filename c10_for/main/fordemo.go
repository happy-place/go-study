package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 标准写法
func print(){
	for i:=0;i<10;i++{
		fmt.Println("hello world")
	}
}

func print2(){
	i := 0
	for i< 10 {
		fmt.Println("hello world")
		i ++
	}
}

func print3(){
	i := 0
	for {
		if i < 10{
			fmt.Println("hello world")
			i ++
		}else{
			break
		}
	}
}

func print4(){
	for i:=0;i<10;i++ {
		if i % 2 ==0 {
			continue
		}
		fmt.Println("hello world")
	}
}

func print5(){
	// for-range 是按字符方式进行遍历，因此有中文也是可以的
	var str = "hello world中国" // range 可以自动识别字符时ascii编码,还是utf-8编码，中文可以正常输出
	for i,s := range str {
		fmt.Printf("第 %v 个字符是 %c\n",i,s)
	}
}

func print6(){
	// str[i] 字符数组切片
	var str = "hello world中国"
	var str2 = []rune(str) // 必须强制转为 []rune 才能避免中英文混写解析异常
	for i:=0;i<len(str2);i++  {
		fmt.Printf("第 %v 个字符是 %c\n",i,str2[i])
	}
}

func print7(){
	var count uint64 = 0
	var sum uint64 = 0
	var i uint64 = 9
	for ;i<=100;i+=9{
		count ++
		sum += i
	}
	fmt.Printf("1~100之间9的倍数个数有%v个，其累计和为：%v\n",count,sum) // 1~100之间9的倍数个数有11个，其累计和为：594
}

func print8(){
	var n int = 6
	for i:=0;i<=n;i++{
		fmt.Printf("%v + %v = %v\n",i,n-i,n)
	}
}

func while(){
	var i = 0
	for {
		if i>=5{
			break
		}
		fmt.Println("hello world")
		i ++
	}
}

func doWhile(){ // 至少执行一次
	var i = 5
	for{
		fmt.Println("hello world")
		i ++
		if i >= 5 {
			break
		}
	}
}

func test1(){
	classes := 3
	kids := 5
	for i:= 0;i<classes;i++{
		var sum = 0
		for j:=0;j<kids;j++{
			var score int
			fmt.Scanf("%v",&score)
			sum += score
		}
		fmt.Printf("%v班平均分：%v \n",i+1,sum/(kids*1.0))
	}
}

func test2(){
	classes := 3
	kids := 5
	for i:= 0;i<classes;i++{
		var passCount = 0
		for j:=0;j<kids;j++{
			var score int
			fmt.Scanf("%v",&score)
			if score>=60{
				passCount ++
			}
		}
		fmt.Printf("%v班及格人数：%v \n",i+1,passCount)
	}
}

func test3(){
	height := 100
	length :=  2 * height +1

	for i:=0;i<height;i++{
		stars := 2 * i +1
		blanks_half := (length - stars)/2
		for j:=0;j<blanks_half;j++{
			fmt.Print(" ")
		}
		for j:=blanks_half;j<blanks_half+stars;j++{
			fmt.Print("*")
		}
		for j:=blanks_half+stars;j<length;j++{
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func test4(){
	for i:=1;i<=9;i++{
		for j:=1;j<=i;j++{
			fmt.Printf("%v*%v=%v\t",j,i,i*j)
		}
		fmt.Println()
	}
}

func test5(){
	count :=0
	for{
		rand.Seed(time.Now().Unix()) // 重复太严重了，很难碰到99
		rand.Seed(time.Now().UnixNano()) // 重复，很快就找到99
		intn := rand.Intn(100)+1
		count ++
		if intn == 99 {
			fmt.Printf("直到首次生成99，经历了%v次随机数计算\n",count)
			break
		}
	}
}

func test6(){
	lable1:
	for i:=0;i<3;i++{
		for j:=0;j<3;j++{
			if j==2{
				fmt.Printf("i=%v,j=%v",i,j)
				break lable1
			}
		}
	}
}

func test7(){
	sum := 0
	recored := false
	for i:=1;i<=100;i++{
		sum += i
		if sum > 20 && !recored{
			fmt.Printf("i=%v,sum=%v",i,sum)
			recored = true
		}
	}
}

func test8(){
	pass := "tom"
	word := "888"
	cnt := 3
	for i:=0;i<cnt;i++ {
		var e_pass string
		var e_word string
		fmt.Println("请输入用户名 密码")
		fmt.Scanf("%s %s",&e_pass,&e_word)
		if e_pass == pass && e_word == word {
			fmt.Println("登录成功")
			break
		} else {
			fmt.Printf("登录失败，还剩%v次机会\n",cnt-i-1)
		}
	}
}

func test9(){
	for i:=0;i<3;i++{
		for j:=0;j<3;j++{
			if j==2 {
				continue
			}
			fmt.Println("i=",i,"j=",j)
		}
	}
}

func test10(){
	here:
	for i:=0;i<3;i++{
		for j:=0;j<3;j++{
			if j==1 {
				continue here
			}
			fmt.Println("i=",i,"j=",j)
		}
	}
}

func test11(){
	for i:=1;i<=100;i++{
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%v",i)
	}
}

func test12(){
	positive := 0
	negative := 0
	for {
		var num int
		fmt.Scanf("%v",&num)
		if num > 0 {
			positive++
		}else if num ==0 {
			break
		}else{
			negative++
		}
		fmt.Printf("输入正数个数；%v，负数个数：%v\n",positive,negative)
	}
}

func test13(){
	var total float64 = 100000
	cnt := 0
	for{
		if total>50000 {
			total -= total * 0.05
			cnt ++
		}else{
			if total>=1000{
				total -= 1000
				cnt ++
			}else{
				fmt.Printf("总共经过了%v次,剩余金额：%v\n",cnt,total)
				break
			}
		}
	}
}

func test14(){
	// goto 直接跳转到label位置，破坏了程序调用流程，不推荐使用
	fmt.Println("hello 1")
	goto label
	fmt.Println("hello 2")
	fmt.Println("hello 3")
	fmt.Println("hello 4")
	label:
	fmt.Println("hello 5")
	fmt.Println("hello 6")
}
/**
hello 1
hello 5
hello 6
 */


func main(){
	//print()
	//print2()
	//print3()
	//print4()
	//print5()
	//print6()
	//print7()
	//print8()
	//while()
	//doWhile()
	//test1()
	//test2()
	//test3()
	//test4()
	//test5()
	//test6()
	//test7()
	//test8()
	//test9()
	//test10()
	//test11()
	//test12()
	//test13()
	test14()
}
