package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func connect(){
	conn, err := net.Dial("tcp", "0.0.0.0:9000")
	if err != nil {
		fmt.Printf("connected to server error: %v\n",err)
	}
	fmt.Println("connected to server success")
	defer conn.Close()
	buffer := make([]byte,1024)
	reader := bufio.NewReader(os.Stdin)
	for {
		bytes, err := reader.ReadBytes('\n')
		if err != nil {
			fmt.Printf("client read from stdin error error: %v\n",err)
		}
		conn.Write(bytes)
		bytesCount, err := conn.Read(buffer)
		fmt.Printf(string(buffer[:bytesCount]))
	}
}

func main(){
	connect()
}