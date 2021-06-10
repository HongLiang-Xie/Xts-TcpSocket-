package znet

import (
	"awesomeProject1/src/zinxtcpsocket/utils"
	"awesomeProject1/src/zinxtcpsocket/zinxface"
	"fmt"
	"net"
)

//IServer 的接口实现 ，定义了一个Server的服务器模块
type Server struct {
	//服务器名称
	Name string
	//服务器绑定的IP版本
	IPVersion string
	//绑定的IP
	IP string
	//服务器绑定的端口
	Port int
	//当前Server消息管理模块 用于来绑定MsgID和对应处理业务API关系
	MsgHandler zinxface.IMessageHandle
	// server conn管理器
	ConnManager zinxface.IConnManger
	//该Server创建链接后自动调用Hook 函数-OnConnStart()
	OnConnStart func(conn zinxface.IConnection)
	//该Server销毁链接后自动调用Hook 函数-OnConnStop()
	OnConnStop func(conn zinxface.IConnection)
}

/*func CCallbackToClient(conn *net.TCPConn,buff []byte,len int) error {
	fmt.Println("Conn Handle CCallbackToClient...")
	if _,err:=conn.Write(buff[:len]);err!=nil{
	fmt.Println("write back err",err)
	return errors.New("CallbackToClient error")
	}
	return nil
}
*/
//初始化 Server
func NewServer() zinxface.IServer {
	init := &Server{
		Name:      utils.GlobalObject.Name,
		IPVersion: "tcp4",
		IP:         utils.GlobalObject.Host,
		Port:       utils.GlobalObject.TcpPort,
		MsgHandler: 	NewMessageHandle(),
		ConnManager: NewConnManager(),
	}
	return init
}
func (s *Server) Start() {
	fmt.Println(utils.GlobalObject)
	go func() {

		//开启工作池和消息队列
		s.MsgHandler.StartWorkerPool()
		//创建TCP
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))

		if err != nil {
			fmt.Println("resolve tcp addr error:\n", err.Error())
			return
		}
		fmt.Printf("resolve tcp addr[%s] sucess\n", fmt.Sprintf("%s:%d", s.IP, s.Port))
		//监听
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen tcp  error:\n", err.Error())
			return
		}
		fmt.Printf("listen tcp[%s]  sucess\n", s.Name)
		//阻塞等待客户端连接
		var cId uint32
		cId =0
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept error:\n", err.Error())
				continue
			}
			//设置最大连接个数，超过则不连接
			if s.ConnManager.Len()>=utils.GlobalObject.MaxConn{
				//TODO 给客户端返回一个超过最大链接的错误包
				conn.Close()
				continue
			}
			//将处理新连接的业务方法 和conn 进行绑定 得到我们要使用的链接模块
			dealConn:=NewConnection(s,conn,cId,s.MsgHandler)
			go dealConn.Start()
			cId++
		}

	}()
}
func (s *Server) Stop() {
//开始资源回收
	s.ConnManager.ClearConn()
	fmt.Println("server stop")

}
func (s *Server) Server() {
	//启动服务器
	s.Start()
	//中间处理额外业务

	//阻塞
	select {

	}
}
//路由功能:给当前的服务注册一个路由方法，供客户端的链接处理使用
func (s *Server) AddRouter(msgID uint32,router zinxface.IRouter){
	s.MsgHandler.AddRouter(msgID,router)
	fmt.Println("Add Router success")

}
func (s *Server)GetConnManager()zinxface.IConnManger{

	return s.ConnManager
}

//注册 OnConnStart 钩子函数的方法

func (s *Server)SetOnConnStart( hookFunc func(conn zinxface.IConnection)){

	s.OnConnStart=hookFunc
}
//注册 OnConnStop 钩子函数的方法
func (s *Server)SetOnConnStop( hookFunc func(conn zinxface.IConnection)){
	s.OnConnStop=hookFunc
}
//调用 OnConnStart 钩子函数的方法
func (s *Server)CallOnConnStart(conn zinxface.IConnection){
		if s.OnConnStart!=nil{
			fmt.Println("CallOnConnStart...")
			s.OnConnStart(conn)
		}
}
//调用 OnConnStop 钩子函数的方法
func (s *Server)CallOnConnStop(conn zinxface.IConnection){
		if s.OnConnStop!=nil{
			fmt.Println("CallOnConnStop...")
			s.OnConnStop(conn)
		}


}
