package main

import "fmt"
import "time"
func main(){
	// 声明一个传入string类型数据的通道
	message := make(chan string)
	// 在协程中往通道发送"ping"
	go func(){
		time.Sleep(10e7)
		message <- "ping"
	}()

	// 主线程中从通道里面取出，如果一直没有发送就一直阻塞
	msg := <- message
	fmt.Println(msg)
}
