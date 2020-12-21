package model

import "fmt"

type person struct {
	Name string
	age  int
	sal  float64
}

// 通过工厂函数 获取私有化结构体的指针，从而间接获取结构体对象
func NewPerson(name string, age int, sal float64) *person {
	return &person{name, age, sal}
}

// 通过绑定工厂方法，实现对私有属性的getter/setter访问
func (p *person) SetAge(age int) {
	if age >0 && age < 150 {
		p.age = age
	}else{
		// 安全控制
		fmt.Println("年龄范围不正常")
	}

}

func (p *person) GetAge() int {
	return p.age
}

func (p *person)SetSal(sal float64){
	p.sal = sal
}

func (p *person)GetSal()float64{
	return p.sal
}
