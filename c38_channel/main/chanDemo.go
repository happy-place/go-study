package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"
)

/**
不使用chan情况写，多线程写入存在并发操作问题 fatal error: concurrent map writes
通过  go build -race xxx.go 编译，./xxx 运行，最后会统计出哪些写入操作存在竞态情形，需要使用锁来解决
152! = 0
156! = 0  <<< 越界导致，int 存不下
177! = 0
Found 3 data race(s) <<<

多线程同步问题解决思路：
全局变量 + 锁
	抢到锁的线程执行，其余阻塞，避免多个线程同时操作共享资源，导致并发写入异常，主线程读取时也需要通过加锁，获取阻塞，避免协程未处理完，主线程提取退出

chan 管道思路
1）channel 本质上是一个数据结构，即 队列，遵循FIFO 先进先出，数据使用原则；
2）channel 是线程安全的，多个gorountine访问时，并不需要加锁，channel底层自己封装了锁的操作；
3) channel 自己是有类型的，一个string类型channel 本质上只能存储string类型数据

chan 是引用类型
chan 是有类型的
chan 必须make才能使用

go程序运行时，编译器其实在不停检查上下文，发现是无意义的阻塞，就直接报死锁
chan 的deadLock机制，编译器执行时，发现生产块，消费慢，会进入阻塞机制，但如果是只有生产， 没有消费，一旦容量满，就直接报死锁异常。

*/
var (
	resMap = make(map[int]int, 10)
	lock   = sync.Mutex{} // 互斥锁
	num    = 20
)

// 不加锁
func sum1(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	resMap[n] = res
}

func everyNumSum1() {
	for i := 1; i <= num; i++ {
		go sum1(i)
	}
	// fatal error: concurrent map writes
	// 多个协程同时往一个map写入，报并发写入异常
	time.Sleep(time.Second * 5)

	for n, res := range resMap {
		fmt.Printf("%d! = %v\n", n, res)
	}
}

// 上锁思路：多个协程写也加锁，主线程读也加锁
func sum2(n int) {
	lock.Lock()
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	resMap[n] = res
	lock.Unlock()
}

func everyNumSum2() {
	for i := 1; i <= num; i++ {
		go sum2(i)
	}
	// fatal error: concurrent map writes
	// 多个协程同时往一个map写入，报并发写入异常
	time.Sleep(time.Second * 5) // 这里的等待时间是估算，下面加锁读取，仍不能避免，协程没处理完，就开始读问题
	lock.Lock()                 // 10秒如果协程可能未执行完毕，此处拿锁是为避免这种情况，实现阻塞机制，保证下面代码执行前，协程全部执行完毕
	for n, res := range resMap {
		fmt.Printf("%d! = %v\n", n, res)
	}
	lock.Unlock()
}

type Person struct {
	Name    string
	Age     int
	Address string
}

func createChan() {
	var intChan chan int = make(chan int, 3)

	// 对象：0xc0000ba000,指针：0xc0000ae018
	// 管道本质上是执行内存队列数据结构的一个指针（引用类型）
	fmt.Printf("对象：%v,指针：%p\n", intChan, &intChan)

	// 存储数据个数不能超过cap，否则报fatal error: all goroutines are asleep - deadlock!
	intChan <- 1
	intChan <- 2
	fmt.Printf("已经存在元素个数:%v，总容量: %v\n", len(intChan), cap(intChan))

	intChan <- 3
	//intChan <- 4 // fatal error: all goroutines are asleep - deadlock! 》》 goroutine 1 [chan send] 存数据

	n := <-intChan // 从管道取出数据
	fmt.Println(n)

	n = <-intChan
	fmt.Println(n)
	n = <-intChan
	fmt.Println(n)
	//n = <- intChan // 取完了 ，报错 fatal error: all goroutines are asleep - deadlock! 》》 goroutine 1 [chan receive]: 取数据
	//fmt.Println(n)

}

func create2() {
	// chan 中存入map，需要对map进行初始化
	var mapChan = make(chan map[string]int, 3)
	m := make(map[string]int, 2)
	m["a"] = 1
	m["b"] = 2
	mapChan <- m

	mm := <-mapChan
	fmt.Printf("%v %T", mm, mm)

	// chan 中存入 Person
	personChan := make(chan Person, 3)
	p := Person{"tom", 21, "Hubei"}
	personChan <- p

	pp := <-personChan
	fmt.Println(pp)

	chan1 := make(chan interface{}, 3)
	chan1 <- 1
	chan1 <- "string"
	chan1 <- p

	<-chan1
	<-chan1
	ppp := <-chan1
	// 无法直接使用 ppp.Name，需要使用类型断言，转为Person
	fmt.Printf("%v %T\n", ppp, ppp)
	person := ppp.(Person)
	fmt.Printf("Name:%s\n", person.Name)
}

func test1() {
	personChan := make(chan Person, 10)
	go func() {
		for i := 0; i < 10; i++ {
			personChan <- Person{fmt.Sprintf("tom-%v", i), 20 + i, fmt.Sprintf("Hubei-%v", i)}
		}
		// 写入完毕，需要关闭chan,否则取数据可能会超取，报错
		// 关闭后，写入方无法继续写入，读出方，读取最后一个元素后，就不在阻塞，直接退出
		close(personChan)
	}()

	// 不能使用普通for循环（基于下标遍历），遍历管道，会报错 fatal error: all goroutines are asleep - deadlock! 》》 goroutine 1 [chan receive]: 取数据

	// for-range 遍历，自动感知管道有没有关闭
	// 未关闭情况：有数据，就取出，无数据就阻塞，等待写入
	// 关闭情况：有数据，就一直取，直到无数据，就退出循环
	// for-range 遍历，等效于下面的for循环
	for p := range personChan {
		fmt.Println(p)
	}

	//for {
	//	p,isOpen := <-personChan // 未关闭时就可以一直取，如果没有就阻塞，等待存入，如果关闭就退出循环
	//	if isOpen {
	//		fmt.Println(p)
	//	}else{
	//		break
	//	}
	//}
}

func writeAndRead() {
	intChan := make(chan int, 3) // 带缓冲管道，可以憋3个
	exitChan := make(chan bool)  // 不带缓冲管道，只能憋一个
	go func() {
		for i := 0; i < 50; i++ {
			intChan <- i
			log.Printf("写入：%v\n", i)
			//time.Sleep(time.Second *2)
		}
		close(intChan)
	}()

	go func() {
		for num := range intChan {
			fmt.Println(num)
			log.Printf("读出：%v\n", num)
		}
		exitChan <- true
		close(exitChan)
	}()

	exit := <-exitChan // 阻塞，直到读完
	fmt.Printf("is over %v\n", exit)
}

/**
一个协程生产
8个协程消费，且同时写出结果到另一个管道
主线程输出结果
g[1] -> g[8] -> main[1]
*/
func test2() {
	numChan := make(chan int, 3)
	resChan := make(chan []int, 3)
	limit := 2000
	go func() {
		for i := 1; i <= limit; i++ {
			numChan <- i
		}
		close(numChan)
	}()

	for i := 0; i < 8; i++ {
		go func() {
			for {
				num, isOpen := <-numChan
				if isOpen {
					res := 0
					for i := 1; i <= num; i++ {
						res += i
					}
					resChan <- []int{num, res}
				} else {
					break
				}
			}
		}()
	}

	count := 0
	for res := range resChan {
		fmt.Printf("res[%d] = %d\n", res[0], res[1])
		count++
		if count == limit {
			close(resChan)
			break
		}
	}
}

func test3() {
	numChan := make(chan int, 3)
	resChan := make(chan []int, 3)
	exitChan := make(chan bool, 4)
	limit := 2000
	go func() {
		for i := 1; i <= limit; i++ {
			numChan <- i
		}
		close(numChan)
	}()

	consumerCount := 8
	for i := 0; i < consumerCount; i++ {
		go func() {
			for {
				num, isOpen := <-numChan
				if isOpen {
					res := 0
					for i := 1; i <= num; i++ {
						res += i
					}
					resChan <- []int{num, res}
				} else {
					exitChan <- true
					break
				}
			}
		}()
	}

	go func() {
		for res := range resChan {
			fmt.Printf("res[%d] = %d\n", res[0], res[1])
		}
	}()

	for i := 0; i < consumerCount; i++ {
		<-exitChan
	}
	close(resChan)
	close(exitChan)

	fmt.Println("over~")

}

func isPrime(num int) bool {
	if num == 1 || num == 2 || num == 3 {
		return true
	}
	for i := 4; i <= num/2+1; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// 使用协程查找素数相比于普通查找，6 vs 27，效率提升4倍以上，cpu使用率接近100%
func test4() {
	startTmstp := time.Now().Unix()
	//cpu := runtime.NumCPU()
	//runtime.GOMAXPROCS(cpu - 1)
	nums := 300000
	checkers := 8

	numChan := make(chan int, 8)
	primeChan := make(chan int, 3)

	exitChan := make(chan bool, checkers)

	// 发数据
	go func() {
		for i := 1; i <= nums; i++ {
			numChan <- i
		}
		close(numChan)
	}()

	// 取数据，然后校验是否是素数，是的话存入primeChan
	for i := 0; i < checkers; i++ {
		go func() {
			for num := range numChan {
				if isPrime(num) {
					primeChan <- num
				}
			}
			exitChan <- true // 无数据可校验就退出，并发送退出信号
		}()
	}

	// 输出素数
	go func() {
		//for prime := range primeChan{
		//	fmt.Println(prime)
		//}

		for {
			_, isOpen := <-primeChan
			if !isOpen {
				break
			}
		}

	}()

	// 检验查找素数的8个协程是否全部工作完毕，如果是，就关闭primeChan，从而触发打印素数的协程关闭
	for i := 0; i < checkers; i++ {
		<-exitChan
	}

	close(primeChan)
	close(exitChan)
	stopTmstp := time.Now().Unix()
	fmt.Printf("total cost: %d,%d,%d\n", stopTmstp, startTmstp, stopTmstp-startTmstp)
	// total cost: 1606088907,1606088901,6
}

func test44() {
	startTmstp := time.Now().Unix()
	//cpu := runtime.NumCPU()
	//runtime.GOMAXPROCS(cpu - 1)
	nums := 300000
	primes := []int{}
	for i := 1; i <= nums; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	for i := 0; i < len(primes); i++ {
		//fmt.Println(primes[i])
	}
	stopTmstp := time.Now().Unix()
	fmt.Printf("total cost: %d,%d,%d\n", stopTmstp, startTmstp, stopTmstp-startTmstp)
	// total cost: 1606089233,1606089206,27
}

/**
一个协程随机生成1000个数，追加写入文件
另一个协程，在上面文件写入完毕时，读取文件中随机数，排序输出到另一个文件
*/
func test5() {
	signalChan := make(chan bool, 1)
	exitChan := make(chan bool, 1)

	wd, _ := os.Getwd()
	randFile := fmt.Sprintf("%s/src/go_study/c38_channel/resources/rand.txt", wd)
	sortFile := fmt.Sprintf("%s/src/go_study/c38_channel/resources/sort.txt", wd)

	go func() {
		nums := [2000]int{}
		for i := 0; i < 2000; i++ {
			rand.Seed(time.Now().UnixNano())
			nums[i] = rand.Intn(2000)
		}
		numBytes, _ := json.Marshal(nums)
		ioutil.WriteFile(randFile, numBytes, 0666)
		signalChan <- true
	}()

	go func() {
		<-signalChan
		close(signalChan)
		var nums [2000]int
		numBytes, _ := ioutil.ReadFile(randFile)
		json.Unmarshal(numBytes, &nums)

		for i := 0; i < len(nums); i++ {
			for j := 0; j < len(nums)-1-i; j++ {
				if nums[j] > nums[j+1] {
					nums[j], nums[j+1] = nums[j+1], nums[j]
				}
			}
		}
		sortBytes, _ := json.Marshal(nums)
		ioutil.WriteFile(sortFile, sortBytes, 0666)
		exitChan <- true
	}()

	<-exitChan
	close(exitChan)
	fmt.Println("over~")
}

func test6() {
	signalChan := make(chan bool, 10)
	exitChan := make(chan bool, 10)

	wd, _ := os.Getwd()

	for i := 0; i < 10; i++ {
		go func(i int) {
			randFile := fmt.Sprintf("%s/src/go_study/c38_channel/resources/rand-%d.txt", wd, i)
			fmt.Println(randFile)
			nums := [1000]int{}
			for i := 0; i < 1000; i++ {
				rand.Seed(time.Now().UnixNano())
				nums[i] = rand.Intn(1000)
			}
			numBytes, _ := json.Marshal(nums)
			ioutil.WriteFile(randFile, numBytes, 0666)
			signalChan <- true
		}(i)
	}

	for j := 0; j < cap(signalChan); j++ {
		<-signalChan
	}
	close(signalChan)

	for i := 0; i < 10; i++ {
		go func(i int) { // <<< 此处必须显示声明参数，否是依据存在下面使用时，又出现竞态条件，即多个i都成为了10
			var nums [1000]int
			randFile := fmt.Sprintf("%s/src/go_study/c38_channel/resources/rand-%d.txt", wd, i)
			numBytes, _ := ioutil.ReadFile(randFile)
			json.Unmarshal(numBytes, &nums)
			for i := 0; i < len(nums); i++ {
				for j := 0; j < len(nums)-1-i; j++ {
					if nums[j] > nums[j+1] {
						nums[j], nums[j+1] = nums[j+1], nums[j]
					}
				}
			}
			sortBytes, _ := json.Marshal(nums)
			sortFile := fmt.Sprintf("%s/src/go_study/c38_channel/resources/sort-%d.txt", wd, i)
			ioutil.WriteFile(sortFile, sortBytes, 0666)
			exitChan <- true
		}(i)
	}

	for j := 0; j < cap(exitChan); j++ {
		<-exitChan
	}
	close(exitChan)

	fmt.Println("over~")
}

/**
管道方向
只写管道：numChan chan <- int
只读管道：numChan <-chan int
读写管道：numChan chan int
*/
func prepareNums(nums int, numChan chan<- int) {
	for i := 1; i <= nums; i++ {
		numChan <- i
	}
	close(numChan)
}

func findPrimes(numChan <-chan int, primeChan chan<- int, exitChan chan<- bool) {
	for num := range numChan {
		if isPrime(num) {
			primeChan <- num
		}
	}
	exitChan <- true // 无数据可校验就退出，并发送退出信号
}

func printPrimes(primeChan <-chan int) {
	for prime := range primeChan {
		fmt.Println(prime)
	}
}

func test7() {
	startTmstp := time.Now().Unix()
	//cpu := runtime.NumCPU()
	//runtime.GOMAXPROCS(cpu - 1)
	nums := 300000
	checkers := 8

	numChan := make(chan int, 8)
	primeChan := make(chan int, 3)

	exitChan := make(chan bool, checkers)

	// 发数据
	go prepareNums(nums, numChan)

	// 取数据，然后校验是否是素数，是的话存入primeChan
	for i := 0; i < checkers; i++ {
		go findPrimes(numChan, primeChan, exitChan)
	}

	// 输出素数
	go printPrimes(primeChan)

	// 检验查找素数的8个协程是否全部工作完毕，如果是，就关闭primeChan，从而触发打印素数的协程关闭
	for i := 0; i < checkers; i++ {
		<-exitChan
	}

	close(primeChan)
	close(exitChan)
	stopTmstp := time.Now().Unix()
	fmt.Printf("total cost: %d,%d,%d\n", stopTmstp, startTmstp, stopTmstp-startTmstp)
}

func noBlockingBySelect() {
	numChan := make(chan int, 3)
	stringChan := make(chan string, 5)
	exitChan := make(chan bool,1)
	go func() {
		for i := 1; i < 5; i++ {
			numChan <- i
		}
		close(numChan)
	}()

	go func() {
		for i := 0; i < 3; i++ {
			stringChan <- fmt.Sprintf("i%d", i)
		}
		close(stringChan)
	}()

	go func() {
		case1Exit := false
		case2Exit := false
		for ;!case1Exit || !case2Exit;{
			/**
				1）先从第一个case 找，找不到就继续向下找，不产生阻塞，都未找到走default
				2）已经关闭的管道，返回基本数据类型，如 chan int 返回 0，chan string 返回""
				3) 需要收集所有case的退出标签，否则会进入死循环
			 */
			select {
			case num := <-numChan:
				if num != 0 {
					fmt.Println(num)
				}else{
					case1Exit = true
				}
				time.Sleep(time.Second)
			case str := <-stringChan:
				if str != ""{
					fmt.Println(str)
				}else{
					case2Exit = true
				}
				time.Sleep(time.Second)
			default:
				fmt.Println("无数据")
				//time.Sleep(time.Second)
			}
		}
		exitChan <- true
		fmt.Println("close(exitChan)")
	}()

	<- exitChan
	fmt.Println("over~")

}

/*
	使用 defer + recover()捕获协程中的异常，保证多个协程操作时，一个协程出问题不影响其余协程，问题协助暂停，其余继续工作
 */
func recoverBy(){
	numChan := make(chan int, 3)
	stringChan := make(chan string, 5)
	exitChan := make(chan bool,1)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("发生异常了，请及时处理：%s\n", err)
			}
		}()
		for i := 1; i < 5; i++ {
			fmt.Println(i/(2-i))
			numChan <- i
		}
		close(numChan)
	}()

	go func() {
		for i := 0; i < 3; i++ {
			stringChan <- fmt.Sprintf("i%d", i)
		}
		close(stringChan)
	}()

	go func() {
		case1Exit := false
		case2Exit := false
		for ;!case1Exit || !case2Exit;{
			/**
			1）先从第一个case 找，找不到就继续向下找，不产生阻塞，都未找到走default
			2）已经关闭的管道，返回基本数据类型，如 chan int 返回 0，chan string 返回""
			3) 需要收集所有case的退出标签，否则会进入死循环
			*/
			select {
			case num := <-numChan:
				if num != 0 {
					fmt.Println(num)
				}else{
					case1Exit = true
				}
				time.Sleep(time.Second)
			case str := <-stringChan:
				if str != ""{
					fmt.Println(str)
				}else{
					case2Exit = true
				}
				time.Sleep(time.Second)
			default:
				fmt.Println("无数据")
				//time.Sleep(time.Second)
			}
		}
		exitChan <- true
		fmt.Println("close(exitChan)")
	}()

	<- exitChan
	fmt.Println("over~")
}


func main() {
	//everyNumSum1()
	//everyNumSum2()
	//createChan()
	//create2()
	//test1()
	//writeAndRead()
	//test2()
	//test3()
	//test4()
	//test44()
	//test5()
	//test6()
	//test7()
	//noBlockingBySelect()
	recoverBy()
}
