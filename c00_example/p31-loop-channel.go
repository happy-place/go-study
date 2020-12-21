package main

import "fmt"

// 创建缓冲通道后，连续输入两次，最后关闭通道，然后遍历通道

func main(){
	queue := make(chan string,2)
	queue <- "one"
	queue <- "two"
	close(queue)
	// 遍历时如果通道开始没有关闭，则会报异常 fatal error: all goroutines are asleep - deadlock!
	for elem := range queue {
		fmt.Println(elem)
	}
}
/**
one
two
 */
