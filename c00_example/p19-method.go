package main

import "fmt"

// 定义结构体
type rect struct {
	width,height int
}

// 将方法绑定到结构体
func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2 * r.width + 2 * r.height
}

func main(){
	r := rect{width:10,height: 5}
	fmt.Println("area:",r.area())
	fmt.Println("perim:",r.perim())

	rp := &r
	fmt.Println("area:",rp.area())
	fmt.Println("perim:",rp.perim())
}

/**
area: 50
perim: 30
area: 50
perim: 30
 */