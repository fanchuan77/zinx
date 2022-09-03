package ziface

//消息管理抽象层
type IMsgHandle interface {
	//调度/执行对应的Router消息处理方法
	DoMsgHandler(request IRequest)

	//为消息添加具体处理逻辑
	AddRouter(msgId uint32, router IRouter)

	//启动 Worker工作池
	StartWokerPool()

	//启动一个 Worker工作流程
	StartOneWorker(wid int, taskQueue chan IRequest)

	//将request消息传入TaskQueue
	SendMsgToTaskQueue(request IRequest)
}
