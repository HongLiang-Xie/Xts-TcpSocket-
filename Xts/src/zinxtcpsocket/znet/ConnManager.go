package znet

import (
	"awesomeProject1/src/zinxtcpsocket/zinxface"
	"errors"
	"fmt"
	"sync"
)

/*链接管理模块的实现层
 */
type ConnManager struct {

	connections map[uint32] zinxface.IConnection
	connLock sync.RWMutex //保护连接集合的读写锁

}

//初始化
func NewConnManager() zinxface.IConnManger{

	return &ConnManager{
		connections: make(map[uint32]zinxface.IConnection),
	}
}



//添加链接
func (cm*ConnManager)Add(conn zinxface.IConnection){
	//添加保护读写锁
	cm.connLock.Lock()
	defer cm.connLock.Unlock()
	//添加连接
	cm.connections[conn.GetConnID()]=conn
	fmt.Println("connection connID=" ,conn.GetConnID(),"add to Connmanager successfully! conn num=",cm.Len())

}
//删除链接
func (cm*ConnManager)Remove(conn zinxface.IConnection){
	//添加保护读写锁
	cm.connLock.Lock()
	defer cm.connLock.Unlock()
	//删除连接
	delete(cm.connections,conn.GetConnID())
	fmt.Println("connection connID=" ,conn.GetConnID(),"remove to Connmanager successfully! conn num=",cm.Len())
}
//根据ConnID获取链接
func (cm*ConnManager)Get(connID uint32)(zinxface.IConnection,error){

	//添加保护读写锁
	cm.connLock.Lock()
	defer cm.connLock.Unlock()
	conn,err:=cm.connections[connID]
	if err{

		return nil,errors.New("connection NOT FOUND")
	}
	return conn,nil

}
//得到链接总数
func (cm*ConnManager)Len()int{

	return len(cm.connections)

}
//清楚并终止所有链接
func (cm*ConnManager)ClearConn(){
	//添加保护读写锁
	cm.connLock.Lock()
	defer cm.connLock.Unlock()
	for connID,conn:=range cm.connections{

		conn.Stop()
		delete(cm.connections,connID)

	}
	fmt.Println("clear all connection successfully")
}