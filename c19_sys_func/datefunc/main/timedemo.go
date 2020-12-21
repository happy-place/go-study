package main

import (
	"fmt"
	"time"
	"math/rand"
	"strconv"
)

func now(){
	n := time.Now()
	fmt.Printf("now: %v, type:%T\n",n,n)
	// now: 2020-11-14 11:20:18.859929 +0800 CST m=+0.000085800, type:time.Time
	fmt.Printf("%v年%v月%v日%v时%v分%v秒%v纳秒\n",n.Year(),n.Month(),n.Day(),n.Hour(),n.Minute(),n.Second(),n.Nanosecond())
	//2020年November月14日11时22分52秒79510000纳秒
}

func format(){
	n := time.Now()
	// 2020-11-14 14:42:12

	// 格式化日期 第一种方式
	dateString := fmt.Sprintf("%d-%d-%d %d:%d:%d\n",n.Year(),n.Month(),n.Day(),n.Hour(),n.Minute(),n.Second())
	fmt.Println(dateString)

	// 格式化日期 第二种方式
	now := time.Now()
	dateAndTime := now.Format("2006-01-02 15:04:05") // 必须写成这个字符串，才能正常解析
	fmt.Printf("dt: %v, type: %T\n",dateAndTime,dateAndTime)

	dtString := now.Format("2006-01-02") // 或 2006/01/02 亦可
	fmt.Printf("dt: %v, type: %T\n",dtString,dtString) // dt: 2020/11/14, type: string

	timeString := now.Format("15:04:05")
	fmt.Printf("dt: %v, type: %T\n",timeString,timeString) // dt: 14:51:52, type: string

	year := now.Format("2006")
	month := now.Format("01")
	day := now.Format("02")
	hour := now.Format("15")
	minute := now.Format("04")
	second := now.Format("05")
	fmt.Printf("%s年%s月%s日 %s时:%s分:%s秒\n",year,month,day,hour,minute,second) // 2020年11月14日 14时:55分:58秒

}

/**
时间日期相关常量
const (
	Nanosecond  Duration = 1
	Microsecond          = 1000 * Nanosecond
	Millisecond          = 1000 * Microsecond
	Second               = 1000 * Millisecond
	Minute               = 60 * Second
	Hour                 = 60 * Minute
)
*/
func constant(){
	count := 0
	for {
		time.Sleep(time.Second * 2)
		count ++
		n := time.Now()
		fmt.Printf("%d-%d-%d %d:%d:%d\n",n.Year(),n.Month(),n.Day(),n.Hour(),n.Minute(),n.Second())
		if count == 5{
			break
		}
	}

}

func randSeed(){
	target := 99
	count := 0
	for{
		// 纳秒维度 随机种子（重复概率低）
		rand.Seed(time.Now().UnixNano()) // UnixNano 时间戳
		// 秒维度 随机种子（重复概率高
		//rand.Seed(time.Now().Unix()) // Unix 时间戳
		num := rand.Intn(100) + 1
		count ++
		if num == target {
			// 1970-01-01至今经历的1605337806866509000纳秒，经历的毫秒 1605337806
			fmt.Printf("首次出现%v共经历%d次随机数计算",target,count)
			break
		}
	}
}

func unixXX(){
	n := time.Now()
	nano := n.UnixNano()
	tmstp := n.Unix()
	fmt.Printf("1970-01-01至今经历的%v纳秒，经历的%v秒\n",nano,tmstp)
}

func countTime(){
	str := ""
	start := time.Now().Unix()
	for i:=0;i<100000;i++{
		str += strconv.Itoa(i)
	}
	end := time.Now().Unix()
	fmt.Printf("总共耗时: %v",end - start)
}


func main(){
	//now()
	//format()
	//constant()
	//randSeed()
	//unixXX()
	countTime()

}

