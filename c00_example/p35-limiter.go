package main

import (
	"fmt"
	"time"
)

func main() {
	//  创建缓冲通道，一次性把数据存进去，然后按打点器节奏消费
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(time.Millisecond * 200)
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// 创建速率控制器，前三个元素几乎同时处理，后面的按time.Tick(time.Millisecond * 200)节奏处理
	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	// 遍历burstyRequests前，需要先关闭burstyRequests
	close(burstyRequests)

	// 前3个消费几乎同时进行，之后的按time.Millisecond * 200节奏消费
	for req := range burstyRequests {
		<- burstyLimiter
		fmt.Println("request",req,time.Now())
	}

}

/**
request 1 2020-11-06 08:18:42.05256 +0800 CST m=+0.203184268
request 2 2020-11-06 08:18:42.254631 +0800 CST m=+0.405253552
request 3 2020-11-06 08:18:42.449631 +0800 CST m=+0.600251472
request 4 2020-11-06 08:18:42.649656 +0800 CST m=+0.800275457
request 5 2020-11-06 08:18:42.851619 +0800 CST m=+1.002237465

request 1 2020-11-06 08:18:42.851686 +0800 CST m=+1.002303953
request 2 2020-11-06 08:18:42.851702 +0800 CST m=+1.002319790
request 3 2020-11-06 08:18:42.851712 +0800 CST m=+1.002329995

request 4 2020-11-06 08:18:43.052437 +0800 CST m=+1.203053531
request 5 2020-11-06 08:18:43.254599 +0800 CST m=+1.405214742
 */
