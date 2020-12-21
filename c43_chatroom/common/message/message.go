package message

const (
	LoginMesType  = "LoginMes"
	LoginResMesType  = "LoginResMes"
	RegisterMesType  = "RegisterMes"
)

type Message struct {
	Type string  `json:"type"`// 消息类型
	Data string  `json:"data"`// 消息
}

// 登录消息 client -> server
type LoginMes struct {
	UserId int `json:userid`
	UserPwd string `json:userpwd`
	UserName string `json:"username"`
}

// 登录响应消息 server -> client
type LoginResMes struct {
	Code int `json:"code"` // 500 未注册 200 登录成功
	Error string  `json:"error"`// 返回错误信息
}

type RegisterMes struct {

}