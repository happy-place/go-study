package main

import (
	"fmt"
	"time"
)

func main(){
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now) // 2020-11-07 12:01:12.281219 +0800 CST m=+0.000074762
	fmt.Println(secs) // 1604721672
	fmt.Println(nanos) // 1604721672281219000

	millis := nanos / 1000000
	fmt.Println(millis) // 1604721672281

	fmt.Println(time.Unix(secs,0)) // 2020-11-07 12:01:12 +0800 CST
	fmt.Println(time.Unix(0,nanos)) // 2020-11-07 12:01:12.281219 +0800 CST
}
