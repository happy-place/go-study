package main

import (
	"syscall"
	"os"
	"os/exec"
)

func main(){
	// 找到 ls 命令绝对路径，然后返回给binary
	binary,lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	args := []string {"ls","-a","-l","-h"}
	env := os.Environ()

	// 基于绝对路径发起调用
	execErr := syscall.Exec(binary,args,env)
	if execErr != nil {
		panic(execErr)
	}
}
