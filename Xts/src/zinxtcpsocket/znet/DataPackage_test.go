package znet

import (
	"fmt"
	"io"
	"net"
	"testing"
)

//只负责测试DataPack拆包 封包的单元测试
func TestDataPackage(t *testing.T) {
	//模拟服务器
	Listener,err:=net.Listen("tcp","localhost:8080")
	if err!=nil{
		fmt.Println("server error:",err)
		return
	}
	go func() {

		for{
			conn,err:=Listener.Accept()
			if err!=nil{
				fmt.Println("server accept error")

			}
			//处理客户端请求
			go func(conn net.Conn) {
				//定义一个拆包对象
				dp:=NewDataPackage()
				for {
					//-----拆包过程------
					//1 第一次从conn读，把包的Head读出来
					headData:=make([]byte,dp.GetHeadLen())
					_,err=io.ReadFull(conn,headData)
					if err!=nil{
						fmt.Println("read head error")
						break
					}

					msgHead ,err:=dp.UnPackage(headData)
					if err!=nil{
					fmt.Println("server unpackage err",err)
						return
					}
					if msgHead.GetMessageLen()>0 {
						//msg有数据，需要进行二次读取
						//2 第二次从conn读，把包的dataLen读出来
						msg:=msgHead.(*Message)
						msg.Data=make([]byte,msg.DataLen)
						_,err:=io.ReadFull(conn,msg.Data)
						if err!=nil{
							fmt.Println("server unpackage data error",err)
							return
						}

						fmt.Println("Recv MsgID",msg.ID,"MsgDataLen",msg.DataLen,"MsgData",string(msg.Data))
					}


				}
			}(conn)

		}
	}()
	//模拟客户端
	conn,err:=net.Dial("tcp","localhost:8080")
	if err!=nil{
		return
	}
	//创建一个封包对象
	dp:=NewDataPackage()
	//模拟粘包过程，封装两个msg一同发送
	//封装第一个
	msg1:=&Message{
		ID: 1,
		DataLen: 5,
		Data: []byte("hello"),
	}
	sendData1,err:=dp.Package(msg1)
	if err!=nil{

		return
	}
	msg2:=&Message{
		ID: 2,
		DataLen: 6,
		Data: []byte(",world"),
	}
	sendData2,err:=dp.Package(msg2)
	if err!=nil{

		return
	}
	sendData1=append(sendData1,sendData2...)
	conn.Write(sendData1)
	select {

	}
}
