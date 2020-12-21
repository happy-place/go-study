package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// 往通道发送和接收数据都是阻塞过程，是select - default 可以实现非阻塞机制
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}

}

/**
no message received
no message sent
no activity
 */