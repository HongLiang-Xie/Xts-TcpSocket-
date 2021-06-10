package zinxface

/*
将请求的消息封装到一个Message中 定义一个Message一个抽象接口
 */
type IMessage interface {

	//获取消息的ID
	GetMessageID() uint32
	//获取消息的长度
	GetMessageLen() uint32
	//获取消息的内容
	GetMessageData() []byte

	//设置消息的ID
	SetMessageID(id uint32)
	//设置消息的长度
	SetMessageLen(len uint32)
	//设置消息的内容
	SetMessageData(data []byte)
}