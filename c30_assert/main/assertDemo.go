package main

import (
	"fmt"
)

type Point struct {
	x int
	y int
}

func assert1(){
	p1 :=Point{1,2}
	var a interface{} = p1 // 空接口能够接受任意类型
	//var p2 Point = a // 类型不一致，无法直接赋值,但可以通过 a.(Point)形式进行类型断言，实现强转
	var p2 Point = a.(Point)
	fmt.Println(p2)
}

// 只有对 interface类型才能进行类型断言
func assert2(){
	var f1 float32 = 3.2
	var f interface{} = f1
	var f2 float64 = f.(float64)
	fmt.Println(f2)
}

func assert3(){
	var f1 string = "3.2"
	var f interface{} = f1
	f2,ok := f.(float64) // 类型断言检测，OK为布尔类型
	if ok != true {
		panic(fmt.Sprintf("f.(float64) error"))
	}
	fmt.Println(f2)
}

func assert4(){
	var f1 string = "3.2"
	var f interface{} = f1
	// 类型断言检测，OK为布尔类型
	// 上面assert3 类型断言检测 缩写
	if f2,ok := f.(float64);ok != true {
		panic(fmt.Sprintf("f.(float64) error"))
	}else{
		fmt.Println(f2)
	}
}

type Usb interface {
	Start()
	Stop()
}

func Work(usb Usb){
	usb.Start()
	// 如果是Phone时，还需要调用Call()方法
	if phone,ok:=usb.(Phone);ok{
		phone.Calll()
	}
	usb.Stop()
}

type Phone struct {

}

func (p Phone)Calll(){
	fmt.Println("手机打电话")
}

func (p Phone)Start(){
	fmt.Println("手机启动")
}

func (p Phone)Stop(){
	fmt.Println("手机关闭")
}

type Camera struct {

}

func (c Camera)Start(){
	fmt.Println("相机启动")
}

func (c Camera)Stop(){
	fmt.Println("相机关闭")
}

type Mp3 struct {

}

func (m Mp3)Start(){
	fmt.Println("MP3启动")
}

func (m Mp3)Stop(){
	fmt.Println("MP3关闭")
}

func assert5(){
	camera := Camera{}
	Work(camera)

	phone := Phone{}
	Work(phone)

	mp3 := Mp3{}
	Work(mp3)
}

// d.(type) 类型断言只能在switch-case中使用
func TypeJudge(data ... interface{}){
	for i,d := range data{
		switch d.(type) {
		case bool:
			fmt.Printf("第%v个参数是#%T类型，值为：%v\n",i,d,d)
		case string:
			fmt.Printf("第%v个参数是#%T类型，值为：%v\n",i,d,d)
		case float64,float32:
			fmt.Printf("第%v个参数是#%T类型，值为：%v\n",i,d,d)
		case int,int8,int16,int32,int64:
			fmt.Printf("第%v个参数是#%T类型，值为：%v\n",i,d,d)
		case Phone:
			fmt.Printf("第%v个参数是#%T类型，值为：%v\n",i,d,d)
		case *Phone:
			fmt.Printf("第%v个参数是#%T类型，值为：%v\n",i,d,d)
		default:
			fmt.Printf("第%v个参数是未知类型，值为：%v\n",i,d)
		}
	}
}

func assert6(){
	d1 := true
	d2 := 1.2
	var d3 float32 = 5.6
	d4 := 1
	d5 := "hello"
	d6 := Phone{}
	d7 := &Phone{}
	TypeJudge(d1,d2,d3,d4,d5,d6,d7)
}

func main(){
	//assert1()
	//assert3()
	//assert5()
	assert6()
}
