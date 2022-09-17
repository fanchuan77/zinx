package ziface

// IMsgHandle 消息管理抽象层
type IMsgHandle interface {
	// DoMsgHandler 调度执行对应的Router消息处理方法
	DoMsgHandler(request IRequest)

	// AddRouter 为消息添加具体处理逻辑
	AddRouter(msgId uint32, router IRouter)

	// StartWokerPool 启动 Worker工作池
	StartWokerPool()

	// StartOneWorker 启动一个 Worker工作流程
	StartOneWorker(wid int, taskQueue chan IRequest)

	// SendMsgToTaskQueue 将request消息传入TaskQueue
	SendMsgToTaskQueue(request IRequest)
}
