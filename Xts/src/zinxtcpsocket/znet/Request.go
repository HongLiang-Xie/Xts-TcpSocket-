package znet

import (
	"awesomeProject1/src/zinxtcpsocket/zinxface"
)

type Request struct {
	// 已经和客户端建立好的链接
	conn zinxface.IConnection
	//客户端请求的数据
	messageData zinxface.IMessage
}

// 得到当前链接
func (r*Request) GetConnection() zinxface.IConnection{

	return r.conn
}
//得到请求的消息数据
func (r*Request) GetData() []byte{

 return  r.messageData.GetMessageData()
}

//得到消息id
func (r*Request) GetDataID() uint32{

	return  r.messageData.GetMessageID()
}