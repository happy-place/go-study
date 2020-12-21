package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

func main(){
	var ops uint64 = 0

	for i:=0;i<50;i++{
		go func(){
			for{
				// 原子累加
				atomic.AddUint64(&ops,1)
				// 为了确保这个 Go 协程不会在调度中饿死，我们在每次操作后明确的使用 runtime.Gosched()进行释放
				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:",opsFinal)

}

/**
ops: 5751517
 */