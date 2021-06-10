package main

import (
	"awesomeProject1/src/mmogame/pb"
	"awesomeProject1/src/zinxtcpsocket/znet"
	"fmt"
	"google.golang.org/protobuf/proto"
	"io"
	"net"
	"time"
)

func main() {


	conn,err:=net.Dial("tcp4","localhost:8080")
	if err!=nil{

		fmt.Println("conn failed")
		return
	}
	for{
		//发送封包消息
		dp:=znet.NewDataPackage()
		data:=&pb.SyncPid{Pid:0}
		login,_:= proto.Marshal(data)
		binaryMsg,err:= dp.Package(znet.NewMessagePackage(1,login))
		if err!=nil{ return}
		if _,err:=conn.Write(binaryMsg);err!=nil{
			return
		}
		binaryHead:=make([]byte,dp.GetHeadLen())
		_,_=io.ReadFull(conn,binaryHead)
		msgHead,err:=dp.UnPackage(binaryHead)
		if msgHead.GetMessageLen()>0{
			msg:=msgHead.(*znet.Message)
			msg.Data=make([]byte,msg.GetMessageLen())
			_,_=io.ReadFull(conn,msg.Data)
			data:=&pb.SyncPid{}
			proto.Unmarshal(msg.GetMessageData(),data)
			fmt.Println("recv from server:msgID=",msg.GetMessageID(),"msg:", data.Pid)
		}

		time.Sleep(1*time.Second)




	}
}
