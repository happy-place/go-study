package main

import (
	"fmt"
	"go_study/c27_factory/model"
)

func createObj(){
	var p1 = model.Person{"tom",21} // 创建结构体变量
	p2 := model.Person{"jack",28}
	fmt.Printf("obj:%v, ptr: %p\n",p1,&p1)
	fmt.Printf("obj:%v, ptr: %p\n",p2,&p2)

	var p3 = &model.Person{"tom",21} // 创建结构体指针
	p4 := &model.Person{"tom",21}
	fmt.Printf("obj:%v, ptr: %p\n",*p3,p3)
	fmt.Printf("obj:%v, ptr: %p\n",*p4,p4)
}

func createByFactory(){
	c1 := model.Car1{"jpk","yellow"}
	fmt.Printf("obj: %v, color: %s\n",c1,c1.Color)

	c2 := model.GetCar2("yyk","black")
	fmt.Printf("obj: %v, color: %s\n",*c2,c2.GetCar2Color())
}

func main(){
	//createObj()
	createByFactory()
}