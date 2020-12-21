package service

import (
	"go_study/c32_customer_management/pojo"
)
type CustomerService struct {
	space           string
	Customers       []*pojo.Customer
	autoIncrementId int
}

func NewCustomerService() *CustomerService{
	space := "\t\t"
	customers := make([]*pojo.Customer,5)
	return &CustomerService{space,customers,0}
}

func (this *CustomerService)Add(name string,age int,gender string,phone string,email string) {
	target := &pojo.Customer{this.autoIncrementId,name,age,gender,phone,email}
	this.Customers = append(this.Customers, target)
	this.autoIncrementId ++
}

func (this *CustomerService)Delete(id int) bool{
	index,target := this.Get(id)
	if target!=nil {
		this.Customers = append(this.Customers[:index],this.Customers[index+1:]...)
		return true
	}else{
		return false
	}
}

func (this *CustomerService)Get(id int) (int,*pojo.Customer) {
	for index,target := range this.Customers {
		if target.Id == id {
			return index,target
		}
	}
	return -1,nil
}

func (this *CustomerService)Update(id int,name string,age int,gender string,phone string,email string) {
	_,target := this.Get(id)
	target.Update(name ,age ,gender ,phone ,email)
}


