package main

import (
	"fmt"
	"errors"
)

// 内置函数
/**
new() 用来分配内存，主要分配值类型(int float struct)，输入需要分配的类型，返回指定类型的指针
 */
func newDemo(){
	num1 := 100
	fmt.Printf("类型: %T， 值: %v，地址:%v\n",num1,num1,&num1)
	// 类型: int， 值: 100，地址:0xc00001e090

	num2 := new(int) // 指针
	num2 = &num1
	fmt.Printf("类型: %T，存储值: %v, 地址映射空间的值: %v，地址:%v\n",num2,num2,*num2,&num2)
	// 类型: *int，存储值: 0xc00001e090, 值: 100，地址:0xc00000e030
}

/**
make 分配内存，主要分配引用类型，比如 chan map slice。
 */
func makeDemo(){
	// 1.无缓冲的通道(无缓冲时，只有出一个才能进一个)
	c1 := make(chan int)
	go func(){
		c1 <- 1 // 发送元素
		close(c1) // 必须close() 否则循环消费时会报错
	}()
	fmt.Println(<-c1) // 接收元素，如果没有，就一直阻塞
	// 1

	// 2.三个长度缓存通道(有缓冲时，缓冲数以内，可以一直发送，直到通道打满)
	c2 := make(chan int,3)
	go func(){
		c2 <-1
		c2 <-2
		c2 <-3
		close(c2)
	}()
	for e := range c2 {
		fmt.Println(e)
	}
	//1
	//2
	//3

	// 3.哈希结构
	m1 := make(map[string]int)
	m1["a"] = 1
	m1["b"] = 2
	for k,v := range m1 {
		fmt.Printf("%v -> %v\n",k,v)
	}
	//a -> 1
	//b -> 2

	fmt.Printf("c=%v, a=%v\n",m1["c"],m1["a"]) // c=0, a=1
	delete(m1,"a") // 删除
	fmt.Printf("c=%v, a=%v\n",m1["c"],m1["a"]) // c=0, a=0

	s1 := make([]int,2)
	m1["a"] = 1
	m1["b"] = 2
	m1["c"] = 3
	for i,e := range s1 {
		fmt.Printf("s1[%v] -> %v\n",i,e)
	}
	//s1[0] -> 0
	//s1[1] -> 0

}

/**
golang 不支持try-catch-finally 异常处理机制，基于defer-recover()-panic()兼容异常，使程序更健壮
1）使用defer声明需要延迟处理的表达式
2) 内置函数recover()捕获函数作用域内异常
3）内置函数panic()相当于 throw new Exception(error)，panic 可以接受interface{}即任意类型值作为参数，输出异常，并退出程序
 */
func errorDemo(){
	defer func(){ // 延迟处理机制
		err := recover() // recover 可以捕获函数作用域异常
		if err != nil {
			panic(err) // 想当与throw new Exception(error)
			fmt.Println("err=",err)
		}
	}()
	num1 := 1
	num2 := 0
	res := num1 / num2 // panic: runtime error: integer divide by zero

	fmt.Println(res)
}

func throwErr(a int,b int)(int,error){
	if b == 2 {
		return -1,errors.New("b = 2") // 手动声明异常
	}
	return a / b,nil
}

func main(){
	//newDemo()
	//makeDemo()
	//errorDemo()
	//fmt.Println("errorDemo over~")

	res,err := throwErr(1,0)
	if err != nil {
		panic(err) // 抛出异常，之后代码不会执行了
	}
	fmt.Println("res=",res)

}

