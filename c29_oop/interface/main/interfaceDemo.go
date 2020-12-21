package main

import (
	"fmt"
	"sort"
)

// 声明接口，里面包含了两个抽象方法
type Usb interface {
	Start()
	Stop()
}

// 声明结构体，通过绑定接口的抽象方法，实现接口
type Phone struct {
}

func (p Phone) Start() {
	fmt.Println("手机开始工作")
}

func (p Phone) Stop() {
	fmt.Println("手机停止工作")
}

// 声明结构体，通过绑定接口的抽象方法，实现接口
type Camera struct {
}

func (p Camera) Start() {
	fmt.Println("相机开始工作")
}

func (p Camera) Stop() {
	fmt.Println("相机停止工作")
}

// 声明结构体，绑定自己的方法，方法接收接口类型参数
type Computer struct {
}

// 相当于 Phone 和 Camera同时实现了Usb 和 Usb2接口
type Usb2 interface {
	Start()
	Stop()
}

// 传入usb不同实现类，实现多态效果
func (c Computer) Working(usb Usb) {
	usb.Start()
	usb.Stop()
}

// 传入usb不同实现类，实现多态效果
func (c Computer) Working2(usb Usb2) {
	usb.Start()
	usb.Stop()
}

type AInterface interface {
	SayHi()
}

// 即便基本类型，也可以通过实现接口，扩展功能
type integer int

func (i integer)SayHi(){
	fmt.Println("Hi,",i)
}

// 结构体同时实现AInterface和BInterface接口
// 接口中只能包含抽象方法，不能包含属性字段
type BInterface interface {
	Hello()
}

type Monster struct {

}

func (m Monster)SayHi(){
	fmt.Println("monster say hi")
}

func (m Monster)Hello(){
	fmt.Println("monster say hello")
}

type CInterface interface {
	SayGoodBy()
}

// 接口通过嵌入隐式接口，实现接口继承效果，其实现类，需要实现自己的方法和继承接口发方法
type DInterface interface {
	BInterface
	CInterface
	Eat()
}

type Student struct {
	Name string
}

func (s Student)SayGoodBy(){
	fmt.Printf("%s say goodBy\n",s.Name)
}

func (s Student)Hello(){
	fmt.Printf("%s say hello\n",s.Name)
}

func (s Student)Eat(){
	fmt.Printf("%s eat something\n",s.Name)
}

func live(d DInterface){
	d.Hello()
	d.SayGoodBy()
	d.Eat()
}

// 空接口，所有类型都是空接口的实现类
type EInterface interface {

}

func show(e EInterface){
	fmt.Printf("%p show\n",&e)
}

type FInterface interface {
	Test01()
	Test02()
}

type GInterface interface {
	Test01()
	Test03()
}

// MM 同时实现了FInterface，GInterface两个接口
type MM struct{

}

func (mm MM)Test01(){
	fmt.Println("Test01")
}

func (mm MM)Test02(){
	fmt.Println("Test02")
}

func (mm MM)Test03(){
	fmt.Println("Test03")
}

// FInterface、GInterface 同时包含Test01()方法，因此不能纳入桶一个接口
//type HInterface interface {
//	FInterface
//	GInterface
//}

type HInterface interface {
	Haha()
}

type GG struct {

}

func (g *GG)Haha(){
	fmt.Println("haha")
}

func test01(){
	g := GG{}
	// g 没有直接实现Haha()因此不能称为 HInterface的子类，不能直接赋值
	// 但通过指针发起对Haha的调用
	//var h HInterface = g
	g.Haha()


	var i HInterface = &g
	i.Haha()
}

// 手动实现sort.Sort(data interface{}) 中 Sort 相关接口函数 Len() int ,Less(i,j int) bool, Swap(i,j)，就可以对自定义结构体进行排序
type Kid struct {
	Name string
	Age int
}

type KidSlice []Kid

func (ss KidSlice)Less(i,j int) bool {
	return ss[i].Age < ss[j].Age
}

func (ss KidSlice)Swap(i,j int){
	ss[i],ss[j] = ss[j],ss[i]
}

func (ss KidSlice)Len()int {
	return len(ss)
}

func selfDefineSort(){
	ss := KidSlice{Kid{"tom",23},Kid{"jerry",21},Kid{"lucy",33}}
	fmt.Println(ss)
	sort.Sort(ss)
	fmt.Println(ss)
}

// 通过继承，子类可以快速获取父类的所有属性和绑定方法，继承更多强调的是复用性 is a
// 实现接口可以看做是对目标对象的功能增强，接口更多强调的是功能扩展 like a
// 在不破坏继承关系基础上，通过实现接口，完成功能扩展
// 实现接口可以看成是对继承的补充，接口比继承更为灵活，接口的价值在于设计，注重范式
// 接口在一定程度上实现代码解耦
type Monkey struct {
	Name string
	Age int
	Color string
}

func (m Monkey)climb(){
	fmt.Println("monkey likes climbing.")
}

type LittleMonkey struct {
	Monkey
}

type Flyable interface {
	fly()
}

type Swimable interface {
	swim()
}

func (m LittleMonkey)fly(){
	fmt.Printf("%s can fly\n",m.Name)
}

func (m LittleMonkey)swim(){
	fmt.Printf("%s can swim\n",m.Name)
}

func extendsAndImpl(){
	monkey := LittleMonkey{Monkey{"tom", 2, "yellow"}}
	monkey.climb()
	monkey.fly()
	monkey.swim()

	var flyer Flyable = monkey
	flyer.fly()

	var swimmer Swimable = monkey
	swimmer.swim()
}

func main() {
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}

	computer.Working(phone)
	computer.Working(camera)

	computer.Working2(phone)
	computer.Working2(camera)

	// 接口不能直接初始化，但可以指向其实现类的变量
	//var usbDevice Usb = Usb{}
	var usbDevice Usb = phone
	usbDevice.Start()

	var a integer = 1
	a.SayHi()

	var b AInterface = a // 父类引用指向子类实现
	b.SayHi()

	// 多实现
	m1 := Monster{}
	m1.SayHi()
	m1.Hello()

	var m2 AInterface = Monster{}
	m2.SayHi()

	var m3 BInterface = Monster{}
	m3.Hello()

	s := Student{"tom"}
	live(s)

	// 接口类型默认是指针（引用类型），没有初始化的接口引用，值为nil
	var d DInterface
	if d == nil {
		fmt.Println("d is not init()")
	}

	show(s)
	show(m1)
	show(1)

	// 任意变量都可以赋值给空接口引用
	var n interface{} = s
	n = m1
	fmt.Println(n)

	mm := MM{}
	var f FInterface = mm
	var g GInterface = mm
	f.Test01()
	g.Test03()

	test01()

	selfDefineSort()

	extendsAndImpl()

}
