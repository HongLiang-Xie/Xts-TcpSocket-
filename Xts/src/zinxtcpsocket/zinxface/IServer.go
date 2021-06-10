package zinxface

//定义一个服务器接口
type IServer interface {
	//启动服务器
	Start()
	//停止服务器
	Stop()
	//运行服务器
	Server()
	//路由功能:给当前的服务注册一个路由方法，供客户端的链接处理使用
	AddRouter(msgID uint32,router IRouter)
	//获取当前server conn管理器
	GetConnManager()IConnManger

	//注册 OnConnStart 钩子函数的方法

	SetOnConnStart( hookFunc func(conn IConnection))
	//注册 OnConnStop 钩子函数的方法
	SetOnConnStop( hookFunc func(conn IConnection))
	//调用 OnConnStart 钩子函数的方法
	CallOnConnStart(conn IConnection)
	//调用 OnConnStop 钩子函数的方法
	CallOnConnStop(conn IConnection)

}
