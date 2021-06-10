package zinxface
type IMessageHandle interface {
	//调度/执行对应的Router消息处理方法
	DoMessageHandler(request IRequest)
	//为消息添加具体逻辑
	AddRouter(msgId uint32 ,router IRouter)
	//启动worker工作池（开启工作池只能开启一次）
	StartWorkerPool()
	//启动一个worker工作流程
	//StartOneWorker(workID int, taskQueue chan IRequest)
	//将消息交给TaskQueue，由worker进行处理
	SendMsgToTaskQueue(reuqest IRequest)
}
