package pojo

import "fmt"

type Customer struct {
	Id int
	Name string
	Age int
	Gender string
	Phone string
	Email string
}

func NewCustomer(id int,name string,age int,gender string,phone string,email string) *Customer {
	return &Customer{id,name,age,gender,phone,email}
}

func (this *Customer)Update(name string,age int,gender string,phone string,email string){
	this.Name = name
	this.Age = age
	this.Gender = gender
	this.Phone = phone
	this.Email = email
}

func (this *Customer)GetInfo()string{
	return fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v",this.Id,this.Name,this.Age,this.Gender,this.Phone,this.Email)
}
