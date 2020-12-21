package main

import "os"

/**
使用panic接口处理非受检异常
 */
func main(){
	//panic("a problem")

	// 成功创建文件
	//_,err := os.Create("/Users/huhao/softwares/go_proj/c00_example/temp.txt")
	_,err := os.Create("/aa/temp.txt")
	if err != nil{
		panic(err)
	}
}

/**
panic: open /aa/temp.txt: no such file or directory

goroutine 1 [running]:
newdemo.newdemo()
        /Users/huhao/softwares/go_proj/c00_example/p41-panic.go:12 +0x75

Process finished with exit code 2
 */