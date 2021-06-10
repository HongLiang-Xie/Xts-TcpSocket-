package znet

type Message struct {
	ID      uint32 //消息的id
	DataLen uint32 //消息的长度
	Data    []byte //消息
}

//创建一个Message消息包
func NewMessagePackage(id uint32, data []byte) *Message {
	return &Message{
		ID:      id,
		DataLen: uint32(len(data)),
		Data:    data,
	}
}

//获取消息的ID
func (m *Message) GetMessageID() uint32 {
	return m.ID
}

//获取消息的长度
func (m *Message) GetMessageLen() uint32 {
	return m.DataLen
}

//获取消息的内容
func (m *Message) GetMessageData() []byte {
	return m.Data
}

//设置消息的ID
func (m *Message) SetMessageID(id uint32) {
	m.ID=id
}

//设置消息的长度
func (m *Message) SetMessageLen(len uint32) {
	m.DataLen=len
}

//设置消息的内容
func (m *Message) SetMessageData(data []byte) {
  m.Data=data
}
