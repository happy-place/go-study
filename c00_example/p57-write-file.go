package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func checkErr(e error){
	if e!=nil{
		panic(e)
	}
}

// 拼接路径
func getPath(name string) string{
	dir,_ := os.Getwd()
	return dir+name
}

func main(){
	// 直接写出
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile(getPath("temp1.txt"),d1,0644)
	checkErr(err)

	// 获取文件句柄，并注册关闭函数
	f,err := os.Create(getPath("temp2.txt"))
	checkErr(err)
	defer f.Close()

	// 手动写出字节数组
	d2 := []byte{115,111,109,101,10}
	n2,err := f.Write(d2)
	checkErr(err)
	fmt.Printf("wrote %d bytes\n",n2) // wrote 5 bytes

	// 手动写出字符串，并立刻同步刷新缓冲到磁盘
	n3,err := f.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n",n3) // wrote 7 bytes
	err = f.Sync()
	checkErr(err)

	// 带缓冲写出字符串，并立即刷新到磁盘
	w := bufio.NewWriter(f)
	n4,err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n",n4) // wrote 9 bytes
	err = w.Flush()
	checkErr(err)

}
