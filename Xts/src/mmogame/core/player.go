package core

import (
	"awesomeProject1/src/mmogame/pb"
	"awesomeProject1/src/zinxtcpsocket/zinxface"
	"fmt"
	"google.golang.org/protobuf/proto"
	"math/rand"
	"sync"
)

//玩家对象
type Player struct {
	Pid int32 //玩家ID TODO 从数据库获取
	Conn zinxface.IConnection //当前玩家的连接（用于和客户端的连接）
	X float32 //平面x轴坐标 TODO 从数据库获取
	Y float32 //高度 TODO 从数据库获取
	Z float32 //平面y轴坐标 TODO 从数据库获取
	V float32 //旋转角度0-360 TODO 从数据库获取
}
var PidGen int32=1 //玩家Pid计数器
var PidGenLock sync.RWMutex //计数器锁
func NewPlayer(conn zinxface.IConnection)*Player{

	PidGenLock.Lock()
	defer PidGenLock.Unlock()
	id:=PidGen
	PidGen++
	p:=&Player{
		Pid:  id,
		Conn: conn,
		X:    float32(160+rand.Intn(50)),
		Y:    0,
		Z:    float32(132+rand.Intn(50)),
		V:    0,
	}
	return p
}
func(p*Player)SendMsg(msgID uint32,data proto.Message){

	//将proto.Message结构体转化为二进制
	msg,err:=proto.Marshal(data)
	if err!=nil{
		fmt.Println("player sendmsg err:(marshal msg err)",err)
		return
	}
	if p.Conn==nil{
		fmt.Println("player sendmsg err:(player is nil)")
		return
	}
	p.Conn.SendMsg(msgID,msg)
}
//告知客户端玩家Pid ,同步生成的玩家ID给客户端
func(p*Player) SyncPid(){
	//组建msgID 1的proto 数据
	data:=&pb.SyncPid{Pid: p.Pid}
	//发送数据
	p.SendMsg(1,data)
}
//广播玩家自己的出生点
func (p* Player) BroadCastStartPosition()  {
	//组建msgID 200的proto 数据
	data:=&pb.BroadCast{
		Pid:  p.Pid,
		Tp:   2,
		Data: &pb.BroadCast_P{P: &pb.Position{
			X: p.X,
			Y: p.Y,
			Z: p.Z,
			V: p.V,
		},
		},
	}
	//发送数据
	p.SendMsg(200,data)

}
//发送世界广播消息
func (p* Player) Talk(conent string ) {

	data:=&pb.BroadCast{
		Pid:  p.Pid,
		Tp:   1,
		Data: &pb.BroadCast_Content{
			Content: conent,
			},
		}
	players:=WorldMgrObj.GetAllPlayer()
	for _,player:=range players{
		player.SendMsg(200,data)
	}
}
//同步玩家上线位置信息
func (p* Player) SyncSurrounding()  {
	//1获取当前周边玩家的位置信息
	Pids:=WorldMgrObj.AOIMgr.GetSpriteByPoint(p.X,p.Z)
	players:=make([]*Player,0,len(Pids))

	for _,pid:=range Pids{
		players=append(players,WorldMgrObj.GetPlayer(int32(pid)))
	}
	//2将当前玩家的位置信息通过msgID=200 发给周边的玩家（我看到周边玩家）
	protoMessage:=&pb.BroadCast{
		Pid:  p.Pid,
		Tp:   2, //tp 2代表广播位置
		Data: &pb.BroadCast_P{
			P: &pb.Position{
			X: p.X,
			Y: p.Y,
			Z: p.Z,
			V: p.V,
		},
		},
	}
	for _,player:= range  players{

		player.SendMsg(200,protoMessage)
	}
	//3 将周围的全部玩家的位置信息通过msgID=202发送给当前的玩家客户端（让其他人看到我）
	//制作msgID 202数据
	players_protoMessage:=make([]*pb.Player,0,len(players))

	for _,player:= range  players{
		player_protoMessage:=&pb.Player{
			Pid: player.Pid,
			Pos: &pb.Position{
				X: player.X,
				Y: player.Y,
				Z: player.Z,
				V: player.V,
			},
		}
		players_protoMessage=append(players_protoMessage,player_protoMessage)
	}

	SyncPlayers_protoMessage:=&pb.SyncPlayers{

		Ps: players_protoMessage[:],
	}

	p.SendMsg(202,SyncPlayers_protoMessage)

}

func (p*Player) UpdatePosition(x,y,z,v float32)  {
	//更新当前玩家坐标
	p.X=x
	p.Y=y
	p.Z=z
	p.V=v
	//组建广播协议 msgID=200 Tp=4
	ProtoMessage:=&pb.BroadCast{
		Pid:  p.Pid,
		Tp:   4,
		Data: &pb.BroadCast_P{P: &pb.Position{
			X: p.X,
			Y: p.Y,
			Z: p.Z,
			V: p.V,
		}},
	}

	//获取周围玩家
	players:=p.GetSurroundingPlayers()
	//给周围玩家发送当前玩家更新的位置
	for _,player:=range players{

		player.SendMsg(200,ProtoMessage)
	}
}
//获取周围玩家
func (p* Player) GetSurroundingPlayers() []*Player {
	//获得AOI九宫格内所有玩家Pid
	Pids:=WorldMgrObj.AOIMgr.GetSpriteByPoint(p.X,p.Z)
	players:=make([]*Player,0,len(Pids))
	//通过Pid获取玩家对象
	for _,pid:=range Pids{

		players=append(players,WorldMgrObj.GetPlayer(int32(pid)))

	}
	return players
}