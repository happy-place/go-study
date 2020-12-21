package main

import (
	"fmt"
	"os"
)

/**
使用 os.exit() 时。defer 最终不会执行
go 中任何非零状态退出都要使用 os.exit(0
 */
func main(){
	defer fmt.Println("!")
	os.Exit(3)
}