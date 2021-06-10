package zinxface

import "net"

//定义链接模块的抽象层
type IConnection interface {
	//启动链接
	Start()
	//停止链接
	Stop()
	//获取当前链接的绑定socket conn
	GetTCPConnection()*net.TCPConn
	//获取当前链接模块的链接ID
	GetConnID() uint32
	//获取客户端的 TCP状态 IP port
	RemoteAddr() net.Addr
	// 发送数据 将数据发送给客户端
	SendMsg(msgid uint32,data []byte ) error
	//设置链接属性
	SetProperty(key string ,value interface{})
	//获取链接属性
	GetProperty(key string) (interface{} ,error)
	//移除链接属性
	RemoveProperty(key string)
	GetAllPlayer()IServer
}
type HandleFunc func(*net.TCPConn,[]byte,int)error