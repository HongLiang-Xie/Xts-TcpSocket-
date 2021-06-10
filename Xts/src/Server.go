package main

import (
	"awesomeProject1/src/zinxtcpsocket/utils"
	"awesomeProject1/src/zinxtcpsocket/zinxface"
	"awesomeProject1/src/zinxtcpsocket/znet"
	"fmt"
)

/*
基于zinx框架 开发的服务器应用程序
*/
// ping test 自定义路由
type PingRouter struct {
	znet.BaseRouter

}
//test PreHandle
func (this*PingRouter) PreHandle(request zinxface.IRequest){
	/*fmt.Println("Call Router PreHandle")
	_,err:=request.GetConnection().GetTCPConnection().Write([]byte("before ping...\n"))
	if err!=nil{

		fmt.Println("call back before ping...error")
	}

	 */

}
//test Handler
func (this*PingRouter)Handler(request zinxface.IRequest){
	fmt.Println("Call Router Handler")
	/*_,err:=request.GetConnection().GetTCPConnection().Write([]byte("ping...ping...ping...\n"))
	if err!=nil{

		fmt.Println("call back  ping...error")
	}
	*/
	fmt.Println("recv from client:msgid=",request.GetDataID(),"msg:", string(request.GetData()))
	_ = request.GetConnection().SendMsg(200, []byte("ping...ping..."))
}
//test PostHandle
func (this*PingRouter)PostHandle(request zinxface.IRequest){
/*	fmt.Println("Call Router PostHandle")
	_,err:=request.GetConnection().GetTCPConnection().Write([]byte("after ping...\n"))
	if err!=nil{

		fmt.Println("call back after ping...error")
	}
*/
}
type HelloRouter struct {
	znet.BaseRouter

}
//test PreHandle
func (this*HelloRouter) PreHandle(request zinxface.IRequest){
	/*fmt.Println("Call Router PreHandle")
	_,err:=request.GetConnection().GetTCPConnection().Write([]byte("before ping...\n"))
	if err!=nil{

		fmt.Println("call back before ping...error")
	}

	*/

}
//test Handler
func (this*HelloRouter)Handler(request zinxface.IRequest){
	fmt.Println("Call Router Handler")
	/*_,err:=request.GetConnection().GetTCPConnection().Write([]byte("ping...ping...ping...\n"))
	if err!=nil{

		fmt.Println("call back  ping...error")
	}
	*/
	fmt.Println("recv from client:msgid=",request.GetDataID(),"msg:", string(request.GetData()))
	_ = request.GetConnection().SendMsg(201, []byte("hello this zinxtcpsocket server"))
}
//test PostHandle
func (this*HelloRouter)PostHandle(request zinxface.IRequest){
	/*	fmt.Println("Call Router PostHandle")
		_,err:=request.GetConnection().GetTCPConnection().Write([]byte("after ping...\n"))
		if err!=nil{

			fmt.Println("call back after ping...error")
		}
	*/
}

func InitStart(conn zinxface.IConnection){

	_ = conn.SendMsg(202, []byte("DoConnection begin...."))
	conn.SetProperty("xhl","成都")
}
func InitStop(conn zinxface.IConnection){

	fmt.Println("ConnID=" ,conn.GetConnID(),"is lost....")
	fmt.Println(conn.GetProperty("xhl"))
	conn.RemoveProperty("xhl")
}
func main()  {

	//初始化服务器
	s:=znet.NewServer()
	//注册自定义Router
	s.AddRouter(0,&PingRouter{})
	s.AddRouter(1,&HelloRouter{})
	utils.GlobalObject.TCPServer=s
	//启动服务器
	s.SetOnConnStart(InitStart)
	s.SetOnConnStop(InitStop)
	s.Server()
}
