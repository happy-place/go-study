package main

import (
"context"
"log"
"os"
"time"
)

var logger *log.Logger

func timeoutHandler() {
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(15*time.Second))
	// go doTimeOutStuff(ctx)
	go doStuff(ctx)

	time.Sleep(10 * time.Second)

	cancel()

	time.Sleep(2 * time.Second)

}

//每1秒work一下，同时会判断ctx是否被取消了，如果是就退出
func doStuff(ctx context.Context) {
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done(): // ctx同时产生的cancel 被调用 或 超时时，就停止阻塞
			logger.Printf("done")
			return
		default:
			logger.Printf("work")
		}
	}
}

func main() {
	logger = log.New(os.Stdout, "", log.Ltime)
	timeoutHandler()
	logger.Printf("end")
}
