package view

import (
	"fmt"
	s "go_study/c32_customer_management/service"
	"strconv"
)

type CustomerView struct {
	key string
	header string
	template string
	footer string
	loop bool
	service s.CustomerService
}

func NewCustomerView() *CustomerView{
	return &CustomerView{
		key: "",
		header: "\n\t\t\t\t\t客户关系管理系统",
		loop: true,
	}
}

func (this *CustomerView) list(){
	customers := this.service.Customers
	if len(customers) ==0 {
		fmt.Println("无客户信息，请添加！")
	}else{
		fmt.Println(this.header)
		fmt.Println("编号\t姓名\t年龄\t性别\t电话\t邮箱")
		fmt.Println("---------------------------------------------------------------------------")
		for _,target := range customers{
			fmt.Printf("%v\t\t%v\t\t%v\t\t%v\t\t%v\t\t%v\n",
				target.Id,target.Name,target.Age,target.Gender,target.Phone,target.Email)
		}
		fmt.Println("---------------------------------------------------------------------------")
	}
}

func (this *CustomerView) add(){
	fmt.Println(this.header)
	fmt.Println("-------------------------------添加客户--------------------------------")
	var name string
	fmt.Printf("请输入名称：")
	fmt.Scanf("%s",&name)

	var age int
	fmt.Printf("请输入年龄：")
	fmt.Scanf("%d",&age)

	var gender string
	fmt.Printf("请输入性别：")
	fmt.Scanf("%s",&gender)

	var phone string
	fmt.Printf("请输入电话：")
	fmt.Scanf("%s",&phone)

	var email string
	fmt.Printf("请输入邮箱：")
	fmt.Scanf("%s",&email)

	fmt.Println("----------------------------------------------------------------------")

	this.service.Add(name,age,gender,phone,email)
}

func (this *CustomerView) delete(){
	fmt.Println(this.header)
	fmt.Println("-------------------------------删除客户--------------------------------")
	var id int
	fmt.Printf("请输入删除客户编码（返回：-1）：")
	fmt.Scanf("%d",&id)
	fmt.Println("----------------------------------------------------------------------")
	if id != -1 {
		this.service.Delete(id)
	}
}

func (this *CustomerView) defaultOrInput(key string,value string) string {
	if key == "" {
		return value
	}else{
		return key
	}
}

func (this *CustomerView) update(){
	fmt.Println(this.header)
	fmt.Println("-------------------------------修改客户--------------------------------")
	var id int
	fmt.Printf("请输入修改客户编码：")
	fmt.Scanf("%d",&id)
	fmt.Println("----------------------------------------------------------------------")

	_,target := this.service.Get(id)

	if target != nil {
		var key string

		fmt.Printf("请输入名称（%s）：",target.Name)
		fmt.Scanf("%s",&key)
		name := this.defaultOrInput(key,target.Name)

		fmt.Printf("请输入年龄（%v）：",target.Age)
		fmt.Scanf("%s",&key)
		ageStr := this.defaultOrInput(key,fmt.Sprintf("%d",target.Age))
		age,_ := strconv.Atoi(ageStr)

		fmt.Printf("请输入性别（%s）：",target.Gender)
		fmt.Scanf("%s",&key)
		gender := this.defaultOrInput(key,target.Gender)

		fmt.Printf("请输入电话（%s）：",target.Phone)
		fmt.Scanf("%s",&key)
		phone := this.defaultOrInput(key,target.Phone)

		fmt.Printf("请输入邮箱（%s）：",target.Email)
		fmt.Scanf("%s",&key)
		email := this.defaultOrInput(key,target.Email)

		this.service.Update(id,name,age,gender,phone,email)
	} else {
		fmt.Print("不存在指定编号为 %v 客户！\n",id)
	}
}

func (this *CustomerView) exit(){
	fmt.Println(this.header)
	fmt.Println("-------------------------------退出系统--------------------------------")
	var signal string
	for{
		fmt.Printf("确认退出吗？（Y/N）：")
		fmt.Scanf("%s",&signal)
		if signal == "y" || signal == "Y" {
			this.loop = false
			break
		}else{
			fmt.Println("无效输入，请重新选择！")
		}
	}

	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("Goodby~")
}

func (this *CustomerView)Menu(){
	for;this.loop;{
		fmt.Println("\t\t\t\t客户关系管理系统")
		fmt.Println("----------------------------------------------------------------------")
		fmt.Println("\t\t\t\t【1】查看全部客户")
		fmt.Println("\t\t\t\t【2】新增客户信息")
		fmt.Println("\t\t\t\t【3】删除客户信息")
		fmt.Println("\t\t\t\t【4】修改客户信息")
		fmt.Println("\t\t\t\t【5】退出信息系统")
		fmt.Println("----------------------------------------------------------------------")
		fmt.Printf("请输入选项：")
		fmt.Scanf("%s",&this.key)
		switch this.key {
		case "1":
			this.list()
		case "2":
			this.add()
		case "3":
			this.delete()
		case "4":
			this.update()
		case "5":
			this.exit()
		default:
			fmt.Println("无效选项！")
		}
	}
}




