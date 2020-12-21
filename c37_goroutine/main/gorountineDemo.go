package main

import (
	"fmt"
	"runtime"
	"time"
)

/**
	进程：操作系统分配资源最基本单位；
	线程：cpu调度基本单位

	单核跑多线程，就是并发 （来回切）
	多个任务跑在一个cpu上，每个时间点只有一个执行，一段时间内，所有任务都有执行

	多核跑多线程，就是并行 （同时执行）
	多个任务作用在多个cpu上，每个时间点，所有任务都在同时执行。

	协程：一个线程中可以起多个协程（约3万个），这是由go编译器，实现的。
	1）有独立栈空间
	2）共享程序堆空间
	3）调度由用户控制
	4）协程是轻量级线程。

	主线程是物理线程，直接作用在cpu上（线程是cpu调度最基本单位，进程是操作系统资源分配基本单位），是重量级的，非常消耗cpu子资源;
	协程是从主线程开启的轻量级线程，是逻辑态的线程（拥有独立栈，共享程序堆，由用户进行调度），资源消耗相对较少。
	Golang的协程机制是重要特点，可以轻松开启上万个协程，其他编程语言的开发机制都是基于线程，而Golang是基于协程，更省资源。

	c、c++、java 的线程是内核态的，一个核支撑几千个线程，很快就将硬件资源消耗完了，非常重
	golang的协程是，逻辑态的线程，可以理解为时用户空间维持的线程，一个上下文环境P，可以很轻松开启上万个协程，比较轻；

	MPG模型
	M - 物理态主线程
	P -	协程执行上下文环境
	G - 协程

	1.一个P（上下文环境）下面以队列形式管理了多个协程（即，可以理解为一个主线程可以轻松起上万个协程）；
	2.某个协程G3阻塞时，会自动让出cpu执行权限，通知P执行G3后面的协程，当G3取消阻塞时，又重新开始执行G3，通过上下文来回切换，保证最高效使用cpu。

 */

/**
	子线程（协程）每个1秒打印输出一次
	主线程每个2秒输出一次
	main-thread print 0
	test-thread print 0
	test-thread print 1
	test-thread print 2
	main-thread print 1
	test-thread print 3
	test-thread print 4
 */
func printNum(thread string,duration time.Duration){
	for i:=0;i<10;i++{
		fmt.Printf("%s print %v\n",thread,i)
		time.Sleep(time.Second * duration)
	}
}

func testPrint(){
	// 拷贝一个分支，用协程执行，主线程退出，协程也一起退出，但协程也可以优先于主线程退出
	go printNum("test-thread",1)

	printNum("main-thread",2)
}


func setCpu(){
	cpu := runtime.NumCPU() // cpu 逻辑核数
	runtime.GOMAXPROCS(cpu -1) // 设置最多使用核数
	fmt.Printf("逻辑cpu数:%v\n",cpu)
	// 默认让程序运行在所有cpu上，通过设置cpu可以更高效使用cpu
}

var (
	resMap = make(map[int]int,10)
)

func sum(n int){
	res := 1
	for i:=1;i<=n;i++{
		res *= i
	}
	resMap[n] = res
}

func everyNumSum(){
	for i:=1;i<=200;i++{
		go sum(i)
	}
	// fatal error: concurrent map writes
	// 多个协程同时往一个map写入，报并发写入异常
	time.Sleep(time.Second * 10)

	for n,res := range resMap {
		fmt.Printf("%d! = %v\n",n,res)
	}
}

func main(){
	//testPrint()
	//setCpu()
	everyNumSum()
}

