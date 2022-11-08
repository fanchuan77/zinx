package znet

import (
	"fmt"
	"lib/zinx/utils"
	"lib/zinx/ziface"
	"strconv"
)

// MsgHandle 消息管理模块
type MsgHandle struct {
	//存放所有 MsgId 与 Router 的对应关系
	Apis map[uint32]ziface.IRouter
	//业务工作池的 worker总量
	WorkerPoolSize uint32
	//与worker对接的消息队列
	TaskQueue []chan ziface.IRequest
}

func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		Apis:           make(map[uint32]ziface.IRouter),
		WorkerPoolSize: utils.GlobalObject.WorkerPoolSize, //从全局配置中获取
		TaskQueue:      make([]chan ziface.IRequest, utils.GlobalObject.WorkerPoolSize),
	}
}

// DoMsgHandler 调度/执行对应的Router消息处理方法
func (mh *MsgHandle) DoMsgHandler(request ziface.IRequest) {
	//handler继承自BaseRouter
	handler, ok := mh.Apis[request.GetMsgId()]
	if !ok {
		fmt.Println("api MsgId =", request.GetMsgId(), "is NOT FOUND! Need Register!")
		return
	}
	handler.PreHandle(request)
	handler.Handle(request)
	handler.PostHandle(request)
}

// AddRouter 为消息添加具体处理逻辑
func (mh *MsgHandle) AddRouter(msgId uint32, router ziface.IRouter) {
	if _, ok := mh.Apis[msgId]; ok {
		//msgId已注册
		panic("repeat api,MsgId =" + strconv.Itoa(int(msgId)))
	}
	//添加 msgId与Router的绑定关系
	mh.Apis[msgId] = router
	fmt.Println("Add api MsgId =", msgId, "succ!")
}

// StartWokerPool 启动 Worker工作池
func (mh *MsgHandle) StartWokerPool() {
	//根据WorkerPoolSize分别开启Worker,每个 Worker用一个go来承载
	for i := 0; i < int(mh.WorkerPoolSize); i++ {
		//当前 Worker对应的channel，第i个 Worker对应第i个channel
		//每个channel中包含最大任务数量
		mh.TaskQueue[i] = make(chan ziface.IRequest, utils.GlobalObject.MaxWorkerTaskLen)
		//启动当前 Worker,阻塞等待消息从channel传入
		go mh.StartOneWorker(i, mh.TaskQueue[i])
	}
}

// StartOneWorker 启动一个 Worker工作流程,等待消息队列传出消息
func (mh *MsgHandle) StartOneWorker(wid int, taskQueue chan ziface.IRequest) {
	fmt.Println("Worker ID:", wid, "is started...")

	//阻塞等待
	for {
		select {
		//从channel中拿到request，执行request绑定的 Router方法
		case request := <-taskQueue:
			mh.DoMsgHandler(request)
		}
	}

}

// SendMsgToTaskQueue 将request消息传入TaskQueue
func (mh *MsgHandle) SendMsgToTaskQueue(request ziface.IRequest) {
	//平均分配消息传入消息队列
	//根据客户端 ConnID分配channel
	WorkerID := request.GetConnection().GetConnID() % mh.WorkerPoolSize
	fmt.Println("Add ConnID:", request.GetConnection().GetConnID(),
		"and request MsgID:", request.GetMsgId(),
		"to WorkerID:", WorkerID)
	//传入消息
	mh.TaskQueue[WorkerID] <- request
}
