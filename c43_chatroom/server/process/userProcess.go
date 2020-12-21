package process

import (
	"encoding/json"
	"fmt"
	"go-study/c43_chatroom/common/message"
	"go-study/c43_chatroom/server/model"
	"go-study/c43_chatroom/server/utils"
	"net"
)

type UserProcessor struct {
	Conn net.Conn

}

func (this *UserProcessor)ServerProcessLogin(mes *message.Message)(err error){
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Printf("json.Unmarshal fail err: %v\n",err)
		return
	}

	var resMes message.Message
	resMes.Type = message.LoginResMesType

	var loginResMes message.LoginResMes

	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if err != nil {
		loginResMes.Code = 500 // 用户不存在
		loginResMes.Error = "该用户不存在，请注册再使用"
	}else {
		loginResMes.Code = 200
		fmt.Printf("%s 登录成功\n",user.UserName)
	}

	//if loginMes.UserId == 200 && loginMes.UserPwd == "123456" {
	//	loginResMes.Code = 200 // 登录成功
	//}else{
	//	loginResMes.Code = 500 // 用户不存在
	//	loginResMes.Error = "该用户不存在，请注册再使用"
	//}

	loginResMesBytes, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal fail ",err)
		return
	}
	resMes.Data = string(loginResMesBytes)
	resMesBytes, err := json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail ",err)
		return
	}
	transfer := &utils.Transfer{
		Conn: this.Conn,
	}
	err = transfer.WritePkg(resMesBytes)
	return err
}



