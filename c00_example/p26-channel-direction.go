package main

import "fmt"

// pings chan<- string 只发送通道（只存）
func ping(pings chan<- string, msg string){
	pings <- msg
}

// pings <-chan string 只接收通道（只取）
// pongs chan<- string 只发送通道 （只存）
func pong(pings <-chan string,pongs chan<- string){
	msg := <- pings
	pongs <- msg
}

func main(){
	pings := make(chan string,1)
	pongs := make(chan string,1)
	ping(pings,"passed message")
	pong(pings,pongs)
	fmt.Println(<-pongs)
}
