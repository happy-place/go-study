package model

import (
	"fmt"
	"encoding/json"
)

type MethodUtils struct {

}

func (m MethodUtils) PrintRect(){
	for i:=0;i<8;i++{
		for j:=0;j<10;j++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func (mu MethodUtils) PrintRect2(m int,n int){
	for i:=0;i<n;i++{
		for j:=0;j<m;j++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func (m MethodUtils) GetArea(len int,width int) int {
	return len * width
}

func (m *MethodUtils) JudgeNum(n int){
	if n % 2 == 0 {
		fmt.Printf("%v 是偶数\n",n)
	}else{
		fmt.Printf("%v 是奇数\n",n)
	}
}

func (mu *MethodUtils) PrintRect3(m int,n int,key string){
	for i:=0;i<n;i++{
		for j:=0;j<m;j++{
			fmt.Print(key)
		}
		fmt.Println()
	}
}

func (m MethodUtils) MultiPrint(n int){
	for i:=1;i<=n;i++{
		for j:=1;j<=i;j++{
			fmt.Printf("%v x %v = %v ",j,i,i*j)
		}
		fmt.Println()
	}
}

func (m MethodUtils) Reverse(arr [][]int){
	for i:=0;i<len(arr);i++{
		for j:=0;j<=i;j++{
			arr[i][j],arr[j][i] = arr[j][i],arr[i][j]
		}
	}
}

type Calculator struct {

}

func (c *Calculator) Add(n int,m int) {
	fmt.Printf("%v + %v = %v\n",m,n,m+n)
}

func (c *Calculator) Sub(n int,m int) {
	fmt.Printf("%v - %v = %v\n",m,n,m-n)
}

func (c *Calculator) Multi(n int,m int) {
	fmt.Printf("%v x %v = %v\n",m,n,m*n)
}

func (c *Calculator) Divide(n int,m int) {
	fmt.Printf("%v / %v = %v\n",m,n,m/n)
}

func (c *Calculator) Cal(n int,m int,opt string){
	switch opt {
	case "+":
		fmt.Printf("%v + %v = %v\n",m,n,m+n)
	case "-":
		fmt.Printf("%v - %v = %v\n",m,n,m-n)
	case "*":
		fmt.Printf("%v x %v = %v\n",m,n,m*n)
	case "/":
		fmt.Printf("%v / %v = %v\n",m,n,m/n)
	default:
		fmt.Printf("不支持%s运算\n",opt)
	}
}

type Student struct {
	Name string `json:"name"`
	Gender string `json:"gender"`
	Age int `json:"age"`
	Id int `json:"id"`
	Score float64 `json:"score"`
}

func (s *Student) Say() string{
	str,_ := json.Marshal(s) // 接收的是地址，但编译器自己会获取地址对应的对象 - "自动按图索骥"
	return string(str)
}

type Dog struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Weigth float64 `json:"weight"`
}

func (dog *Dog) Say() string {
	str,err := json.Marshal(dog)
	if err !=nil {
		panic(err)
	}
	return string(str)
}

type Box struct {
	Len float64
	Width float64
	Height float64
}

func (box *Box) GetVolumn() float64{
	return box.Width * box.Len * box.Height
}

type Visitor struct {
	Name string
	Age int
}

func (visitor *Visitor) ShowPrice(){
	if visitor.Age <= 18 {
		fmt.Println("年龄不足18岁，免费")
	}else if visitor.Age <= 80{
		fmt.Println("年龄超过18岁，收费20元")
	}else {
		fmt.Println("年龄超过80岁，不建议游玩")
	}
}