package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func check(e error){
	if e != nil {
		panic(e)
	}
}

func main(){
	// 一次性渠道全部文件内容
	dat,err := ioutil.ReadFile("/Users/huhao/softwares/go_proj/c00_example/temp.txt")
	check(err)
	fmt.Println(string(dat))

	// 打开文件，并注册关闭事件
	f,err := os.Open("/Users/huhao/softwares/go_proj/c00_example/temp.txt")
	defer f.Close()
	check(err)

	b1 := make([]byte,5)
	n1,err := f.Read(b1) // 读取5个字节长度，n1 返回的就是读取字节长度，b1承接的就是内容
	check(err)
	fmt.Printf("%d bytes: %s\n",n1,string(b1)) // 5 bytes: /usr/

	o2,err := f.Seek(6,0) // 手动拨动游标到6位置， 不设置读取长度（有多大缓冲，读多大）
	check(err)
	b2 := make([]byte,2)
	n2,err := f.Read(b2) // n2 为读取字节长度
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n",n2,o2,string(b2)) // 2 bytes @ 6: oc

	o3,err := f.Seek(6,0) // 重新拨回到下标为6位置不设置读取长度
	check(err)
	b3 := make([]byte,2)
	n3,err := io.ReadAtLeast(f,b3,2) // 最少读取2位
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n",n3,o3,string(b3)) // 2 bytes @ 6: oc

	_,err = f.Seek(0,0) // 重新拨回到0位置
	check(err)

	r4 := bufio.NewReader(f) // 创建buffer reader对象
	b4,err := r4.Peek(5)  // 读取5个字节长度，b4承接的就是读取内容
	check(err)
	fmt.Printf("5 bytes: %s\n",string(b4)) // 5 bytes: /usr/

}
