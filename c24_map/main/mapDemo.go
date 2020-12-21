package main

import (
	"fmt"
	"sort"
)

/**
 kv结构
 1）key 可以是任意类型，比如: 数值类型，bool，string,指针,chan ，通常为 int 或 string，但不能是slic、map、function，因为不能进行'=='运算
 2）value 可以是任意类型；可以是 map string struct
 3）必须使用make() 申请内存完成初始化，不指定长度的话，默认长度为1
 4）维持的顺序是key的顺序，不是插入顺序

 */

func create(){
	// 显示声明需要的容量，后期依旧可以递增
	var m1 = make(map[string]int,2)
	m1["a"] = 1
	m1["b"] = 2
	m1["c"] = 11 // 覆盖
	fmt.Println(m1)

	// 不声明容量
	var m2 = make(map[string]int)
	m2["a"] = 1
	m2["b"] = 2
	m2["c"] = 3
	fmt.Println(m2)

	// 创建同时进行赋值
	var m3 = map[string]int{
		"a":1,
		"b":2,
	}
	fmt.Println(m3)

}

/**
	嵌套map
 */
func test1(){
	var students = make(map[string]map[string]string,2)

	students["stu01"] = make(map[string]string,3)
	students["stu01"]["name"] = "tom"
	students["stu01"]["sex"] = "男"
	students["stu01"]["address"] = "北京长安街"

	students["stu02"] = make(map[string]string,3)
	students["stu02"]["name"] = "jack"
	students["stu02"]["sex"] = "女"
	students["stu02"]["address"] = "上海黄浦江"

	fmt.Println(students)
}

func crud(){
	var cities  = make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "上海"
	cities["no3"] = "广州"

	// 存在就覆盖，否则就插入
	cities["no3"] = "深圳"

	 // 存在就删除，否则不执行任何操作
	delete(cities,"no4")
	delete(cities,"no3")
	fmt.Println(cities)

	// 回收map空间
	// 方案1：遍历删除所有key
	for k := range cities{
		delete(cities,k)
	}
	// 方案2：指向新map，原来map失去引用，自动gc
	cities  = make(map[string]string)
}

// map 查找
func existed(){
	var m1 = make(map[string]string)
	m1["no1"] = "tom"
	m1["no2"] = "jack"
	v,isExisted := m1["no3"]
	if isExisted {
		fmt.Println(v)
	}else{
		fmt.Println("not existed")
	}
}

func loop(){
	m1 := make(map[string]string)
	m1["no1"] = "a"
	m1["no2"] = "b"
	m1["no3"] = "c"
	m1["no4"] = "d"

	// 只接收key
	for k := range m1 {
		fmt.Printf("%s -> %s\n",k,m1[k])
	}

	// 同时接收key 和 value
	for k,v := range m1 {
		fmt.Printf("%s -> %s\n",k,v)
	}

	var students = make(map[string]map[string]string,2)

	students["stu01"] = make(map[string]string,3)
	students["stu01"]["name"] = "tom"
	students["stu01"]["sex"] = "男"
	students["stu01"]["address"] = "北京长安街"

	students["stu02"] = make(map[string]string,3)
	students["stu02"]["name"] = "jack"
	students["stu02"]["sex"] = "女"
	students["stu02"]["address"] = "上海黄浦江"

	for k,v := range students {
		fmt.Println("key=",k)
		for kk,vv := range v {
			fmt.Printf("%s=%v,",kk,vv)
		}
		fmt.Println()
	}

	fmt.Println(len(students)) // k-v 个数

}

func mapAppend(){
	s1 := make([]map[string]string,2)
	if s1[0] == nil {
		s1[0] = map[string]string{
			"no01":"a",
			"no02":"b",
			"no03":"c",
		}
	}

	if s1[1] == nil {
		s1[1] = map[string]string{
			"no03":"A",
			"no04":"b",
		}
	}

	m3 := map[string]string {
		"no05":"d",
		"no06":"e",
	}

	// index out of range [2] with length 2 直接增加长度超限
	//if s1[2] == nil {
	//	s1[2] = m3
	//}

	// 使用切片增加，可以
	s1 = append(s1,m3)
	fmt.Println(s1)

}

/**
	map 遍历默认是无序的，可以将key收集到切片，统一排序，然后基于排好序的切片进行遍历
 */
func mapSort(){
	m1 := make(map[int]int,10)
	m1[1] = 13
	m1[10] = 100
	m1[4] = 56
	m1[8] = 90

	fmt.Println(m1) // map[1:13 4:56 8:90 10:100]

	var keys []int
	for k,_ := range m1 {
		keys = append(keys,k)
	}
	sort.Ints(keys)
	fmt.Println(keys)

	for _,k := range keys {
		fmt.Printf("%v：%v\n",k,m1[k])
	}
}

func test2(m map[string]string){
	m["no1"] = "aa"
}

/**
 map 是引用类型，作为参数传递时，方法内对map的修改会影响全局
 */
func referPass(){
	m := make(map[string]string)
	m["no1"] = "ab"
	m["no2"] = "ab"
	test2(m)
	fmt.Println(m)
}

/**
	map 自动扩容了
 */
func dynamicExpand(){
	m1 := make(map[string]int,2)
	m1["a"] = 1
	m1["b"] = 2
	m1["c"] = 3
	fmt.Println(m1)
}

type student struct {
	Name string
	Age int
	Address string
}


func test4(m map[string]student){
	s := m["no02"]
	s.Address = "wuhan"
	fmt.Println(m)
}

/**
	通常使用结构体管理复杂对象
 */
func test3(){
	students := make(map[string]student,10)
	stu1 := student{Name: "tom",Age: 23,Address: "beijing"}
	stu2 := student{Name: "jack",Age: 23,Address: "beijing"}
	students["no01"] = stu1
	students["no02"] = stu2
	test4(students)
	fmt.Println(students)
}

func modifyUser(users map[string]map[string]string,name string){
	if users[name] != nil {
		users[name]["pwd"] = "888888"
	}else{
		users[name] = make(map[string]string,2)
		users[name]["pwd"] = "888888"
		users[name]["nickname"] = "昵称~" + name
	}
}
func test5(){
	users := make(map[string]map[string]string,10)
	modifyUser(users,"tom")
	modifyUser(users,"mary")
	users["smith"] = make(map[string]string)
	users["smith"]["pwd"] = "99999"
	users["smith"]["nickname"] = "Mr Smith"
	modifyUser(users,"smith")
	fmt.Println(users)
}


func main(){
	//create()
	//test1()
	//crud()
	//existed()
	//loop()
	//mapAppend()
	//mapSort()
	//referPass()
	//dynamicExpand()
	//test3()
	test5()
}
