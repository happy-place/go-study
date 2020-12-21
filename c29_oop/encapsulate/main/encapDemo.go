package main

import (
	"fmt"
	"go_study/c29_oop/encapsulate/model"
)

/**
	封装性：通过工厂函数，和工厂方法实现对私有结构体，及私有结构体属性的访问
 */
func encapsulate1(){
	person := model.NewPerson("tom", 23, 5000)
	fmt.Printf("%+v\n",person)
	fmt.Println("name=",person.Name)
	person.SetAge(33)
	fmt.Println("age=",person.GetAge())
	person.SetSal(10000)
	fmt.Println("sal=",person.GetSal())
}

func encapsulate2(){
	acc := model.NewAcc("686886",100,"888888")
	if acc == nil {
		fmt.Println("NewAcc创建对象失败")
		return
	}
	acc.SetAccNo("6868")
	acc.SetAccNo("686868")
	fmt.Println(acc.GetAccNo())

	acc.SetBalance(10.0)
	acc.SetBalance(100)
	fmt.Println(acc.GetBalance())

	acc.SetPwd("666")
	acc.SetPwd("666666")
	fmt.Println(acc.GetPwd())

}

func main(){
	//encapsulate1()
	encapsulate2()
}