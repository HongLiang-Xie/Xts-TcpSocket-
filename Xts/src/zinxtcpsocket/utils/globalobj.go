package utils

import (
	"awesomeProject1/src/zinxtcpsocket/zinxface"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

/*存储一切有关zinx框架的全局参数，供其他模块使用
一些参数是可以通过zinx.json由用户进行配置
 */
type GlobalObj struct {
	TCPServer zinxface.IServer //当前zinx全局的Server对象
	Host string //当前服务器主机地址
	TcpPort int //当前服务器主机监听的端口号
	Name string //当前服务器名称
	Version string //当前zinx服务器框架版本号
	MaxConn int //当前服务器允许的最大链接数
	MaxPackageSize uint32 //当前服务器允许的最大数据包
	WorkPoolSize uint32 //当前业务工作池的Goroutine数量
	MaxWorkTaskSize uint32 //每个worker对应的消息队列的任务数量最大值
}
/*
定义一个全局对外的Globalobj
 */
var GlobalObject*GlobalObj
/*
 提供一个init 方法 初始化当前GlobalObject
 */
func init(){
	GlobalObject=&GlobalObj{
		Name: "zinxServerApp",
		Version: "v0.9",
		TcpPort: 8080,
		Host: "192.168.2.87",
		MaxConn: 3000,
		MaxPackageSize: 4096,
		WorkPoolSize: 10,
		MaxWorkTaskSize: 10,
	}

	GlobalObject.Reload()
}
//从zinx.json去加载用于自定义的参数
func(g*GlobalObj)Reload(){
	//dir,_:=os.Getwd()
  data,err:=ioutil.ReadFile("/bin/conf/zinxtcpsocket.json")
  if err!=nil{
  	fmt.Println("conf/zinxtcpsocket.json error",err)
  }
  //将json文件数据解析到Globalobj中
  err=json.Unmarshal(data,&GlobalObject)
}




