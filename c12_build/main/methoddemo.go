package main

import (
	"fmt"
	"go_study/c11_method/model"
	m2 "go_study/c11_method/model2" // 给包起别名
)

// go build methoddemo.go 编译好的可执行文件默认与 *.go 放在一起
// go build -o $GOPATH/bin/methoddemo methoddemo.go 指定编译输出路径
// go build -buildmode=plugin -o $GOPATH/pkg/methoddemo.so methoddemo.go
/**
go build methoddemo.go 编译好的可执行文件默认与 *.go 放在一起
go build -o $GOPATH/bin/methoddemo methoddemo.go 指定编译输出路径
go build -buildmode=plugin -o $GOPATH/pkg/methoddemo.so methoddemo.go 生成库，后期配合文档可以直接交付
 */


func main(){
	v,err := model.Operation(1,2,'+')
	if err !=nil{
		panic(err)
	}
	fmt.Println(v)

	//
	v,err = m2.Operation(1,2,'+')
}
