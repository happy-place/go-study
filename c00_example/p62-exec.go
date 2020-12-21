package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)


func main(){
	dateCmd := exec.Command("date")
	dateOut,err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut)) // 2020年11月 8日 星期日 16时50分03秒 CST

	grepCmd := exec.Command("grep","hello")
	grepIn,_ := grepCmd.StdinPipe()
	grepOut,_ :=grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes,_ := ioutil.ReadAll(grepOut)
	grepCmd.Wait() // 阻塞等待执行完毕
	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes)) // 全部输出 hello grep

	lsCmd := exec.Command("bash","-c","ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))

}
