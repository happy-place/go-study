package main

import (
"context"
"log"
"os"
"time"
)

var logger *log.Logger

func timeoutHandler() {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	go doTimeOutStuff(ctx)

	time.Sleep(10 * time.Second)

	cancel()

	time.Sleep(2 * time.Second)

}

func doTimeOutStuff(ctx context.Context) {
	for {
		time.Sleep(1 * time.Second)

		select {
		case <-ctx.Done():
			// ctx 同时创建的cancel 被调用 或 超时时，停止阻塞
			if deadline, ok := ctx.Deadline(); ok { // 判断是否设置了超时
				logger.Printf("deadline set")
				if time.Now().After(deadline) { // 如果超时，输出超时异常，并退出
					logger.Printf(ctx.Err().Error())
					return
				}
			}
			logger.Printf("done") // 未超时，调用cancel 结束
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

