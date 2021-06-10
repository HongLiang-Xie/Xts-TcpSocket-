package zinxface
/*
链接管理模块抽象层
 */
type IConnManger interface {
	//添加链接
	Add(conn IConnection)
	//删除链接
	Remove(conn IConnection)
	//根据ConnID获取链接
	Get(connID uint32)(IConnection,error)
	//得到链接总数
	Len()int
	//清楚并终止所有链接
	ClearConn()

}