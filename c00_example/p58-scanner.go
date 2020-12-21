package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)

	// 从标准输入流接收，然后转为大写输出
	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}

	// 处理标准输入错误
	if err := scanner.Err();err != nil {
		fmt.Println(os.Stderr,"error:",err)
		os.Exit(1)
	}
}
