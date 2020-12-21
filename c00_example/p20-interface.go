package main

import "fmt"
import "math"

// 声明接口: 凡是能调用 area() 和 perim() 方法的对象都可以看成是geometry对象
type geometry interface {
	// 接口方法
	area() float64
	perim() float64
}

// 声明结构体
type rectangle struct {
	width,height float64
}

// 声明结构体
type circle struct {
	radius float64
}

// 结构体绑定方法
func (r rectangle) area() float64 {
	return r.height * r.width
}

func (r rectangle) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64{
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64{
	return math.Pi * c.radius * 2
}


func measure(g geometry){
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main(){
	r := rectangle{width: 3,height: 4}
	c := circle{radius: 5}
	measure(r)
	measure(c)
}

/**
{3 4}
12
14
{5}
78.53981633974483
31.41592653589793
 */