package znet

import (
	"fmt"
	"strconv"
	"zinx/src/zinx/ziface"
)

//消息管理模块
type MsgHandle struct {
	//存放所有 MsgId 与 Router 的对应关系
	Apis map[uint32]ziface.IRouter
}

func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		Apis: make(map[uint32]ziface.IRouter),
	}
}

//调度/执行对应的Router消息处理方法
func (mh *MsgHandle) DoMsgHandler(request ziface.IRequest) {
	handler, ok := mh.Apis[request.GetMsgId()]
	if !ok {
		fmt.Println("api MsgId =", request.GetMsgId(), "is NOT FOUND! Need Register!")
	}
	handler.PreHandle(request)
	handler.Handle(request)
	handler.PostHandle(request)
}

//为消息添加具体处理逻辑
func (mh *MsgHandle) AddRouter(msgId uint32, router ziface.IRouter) {
	if _, ok := mh.Apis[msgId]; ok {
		//msgId已注册
		panic("repeat api,MsgId =" + strconv.Itoa(int(msgId)))
	}
	//添加 msgId与Router的绑定关系
	mh.Apis[msgId] = router
	fmt.Println("Add api MsgId =", msgId, "succ!")
}
