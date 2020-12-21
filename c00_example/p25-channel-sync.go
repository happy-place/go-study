package main

import (
	"fmt"
	"time"
)

// 接收一个bool类型的channel，调用中，先休眠1秒，然后通过通道，发送true
func worker(done chan bool){
	fmt.Println("working ...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

func main(){
	done := make(chan bool,1)
	// 协程生产数据
	go worker(done)
	// 主线程消费数据
	fmt.Println(<- done)
}

/**
working ...
done
true
 */