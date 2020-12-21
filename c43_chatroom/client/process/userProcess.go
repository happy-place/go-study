package process

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-study/c43_chatroom/client/utils"
	"go-study/c43_chatroom/common/message"
	"net"
)

type UserProcess struct {

}

func (this *UserProcess)Login(userId int,userPwd string)(err error){
	conn,err := net.Dial("tcp","localhost:8889")
	if err != nil {
		fmt.Printf("net dial err: %v\n",err)
		return err
	}
	defer conn.Close() // 延时关闭
	var mes message.Message
	mes.Type = message.LoginMesType
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd
	loginMesBytes, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Printf("loginMes marshal err: %v\n",err)
	}
	mes.Data = string(loginMesBytes)
	mesBytes, err := json.Marshal(mes)
	if err != nil {
		fmt.Printf("mes marshal err: %v\n",err)
		return err
	}

	transfer := &utils.Transfer{
		Conn: conn,
	}

	transfer.WritePkg(mesBytes)

	// 处理服务器响应
	mes, err = transfer.ReadPkg()
	if err != nil {
		fmt.Println("common.readPkg(conn) err:",err)
	}

	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if err != nil {
		fmt.Println("json.Unmarshal err:",err)
	}
	if loginResMes.Code == 200 {
		// 起一个协程，接收服务端推送消息
		go serverProcessMes(conn)
		for {
			ShowMenu()
		}
	}else if loginResMes.Code == 500 {
		// 登录失败
		fmt.Println(loginResMes.Error)
		err = errors.New(loginResMes.Error)
	}
	return err
}

func serverProcessMes(Conn net.Conn){
	// 不停监听服务端发送消息
	tf := &utils.Transfer{Conn:Conn}
	for {
		fmt.Printf("客户端正在读取服务器发送消息\n")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Printf("tf.ReadPkg err: %v\n",err)
			return
		}
		fmt.Printf("mes: %v\n",mes)
	}
}

