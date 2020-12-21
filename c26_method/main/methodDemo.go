package main

import (
	"fmt"
	"math"
	"go_study/c26_method/model"
)

/**
	1）结构体类型是值类型，默认方法中是值传递
	2）如果想引用传递，请传指针
	3）不光结构体，其他所有类型都能以绑定方法方式，进行拓展
	4) 方法名大小写分别对应public、private 访问控制符

	函数与方法区别：
	1）如果是函数，形参是什么，就必须传什么，如果是方法，既可以用结构体变量调用，又可以用结构体指针调用
	2）方法与值绑定-值拷贝 (p Person)，方法与引用绑定-引用拷贝 (p *Person)
 */
type Person struct {
	Name string
}

func (p Person) speak(){
	fmt.Printf("%s is a good man\n",p.Name)
}

func (p Person) jisuan() int {
	sum := 0
	for i:=1;i<=1000;i++{
		sum += i
	}
	return sum
}

func (p Person) jisuan2(n int) int {
	sum := 0
	for i:=1;i<=n;i++{
		sum += i
	}
	return sum
}

func(p Person) getSum(a int,b int) int {
	return a + b
}

type Circle struct {
	radius float64
}

func (c Circle) getArea() float64{
	return math.Pi * math.Pow(c.radius,2)
}

func (c Circle) getInfo() (float64,float64){
	mianJi := math.Pi * math.Pow(c.radius,2)
	zhouChang := math.Pi * c.radius * 2
	return mianJi,zhouChang
}

type Float float64
func (n Float) intPrint(){
	fmt.Println(n)
}

func (n *Float) intPrint2(){
	fmt.Println(*n)
}

func (p Person) String() string{
	return fmt.Sprintf("name: %v",p.Name)
}

func main(){
	p1 := Person{"tom"}
	p1.speak()
	fmt.Println(p1.jisuan())
	fmt.Println(p1.jisuan2(100))
	fmt.Println(p1.getSum(100,200))

	c1 := Circle{2.0}
	fmt.Printf("area: %.2f\n",c1.getArea())
	mianJi,zhouChang := c1.getInfo()
	fmt.Printf("面积: %.2f，周长：%.2f\n",mianJi,zhouChang)

	var f1 Float = 3.0
	f1.intPrint()
	(&f1).intPrint2()

	fmt.Println(p1.String())

	m := model.MethodUtils{}
	m.PrintRect()
	m.PrintRect2(6,3)
	area := m.GetArea(3,5)
	fmt.Printf("%v x %v = %v\n",3,5,area)

	(&m).JudgeNum(100)
	(&m).PrintRect3(20,7,"#")

	cal := model.Calculator{}
	cal.Add(1,2)
	cal.Cal(1,2,"+")

	m.MultiPrint(9)

	arr := [][]int{{1,2,3},{4,5,6},{7,8,9}}
	fmt.Println(arr)
	m.Reverse(arr)
	fmt.Println(arr)

	stu := model.Student{"tom","male",12,100,60.3}
	fmt.Println(stu.Say())

	dog := model.Dog{"jerry",3,12.3}
	fmt.Println(dog.Say())

	box := model.Box{1,2,3}
	fmt.Println(box.GetVolumn())

	v1 := model.Visitor{"tom",23}
	v1.ShowPrice()
	v2 := model.Visitor{"tom",16}
	v2.ShowPrice()
	v3 := model.Visitor{"tom",81}
	v3.ShowPrice()

}
