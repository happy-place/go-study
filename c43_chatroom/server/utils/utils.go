package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"go-study/c43_chatroom/common/message"
)

// 涉及一起消息传输
type Transfer struct{
	Conn net.Conn
	Buf [8064]byte// 传输时使用的缓存
}

// 传入连接，返回解析后的消息，以及异常
// 注：mes 是结构体，因此是引用类型，Unmarshal时，需要穿地址才能读入
// error 是接口，直接就是引用类型，直接赋值
func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		//err = errors.New(fmt.Sprintf("read pkg len err: %v\n",err))
		return
	}
	pkgLen := binary.BigEndian.Uint32(this.Buf[0:4])
	realLen, err := this.Conn.Read(this.Buf[:pkgLen])
	if realLen != int(pkgLen) || err != nil {
		//err = errors.New(fmt.Sprintf("read pkg body err: %v\n",err))
		return
	}
	err = json.Unmarshal(this.Buf[:pkgLen], &mes) // 装载进去
	if err != nil {
		//err = errors.New(fmt.Sprintf("unmarshal mes err: %v\n",err))
		return
	}
	fmt.Printf("mes: %v\n", mes)
	return
}

func (this *Transfer) WritePkg(data []byte) (err error) { // 消息通信前都先发送消息体长度。然后再发送具体消息
	var pkgLen = uint32(len(data)) // 消息序列化长度转为uint32
	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen) // 整数长度，转为[]byte
	n, err := this.Conn.Write(this.Buf[:4])               // 发送[]byte长度信息
	if n != 4 || err != nil {
		fmt.Printf("send message length err: %v\n", err)
		return // (err error) 匿名返回
	}
	mesLen, err := this.Conn.Write(data)
	if mesLen != int(pkgLen) || err != nil {
		fmt.Printf("send message body err: %v\n", err)
		return
	}
	return
}
