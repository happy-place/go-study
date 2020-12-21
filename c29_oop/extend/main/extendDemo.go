package main

import (
	"fmt"
	"go_study/c29_oop/extend/model"
)

func extend1(){
	// 直接通过构造器创建
	pupil := model.Pupil{model.Student{"tom", 23, 60}}
	pupil.Testing()
	pupil.SetScore(90)
	pupil.ShowInfo()

	// 先获取指针，然后逐个初始化，最后调用
	pupil2 := &model.Pupil{}
	pupil2.Student.Name = "tom"
	pupil2.Student.Age = 21
	pupil2.Student.Score = 67
	pupil2.Testing()
	pupil2.SetScore(90)
	pupil2.ShowInfo()

}

func extend2(){
	granduate := model.Granduate{model.Student{"tom", 31, 200}}
	granduate.Testing()
	granduate.SetScore(90)
	granduate.ShowInfo()
}

type A struct{
	Name string
	age int
}

func (a *A) SayOk(){
	fmt.Println("A say ok",a.Name)
}

func (a *A)hello(){
	fmt.Println("A say hello",a.Name)
}

// 就近访问原则，现在本级找，找不到就去父类找
// B 没有Name字段时，b.Name 会自动识别成其嵌套匿名结构体的Name，即 b.A.Name
// B 有自己的Name时，b.Name 就指自己的Name，b.A.Name 是A的Name
type B struct {
	A
	Name string
}

// 继承后，子类能访问父类的私有属性和方法
func extend3(){
	b := &B{A{"a",21},"jerry"}
	fmt.Printf("b.Name=%s, b.A.Name=%s\n",b.Name,b.A.Name)
	fmt.Println("age=",b.A.age)
	b.hello()
	b.SayOk()
}


type AA struct {
	Name string
	age int
}

type BB struct {
	Name string
	age int
}

// 多重继承，嵌套了两个匿名结构体
type CC struct {
	AA
	BB
}

func multiExtend(){
	c := CC{AA{"tom",10},BB{"henrry",20}}
	//fmt.Println(c.Name,c.age) 报错，编译不通过，无法确定从哪里取，必须明确指出嵌套匿名结构体
	fmt.Println(c.AA.Name,c.AA.age)
	fmt.Println(c.BB.Name,c.BB.age)
}

// 嵌套了有名结构体，称为组合关系，则访问有名结构体必须带上结构体名称 c.a.xx
type C struct {
	a A
}

func namedSupper(){
	c:= C{A{"tom",21}}
	fmt.Println(c.a.age)
	fmt.Println(c.a.Name)
	c.a.SayOk()
	c.a.hello()
}

type DD struct {
	*AA
	*BB
}

func initWhenCreate(){
	cc := CC{AA{"tom",10},BB{"jerry",22}}
	fmt.Println(cc.AA.Name,cc.AA.age,cc.BB.Name,cc.BB.age)
	fmt.Printf("%+v\n",cc)

	dd := DD{&AA{"tom",10},&BB{"jerry",22}}
	fmt.Println(dd.AA.Name,dd.AA.age,dd.BB.Name,dd.BB.age)
	fmt.Println((*dd.AA).Name,(*dd.AA).age,(*dd.BB).Name,(*dd.BB).age) // 等效于上面的
	fmt.Printf("%+v\n",dd)
}

type Monster struct {
	Name string
	Age int
}

// 结构体中不光可以嵌入另外结构体，还可以嵌入基本数据类型
type E struct {
	Monster
	int
	n int // 多个int字段，后面的int必起取名字
}

func basicTypeAsAmbigusField(){
	e := E{Monster{"tom",23},10,20}
	fmt.Printf("%+v\n",e)
	fmt.Println(e.int)
	fmt.Println(e.Monster)
	fmt.Println(e.n)
}


func main(){
	//extend1()
	//extend2()
	//extend3()
	//multiExtend()
	//namedSupper()
	initWhenCreate()
	basicTypeAsAmbigusField()
}
