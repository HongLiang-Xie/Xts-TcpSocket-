package znet

import (
	"awesomeProject1/src/zinxtcpsocket/utils"
	"awesomeProject1/src/zinxtcpsocket/zinxface"
	"errors"
	"fmt"
	"io"
	"net"
	"sync"
)

/*
链接模块
*/
type Connection struct {

	//当前Conn属于的Server
	TcpServer zinxface.IServer
	//当前链接的socket TCP套接字
	Conn *net.TCPConn
	//链接的ID
	ConnID uint32
	//当前链接的状态
	isClosed bool

	//告知当前链接已经退出的/停止 channel
	ExitChannel chan bool
	//消息的管理 msgID 和对应处理业务API关系
	MsgHandler zinxface.IMessageHandle

	//消息通道
	MsgChannel chan []byte
	//链接属性集合
	Property map[string] interface{}
	//保护连接属性的锁
	PropertyLock sync.RWMutex
}

func NewConnection(tcpserver zinxface.IServer,conn *net.TCPConn, id uint32, msgHandler zinxface.IMessageHandle) zinxface.IConnection {
	init := &Connection{
		tcpserver,
		conn,
		id,
		false,
		make(chan bool),
		msgHandler,
		make(chan []byte),
		make(map[string] interface{}),
		*new(sync.RWMutex),
	}
	//将conn加入conn管理器中
	tcpserver.GetConnManager().Add(init)
	return init
}
func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running..")
	defer fmt.Println("ConnID=", c.ConnID, "Reader is exit", c.RemoteAddr().String())
	defer c.Stop()
	for {
		/*buff:=make([]byte,utils.GlobalObject.MaxPackageSize)
		_,err:=c.Conn.Read(buff)
		if err!=nil{
			fmt.Println("recv buf err",err.Error())
			continue
		}*/

		//创建一个拆包解包对象
		dp := NewDataPackage()
		//-----拆包过程------
		//1 第一次从conn读，把包的Head读出来
		headData := make([]byte, dp.GetHeadLen())
		_, err := io.ReadFull(c.GetTCPConnection(), headData)
		if err != nil {
			fmt.Println("read head error", err)
			break
		}

		msgHead, err := dp.UnPackage(headData)
		if err != nil {
			fmt.Println("unpackage err", err)
			break
		}
		var data []byte
		if msgHead.GetMessageLen() > 0 {
			//msg有数据，需要进行二次读取
			//2 第二次从conn读，把包的dataLen读出来
			data = make([]byte, msgHead.GetMessageLen())
			_, err := io.ReadFull(c.GetTCPConnection(), data)
			if err != nil {
				fmt.Println("server unpackage data error", err)
				break
			}
		}
		msgHead.SetMessageData(data)
		//得到当前conn数据的Request请求数据
		req := Request{
			conn:        c,
			messageData: msgHead,
		}
		//c.Conn.Write(msgHead.GetMessageData())
		if utils.GlobalObject.WorkPoolSize>0{
			c.MsgHandler.SendMsgToTaskQueue(&req)
		}else{
			//执行注册的路由方法
			go c.MsgHandler.DoMessageHandler(&req)
		}


		//从路由中，找到注册绑定的conn对应的router调用
		//调用当前链接所绑定的HandleAPI
		/*if err:=c.handleAPI(c.Conn,buff,len);err!=nil{
			fmt.Println("ConnID",c.ConnID,"handle is error",err.Error())
			break
		}*/

	}

}
func (c *Connection) StartWriter() {
	fmt.Println("Sender Goroutine is running..")
	for {

		select {
		case data := <-c.MsgChannel:
			if _, err := c.Conn.Write(data); err != nil {
				fmt.Println("Send data error:", err)
				return
			}
		//reader 获取客户端退出，这里也直接退出线程
		case <-c.ExitChannel:
			return

		}
	}

}

//启动链接 开始准备工作
func (c *Connection) Start() {
	fmt.Println("Conn Start()..ConnID=", c.ConnID)
	//启动从当前链接的读数据的业务
	go c.StartReader()
	//启动从当前链接写的业务
	go c.StartWriter()
	//创建链接后用户需要调用 服务器OnConnStart处理启动链接业务
	c.TcpServer.CallOnConnStart(c)

}

//停止链接
func (c *Connection) Stop() {
	defer fmt.Println("ConnID=", c.ConnID, "Sender is exit", c.RemoteAddr().String())
	if c.isClosed == true {
		return
	}
	c.isClosed = true
	//销毁链接后给用户提供调用服务器OnConnStop处理销毁业务
	c.TcpServer.CallOnConnStop(c)

	//关闭socket链接
	c.Conn.Close()

	//将conn 从conn管理器中去掉
	c.TcpServer.GetConnManager().Remove(c)
	//给writer 发送关闭消息
	c.ExitChannel<-true
	close(c.ExitChannel)
	close(c.MsgChannel)
}

//获取当前链接的绑定socket conn
func (c *Connection) GetTCPConnection() *net.TCPConn {

	return c.Conn
}

//获取当前链接模块的链接ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

//获取客户端的 TCP状态 IP port
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()

}

// 发送数据 将数据发送给客户端
func (c *Connection) SendMsg(msgid uint32, data []byte) error {
	if c.isClosed == true {
		return errors.New("Connection close when send msg")
	}
	//将data进行封包 Msgid Msglen  msgdata
	dp := NewDataPackage()
	binaryMsg, err := dp.Package(NewMessagePackage(msgid, data))
	if err != nil {
		fmt.Println("Pack error msg id=", msgid)
		return errors.New("Pack error msg")
	}
	//给writer 发送需要给客户端发送的消息
	c.MsgChannel<-binaryMsg

	return nil
}
//设置链接属性
func (c *Connection) SetProperty(key string ,value interface{}){

	c.PropertyLock.Lock()
	defer c.PropertyLock.Unlock()
	c.Property[key]=value
}
//获取链接属性
func (c *Connection)GetProperty(key string) (interface{} ,error){
	c.PropertyLock.Lock()
	defer c.PropertyLock.Unlock()
	value,err:=c.Property[key]
	if err{

		return value ,nil
	}
return nil ,errors.New("not found key="+key+"property")
}
//移除链接属性
func (c *Connection)RemoveProperty(key string){
	c.PropertyLock.Lock()
	defer c.PropertyLock.Unlock()

	delete(c.Property,key)

}
func (c *Connection)GetAllPlayer()zinxface.IServer{

	 return c.TcpServer
}