package main

import (
	"fmt"
	"math/rand"
	"sync"
	"runtime"
	"sync/atomic"
	"time"
)

func main(){
	var state = make(map[int]int)
	var mutex = &sync.Mutex{}

	var ops int64 = 0

	for r :=0;r<100;r++{
		go func(){
			total := 0
			for{
				key := rand.Intn(5)
				// 使用互斥锁确保对state独占访问
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddInt64(&ops,1)
				// 为了确保这个 Go 协程不会在调度中饿死，我们在每次操作后明确的使用 runtime.Gosched()进行释放。
				runtime.Gosched()
			}
		}()
	}

	for w:=0;w<10;w++{
		go func(){
			key := rand.Intn(5)
			val := rand.Intn(100)
			// 使用互斥锁确保对state独占访问
			mutex.Lock()
			state[key] = val
			mutex.Unlock()
			atomic.AddInt64(&ops,1)
			// 为了确保这个 Go 协程不会在调度中饿死，我们在每次操作后明确的使用 runtime.Gosched()进行释放。
			runtime.Gosched()
		}()
	}

	time.Sleep(time.Second)
	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops:",opsFinal)

	mutex.Lock()
	fmt.Println("state:",state)
	mutex.Unlock()
}

/**
ops: 4678417
state: map[0:20 1:95 3:8 4:84]
 */