package main

import (
	"fmt"
	"time"
)

func main(){
	p := fmt.Println

	now := time.Now() // 当前实际
	p(now) // 2020-11-07 11:55:31.198298 +0800 CST m=+0.000072335

	then := time.Date(2009,11,17,20,34,58,651387237, time.UTC)
	p(then) // 手动构建日期 2009-11-17 20:34:58.651387237 +0000 UTC

	p(then.Year()) // 2009
	p(then.Month()) // November
	p(then.Day()) // 17
	p(then.Hour()) // 20
	p(then.Minute()) // 34
	p(then.Second()) // 58
	p(then.Nanosecond()) // 651387237
	p(then.Location()) // UTC

	p(then.Weekday()) // Tuesday

	p(then.Before(now)) // true
	p(then.After(now)) // false
	p(then.Equal(now)) // false

	diff := now.Sub(then) // 96175h20m32.546910763s
	p(diff)

	p(diff.Hours()) // 96175.34237414188
	p(diff.Minutes()) // 5.770520542448512e+06
	p(diff.Seconds()) // 3.4623123254691076e+08
	p(diff.Nanoseconds()) // 346231232546910763

	p(then.Add(diff)) // 2020-11-07 03:55:31.198298 +0000 UTC
	p(then.Add(-diff)) // 1998-11-28 13:14:26.104476474 +0000 UTC
}
