package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main(){
	// 接收信号通道
	sigs := make(chan os.Signal,1)
	done := make(chan bool,1)

	// 注册接收到interrupt 或 terminated 信号，就将信号传入sigs
	signal.Notify(sigs,syscall.SIGINT,syscall.SIGTERM)

	go func(){
		sig := <- sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<- done
	fmt.Println("existing")

}
/**
awaiting signal
^C
interrupt
existing
 */