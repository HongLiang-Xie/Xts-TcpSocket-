package apis

import (
	"awesomeProject1/src/db"
	"awesomeProject1/src/mmogame/pb"
	"awesomeProject1/src/zinxtcpsocket/zinxface"
	"awesomeProject1/src/zinxtcpsocket/znet"
	"fmt"
	"google.golang.org/protobuf/proto"
	"strings"
)

type Login struct {
	znet.BaseRouter
}

func (p* Login) Handler(request zinxface.IRequest)  {
	//解析用户传进来的Proto协议
	protoMessage:=&pb.Login{}
	err:=proto.Unmarshal(request.GetData(),protoMessage)
	if err!=nil{
		fmt.Println("Position Unmarshal err ",err)
		return
	}
	//获取数据库对象
	mysqldb:=db.MySqlPoos.GetMySqlManger()
	//查询是否正确
	var person [] string
	switch protoMessage.Type {

	case 1:

		mysqldb.Select(&person,db.FindPassword,protoMessage.Userid)

		if person==nil{
			ReturnMessage:=&pb.Login{
				Type: 1,
				ReturnType: 3,
			}
			_data,_:=proto.Marshal(ReturnMessage)
			request.GetConnection().SendMsg(203,_data)
			return
		}

		if strings.Compare(person[0],protoMessage.Password)==0 {
			ReturnMessage:=&pb.Login{
				Type: 1,
				ReturnType: 1,
			}
			_data,_:=proto.Marshal(ReturnMessage)
			request.GetConnection().SendMsg(203,_data)

		}

	case 2:


	}




}
