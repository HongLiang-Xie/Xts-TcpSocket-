package znet

import (
	"awesomeProject1/src/zinxtcpsocket/utils"
	"awesomeProject1/src/zinxtcpsocket/zinxface"
	"bytes"
	"encoding/binary"
	"errors"
)

/* 拆包 封包的方法*/
type DataPackage struct {
	
}

func NewDataPackage() *DataPackage {

	return &DataPackage{}
}

//获取包的长度方法
func (dp*DataPackage)GetHeadLen()uint32{
	//Datalen uint32(4字节)+ID uint32(4字节)
	return 8
}
//封包方法
func (dp*DataPackage)Package(msg zinxface.IMessage) ([]byte,error){
	//创建一个存放bytes字节的缓冲
	dataBuff:=bytes.NewBuffer([]byte{})
	//将datalen 写进databuff中
	if err:=binary.Write(dataBuff,binary.LittleEndian,msg.GetMessageLen());err!=nil{

		return nil ,err
	}
	//将msgID 写进databuff中
	if err:=binary.Write(dataBuff,binary.LittleEndian,msg.GetMessageID());err!=nil{

		return nil ,err
	}
	//将data数据写进databuff中
	if err:=binary.Write(dataBuff,binary.LittleEndian,msg.GetMessageData());err!=nil{

		return nil ,err
	}
	return dataBuff.Bytes(),nil
}
//拆包方法 (将包的Head信息读出来)之后再根据Head信息里的data长度 再进行一起读取
func (dp*DataPackage)UnPackage(binaryData []byte)(zinxface.IMessage,error){
	//创建一个从输入二进制数据的ioReader
	dataBuff:=bytes.NewReader(binaryData)
	//解压Head信息,得到dataLen和msgID
	msg:=&Message{}
	//读dataLen
	if err:=binary.Read(dataBuff,binary.LittleEndian,&msg.DataLen);err!=nil{
		return nil ,err
	}
	//读msgID
	if err:=binary.Read(dataBuff,binary.LittleEndian,&msg.ID);err!=nil{
		return nil ,err
	}
	//判断datalen是否超出了允许的最大包长度
	if msg.DataLen>utils.GlobalObject.MaxPackageSize{

		return nil ,errors.New("too large msg data recv")

	}
	return msg,nil

}

