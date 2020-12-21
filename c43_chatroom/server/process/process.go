package process

import (
	"fmt"
	"io"
	"net"
	"go-study/c43_chatroom/common/message"
	"go-study/c43_chatroom/server/utils"
)

type Processor struct {
	Conn net.Conn
}

// ServverProcessMes 根据客户端发送消息种类不同，决定使用哪个函数来处理
func (this *Processor)serverProcessMes(mes *message.Message)(err error){
	switch mes.Type {
	case message.LoginMesType:
		userProcessor := &UserProcessor{Conn: this.Conn}
		// 处理登录逻辑
		err = userProcessor.ServerProcessLogin(mes)
	case message.RegisterMesType:
		// 处理注册逻辑
    default:
		fmt.Println("消息类型不存在，无法处理...")
	}
	return
}


func (this *Processor) Process()(err error){
	transfer := &utils.Transfer{
		Conn: this.Conn,
	}
	for {
		mes, err := transfer.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务器也退出...")
			}else{
				fmt.Printf("read msg err: %v\n",err)
			}
			return err
		}
		fmt.Println(mes)
		err = this.serverProcessMes(&mes)
		if err != nil {
			return err
		}
	}
}