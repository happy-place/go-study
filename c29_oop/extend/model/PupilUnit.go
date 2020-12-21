package model

import (
	"fmt"
)

/**
	继承：使编程更加贴近人类思维
	1）抽取公共字段放入公共父类；
	2）子类通过嵌入匿名结构体形式引入父类，实现继承效果，从而可以直接访问父类属性，以及父类绑定的方法；
	3）子类可以复写父类绑定的方法，实现自己的逻辑

	继承提高了代码复用性
	扩展性和维护性：
		1）对父类的扩展，子类自动感知；
		2）对子类可以单独扩展；
 */

type Student struct {
	Name string
	Age int
	Score int
}

func (p *Student) ShowInfo(){
	fmt.Printf("学生名称=%v 年龄=%v 成绩=%v\n",p.Name,p.Age,p.Score)
}

func (p *Student)SetScore(score int){
	p.Score = score
}

func (p *Student)Testing(){
	fmt.Printf("学员%v正在考试\n",p.Name)
}

// 通过嵌套结构体，实现继承关系，不光能访问到父类属性，还可以访问父类绑定的方法
type Pupil struct {
	Student
}

func (p *Pupil)Testing(){
	fmt.Printf("小学生%v正在考试\n",p.Name)
}

type Granduate struct {
	Student
}

func (p *Granduate)Testing(){
	fmt.Printf("大学生%v正在考试\n",p.Name)
}