package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

/**
在前面的例子中，我们用互斥锁进行了明确的锁定来让共享的state 跨多个 Go 协程同步访问。
另一个选择是使用内置的 Go协程和通道的的同步特性来达到同样的效果。这个基于通道的方法和
Go 通过通信以及 每个 Go 协程间通过通讯来共享内存，确保每块数据有单独的 Go 协程所有的思路是一致的。
 */
type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {
	// 计数器，统计累计读/写操作次数
	var ops int64

	// 读通道
	reads := make(chan *readOp)
	// 写通道
	writes := make(chan *writeOp)

	// 协程1：接收其他协程的读写请求
	go func() {
		// 读写通道交换数据，中间寄存器
		var state = make(map[int]int)
		for {
			select {
			// 从读通道reads读取读请求，从state提取数据，然后吐出
			case read := <-reads:
				read.resp <- state[read.key]
			// 从写通道writes读取写请求，写入state，并返回响应true
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// 创建100个读协程，从reads通道中读取
	for r:=0;r<100;r++{
		go func(){
			for{
				// 创建&readOp
				read := &readOp{
					key: rand.Intn(5),
					resp: make(chan int),
				}
				// 往读通道，发送读请求
				reads <- read
				// 读取响应
				<- read.resp
				// 累计
				atomic.AddInt64(&ops,1)
			}
		}()
	}

	// 创建10个写协程，从writes中写出
	for w:=0;w<10;w++{
		go func(){
			for {
				// 封装&writeOp
				write := &writeOp{
					key: rand.Intn(5),
					val: rand.Intn(100),
					resp: make(chan bool)}
				// 往写通道，发送写请求
				writes <- write
				// 读取写入成功响应
				<- write.resp
				// 累计写次数
				atomic.AddInt64(&ops,1)
			}
		}()
	}

	// 主线程阻塞1秒
	time.Sleep(time.Second)

	// 读取累加器
	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops:",opsFinal)
}
