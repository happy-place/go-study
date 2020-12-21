package main

import (
	"fmt"
	"time"
)

func main(){
	// 启动定时器1
	timer1 := time.NewTimer(time.Second * 2)

	// 运行到此处，开始阻塞time.Second * 2
	<- timer1.C
	fmt.Println("Timer 1 expired")

	// 启动定时器2
	timer2 := time.NewTimer(time.Second)

	// 协程中 timer2发起阻塞
	go func(){
		<- timer2.C
		fmt.Println("Timer 2 expired")
	}()

	// 定时器2在阻塞停止前，已经终止运行（取消）
	stop2 := timer2.Stop()
	if stop2{
		fmt.Println("Timer 2 stopped")
	}

}
