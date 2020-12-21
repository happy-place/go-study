package main

import (
	"fmt"
)

type Usb interface {
	plugin()
}

func update(usb Usb){
	usb.plugin()
	fmt.Println("USB外设更新驱动")
}

type Phone struct {

}

func (p Phone)call(){
	fmt.Println("手机打电话")
}

func (p Phone)plugin(){
	fmt.Println("手机下载驱动")
}

type Camera struct {

}

func (c Camera)photo(){
	fmt.Println("相机照相")
}

func (c Camera)plugin(){
	fmt.Println("相机下载驱动")
}

func main(){
	phone := Phone{}
	phone.call()
	phone.plugin() // 多态3：相同抽象方法，不同具体实现

	update(phone) // 多态2：声明接收接口，传入接口实现类子类
	var usb1 Usb = phone // 多态1：父类引用指向子类实例
	usb1.plugin()

	camera := Camera{}
	camera.photo()
	camera.plugin()
	update(camera)

}


