package main

import (
	"fmt"
	"os"
)

// 创建文件，并返回文件句柄
func createFile(p string) *os.File {
	fmt.Println("createing")
	f,err := os.Create(p)
	if err != nil{
		panic(err)
	}
	return f
}

// 通过文件句柄，写入数据
func writeFile(f *os.File){
	fmt.Println("writing")
	fmt.Fprint(f,"data")
}

// 关闭文件句柄
func closeFile(f *os.File){
	fmt.Println("closing")
	f.Close()
}

// 创建文件，然后立即声明程序结束时，需要关闭文件，再执行写入操作
// defer closeFile(f) 如果放在 writeFile 之后，可能会在写过程中出异常，导致无法关闭
func main(){
	f := createFile("/Users/huhao/softwares/go_proj/c00_example/temp.txt")
	defer closeFile(f) // 等效于 java try-catch-finally
	writeFile(f)
}
