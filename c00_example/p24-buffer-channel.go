package main

import "fmt"

func main(){
	// 创建允许缓冲2个元素的缓冲通道
	messages := make(chan string,2)
	messages <- "buffered"
	//messages <- "channel"
	fmt.Println(<- messages)
	//fmt.Println(<- messages) //只存一个元素时，如果取超了，抛异常 fatal error: all goroutines are asleep - deadlock!
	fmt.Println(len(messages), &messages, cap(messages))
	// cap(messages) 容量
	// len(messages) 当前元素个数
	// &messages 取地址
}
