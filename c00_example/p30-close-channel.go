package main

import "fmt"

func main(){
	jobs := make(chan int,5)
	done := make(chan bool)

	go func(){
		for{
			// more 通道关闭为false，否则为true
			j,more := <- jobs
			if more {
				fmt.Println("received job",j)
			}else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j:=1;j<=3;j++{
		jobs <- j
		fmt.Println("sent job",j)
	}
	// 关闭通道后，不能继续往通道传数据了
	close(jobs)
	fmt.Println("sent all jobs")
	fmt.Println(<- done)
}

/**
sent job 1
sent job 2
sent job 3
sent all jobs
received job 1
received job 2
received job 3
received all jobs
true
 */
