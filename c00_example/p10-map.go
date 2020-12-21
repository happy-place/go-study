package main

import "fmt"

func main(){
	// 声明 string -> int 类型 map
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:",m)

	// 取出指定元素，并检查map长度
	v1 := m["k1"]
	fmt.Println("v1:",v1)
	fmt.Println("len:",len(m))

	// 删除指定元素
	delete(m,"k2")
	fmt.Println("map:",m)

	// 指定key取value时，通过第二个参数，反馈指定key是否存在
	_,prs := m["k2"]
	fmt.Println("prs:",prs)

	// 声明map同时执行初始化操作
	n := map[string]int{"foo":1,"bar":2}
	fmt.Println("map:",n)

}
/**
map: map[k1:7 k2:13]
v1: 7
len: 2
map: map[k1:7]
prs: false
map: map[bar:2 foo:1]
 */
