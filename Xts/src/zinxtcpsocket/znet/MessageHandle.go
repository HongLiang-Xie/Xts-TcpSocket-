package znet

import (
	"awesomeProject1/src/zinxtcpsocket/utils"
	"awesomeProject1/src/zinxtcpsocket/zinxface"
	"fmt"
	"strconv"
)

/*消息处理的实现模块

 */
type MessageHandle struct {
	//存放每个MessageId 所对应的处理方法
	Apis map[uint32]zinxface.IRouter
	//负责worker取任务的消息队列
	TaskQueue [] chan zinxface.IRequest
	//业务工作worker池的worker数量
	WorkPoolSize uint32
}

//初始化一个MessageHandle
func NewMessageHandle() *MessageHandle {

	return &MessageHandle{
		Apis: make(map[uint32]zinxface.IRouter),
		TaskQueue: make([]chan zinxface.IRequest,utils.GlobalObject.MaxWorkTaskSize),//从全局配置中获取
		WorkPoolSize: utils.GlobalObject.WorkPoolSize,//从全局配置中获取
	}

}

//调度/执行对应的Router消息处理方法
func (mh *MessageHandle) DoMessageHandler(request zinxface.IRequest) {

	Handler, it := mh.Apis[request.GetDataID()]
	if !it {

		fmt.Println("api msgID=", request.GetDataID(), "NOT FOUND! Need  Register")
		return

	}
	//根据MessageID 调度Router执行相应业务
	Handler.PreHandle(request)
	Handler.Handler(request)
	Handler.PostHandle(request)

}

//为消息添加具体逻辑
func (mh *MessageHandle) AddRouter(msgId uint32, router zinxface.IRouter) {

	if _, it := mh.Apis[msgId]; it {

		panic("repeat api MsgID=" + strconv.Itoa(int(msgId)))
	}
	//添加MsgId 与router的绑定关系
	mh.Apis[msgId] = router
	fmt.Println("Add api msgID=", msgId, "success")
}
//启动worker工作池（开启工作池只能开启一次）
func (mh *MessageHandle)StartWorkerPool(){
		//根据WorkerPoolSize 分别开启worker 每一个worker都用一个go来承载
	for i:=0;i<int(mh.WorkPoolSize);i++{

		mh.TaskQueue[i]=make(chan zinxface.IRequest,utils.GlobalObject.MaxWorkTaskSize)
		go mh.StartOneWorker(i,mh.TaskQueue[i])
	}

}

//启动一个worker工作流程

func (mh *MessageHandle)StartOneWorker(workID int, taskQueue chan zinxface.IRequest){
	fmt.Println("Worker ID=",workID,"is started....")
	//阻塞等待对应消息队列的消息
	for{
		select {
		//如果有消息过来,出列的就是一个客户端的Request，执行当前router所绑定的业务
			case request:=<-taskQueue:
				mh.DoMessageHandler(request)

		}

	}

}
//将消息交给TaskQueue，由worker进行处理
func (mh *MessageHandle) SendMsgToTaskQueue(reuqest zinxface.IRequest){
		//1 将消息平均分配给不同的worker
			//根据客户端建立的	ConnID来进行分配
		workerID:=reuqest.GetConnection().GetConnID()%mh.WorkPoolSize
		fmt.Println("Add ConnID=",reuqest.GetConnection().GetConnID(),
		"request msgID=",reuqest.GetDataID(),
		"to worker=" ,workerID)
		//2 将消息发送给对应的worker的TaskQueue即可
		mh.TaskQueue[workerID]<-reuqest
}