package apis

import (
	"awesomeProject1/src/mmogame/core"
	"awesomeProject1/src/mmogame/pb"
	"awesomeProject1/src/zinxtcpsocket/zinxface"
	"awesomeProject1/src/zinxtcpsocket/znet"
	"fmt"
	"google.golang.org/protobuf/proto"
)

type WorldChatApi struct {
	znet.BaseRouter
}

func (wc* WorldChatApi) Handler(request zinxface.IRequest)  {
	//解析用户传进来的Proto协议
	protoMessage:=&pb.Talk{}
	err:=proto.Unmarshal(request.GetData(),protoMessage)
	if err!=nil{
		fmt.Println("Talk Unmarshal err ",err)
		return
	}

	//当前的聊天数据是属于那个玩家发送的
	PlayerInterFace ,_:=request.GetConnection().GetProperty("Pid")
	if PlayerInterFace ==nil{return}
	//根据PlayerInterFace 获得对应player玩家对象
	player:= PlayerInterFace.(*core.Player)
	//将这个消息广播给其他玩家
	player.Talk(protoMessage.Content)
}
