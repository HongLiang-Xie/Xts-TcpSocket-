package main

import (
	"awesomeProject1/src/mmogame/apis"
	"awesomeProject1/src/mmogame/core"
	"awesomeProject1/src/zinxtcpsocket/utils"
	"awesomeProject1/src/zinxtcpsocket/zinxface"
	"awesomeProject1/src/zinxtcpsocket/znet"
	"fmt"
)

// 当前客户端建立连接之后的hook函数

func OnConnectionAdd(conn zinxface.IConnection){
	//创建player 对象
	player:=core.NewPlayer(conn)
	//告知客户端玩家Pid ,同步生成的玩家ID给客户端
	player.SyncPid()
	//广播玩家自己的出生点
	player.BroadCastStartPosition()
	//同步周边玩家，告知他们当前玩家上线。广播当前玩家的位置信息
	player.SyncSurrounding()
	//添加到世界管理器中
	core.WorldMgrObj.AddPlayer(player)
	//将该连接绑定一个Pid 玩家ID的属性
	conn.SetProperty("Pid",player)
	fmt.Println("player pid=",player.Pid)
}
// 当前客户端断开连接之后的hook函数
func OnConnectionLost(conn zinxface.IConnection){
	//获取player 对象
	playerInterFace, _ :=conn.GetProperty("Pid")
	if playerInterFace==nil{
		return
	}
	player:=playerInterFace.(*core.Player)
	//到世界管理器中删除玩家
	core.WorldMgrObj.RemovePlayer(player)
	fmt.Println("delete player pid=",player.Pid)

}

func main() {

	s:=znet.NewServer()
	utils.GlobalObject.TCPServer=s
	s.SetOnConnStart(OnConnectionAdd)
	s.SetOnConnStop(OnConnectionLost)

	s.AddRouter(2,&apis.WorldChatApi{})

	s.AddRouter(3,&apis.MoveApi{})

	s.AddRouter(4,&apis.Login{})

	s.Server()
}
