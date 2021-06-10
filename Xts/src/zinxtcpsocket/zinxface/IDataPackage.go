package zinxface

/* 封包，拆包 模块
 直接面向TCP连接中的数据流，用于处理TCP粘包问题
*/
type IDataPackage interface {
	//获取包的长度方法
	GetHeadLen()uint32
	//封包方法
	Package(msg IMessage) ([]byte,error)
	//拆包方法
	UnPackage([]byte)(IMessage,error)
}