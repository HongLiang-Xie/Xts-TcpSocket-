
package apis

import (
	"awesomeProject1/src/mmogame/core"
	"awesomeProject1/src/mmogame/pb"
	"awesomeProject1/src/zinxtcpsocket/zinxface"
	"awesomeProject1/src/zinxtcpsocket/znet"
"fmt"
"google.golang.org/protobuf/proto"
)

type MoveApi struct {
	znet.BaseRouter
}

func (ma* MoveApi) Handler(request zinxface.IRequest)  {
	//解析用户传进来的Proto协议
	protoMessage:=&pb.Position{}
	err:=proto.Unmarshal(request.GetData(),protoMessage)
	if err!=nil{
		fmt.Println("Position Unmarshal err ",err)
	}

	//当前的聊天数据是属于那个玩家发送的
	PlayerInterFace ,_:=request.GetConnection().GetProperty("Pid")

	//根据PlayerInterFace 获得对应player玩家对象
	player:= PlayerInterFace.(*core.Player)
	//fmt.Println("Pid=",player.Pid,"position:",protoMessage)
	//将这个消息广播给其他玩家
	player.UpdatePosition(player.X,player.Y,player.Z,player.V)



}
