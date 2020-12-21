package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

func process(conn net.Conn){
	defer conn.Close()
	remoteIp := conn.RemoteAddr()
	fmt.Printf("removte ip: %s connnected server\n",remoteIp)
	buffer := make([]byte,1024)
	for {
		bytesCount, err := conn.Read(buffer)
		if err != nil{
			fmt.Printf("server process read error: %s\n",err)
			return
		}else{
			msg := string(buffer[:bytesCount])
			fmt.Printf("receive msg from %s: %s",remoteIp,msg)
			reply := fmt.Sprintf("message received at %v\n",time.Now())
			conn.Write([]byte(reply))
			if strings.EqualFold(strings.Trim(msg,"\n"),"exit"){
				conn.Write([]byte("goodby~"))
				return
			}
		}
	}
}

func server(){
	//conn, err := net.Dial("tcp", "0.0.0.0:9000")
	listener, err := net.Listen("tcp", "0.0.0.0:9000")
	if err != nil {
		fmt.Printf("server start listen error: %v\n",err)
	}
	fmt.Println("server listen at 0.0.0.0:9000")
	for{
		accepted, err := listener.Accept()
		if err != nil{
			fmt.Printf("server listener accept error: %v\n",err)
		} else{

			go process(accepted)
		}
	}

}

func main(){
	server()
}
