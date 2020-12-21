package main

import (
"context"
"log"
"os"
"time"
)

var logger *log.Logger

func someHandler() {
	ctx, cancel := context.WithCancel(context.Background())
	go doStuff(ctx)

	//10秒后取消doStuff
	time.Sleep(10 * time.Second)
	cancel()
	time.Sleep(2 * time.Second)
}

//每1秒work一下，同时会判断ctx是否被取消了，如果是就退出
func doStuff(ctx context.Context) {
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done(): // ctx同时产生是的cancel 被调用时，就停止阻塞
			logger.Printf("done")
			return
		default:
			logger.Printf("work")
		}
	}
}

func main() {
	logger = log.New(os.Stdout, "", log.Ltime)
	someHandler()
	logger.Printf("down")
}

/**
17:46:47 work
17:46:48 work
17:46:49 work
17:46:50 work
17:46:51 work
17:46:52 work
17:46:53 work
17:46:54 work
17:46:55 work
17:46:56 down
*/