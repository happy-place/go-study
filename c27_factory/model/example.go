package model

type Person struct {
	Name string
	Age int
}

type Car1 struct {
	Name string
	Color string
}

type car2 struct {
	name string
	color string
}

// 工厂函数解决结构体访问权限问题
func GetCar2(name string,color string) *car2{
	return &car2{name,color}
}

// 绑定方法，解决结构体属性访问权限问题
func (ptr *car2) GetCar2Color() string {
	return ptr.color
}