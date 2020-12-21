package main

import (
	"fmt"
	"time"
)

// 协程中打点，主进程阻塞等待time.Millisecond * 1600，然后停止打点器
func main(){
	ticker := time.NewTicker(time.Millisecond * 500)
	go func(){
		for t := range ticker.C {
			fmt.Println("Tick at",t)
		}
	}()

	time.Sleep(time.Millisecond * 1600)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
