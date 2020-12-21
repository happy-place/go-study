package main

import (
	"fmt"
	"go-study/c43_chatroom/server/model"
	"go-study/c43_chatroom/server/process"
	"net"
	"time"
)

func handler(conn net.Conn){
	defer conn.Close()
	processor := &process.Processor{Conn: conn}
	err := processor.Process()
	if err != nil {
		fmt.Printf("客户端和服务端通信协程出问题：%v\n",err)
		return
	}
}

func initUserDao(){
	model.MyUserDao = model.NewUserDao(pool)
}

func main(){
	// 初始化redis连接池
	initPool("localhost:6379",16,0,300 * time.Second)
	initUserDao()
	fmt.Println("服务器在8889端口监听")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()
	if err != nil {
		fmt.Printf("server listen err: %v\n",err)
		return // 异常直接退出
	}
	for {
		fmt.Println("等待客户接入...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("Listen accept err: %v\n",err)
		}
		go handler(conn)
	}
}
