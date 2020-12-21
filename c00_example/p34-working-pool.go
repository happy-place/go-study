package main

import (
	"fmt"
	"time"
)

// 声明worker处理过程耗时按1秒模拟
func jobworker(id int,jobs <- chan int,results chan <- int){
	for j:= range jobs {
		fmt.Println("jobworker",id,"processing job",j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func main(){
	// 任务缓存通道
	jobs := make(chan int,100)
	// 结果缓冲通道
	results := make(chan int,100)

	// 启动三个worker协程
	for w:=1;w<=3;w++{
		go jobworker(w,jobs,results)
	}

	// 任务队列插入9个任务
	for j:=1;j<=9;j++{
		jobs <- j
	}
	// 如果不关闭通道，将worker协程将无法退出
	close(jobs)

	// 取出结果
	for a:=1;a<=9;a++{
		fmt.Println(<- results)
	}
}

/**
jobworker 3 processing job 1
jobworker 1 processing job 2
jobworker 2 processing job 3
jobworker 2 processing job 4
jobworker 1 processing job 5
6
4
2
jobworker 3 processing job 6
jobworker 1 processing job 8
jobworker 3 processing job 7
12
10
8
jobworker 2 processing job 9
18
14
16
 */
