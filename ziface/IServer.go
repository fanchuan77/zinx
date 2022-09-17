package ziface

//定义一个服务器接口
type IServer interface {
	// Start 启动服务器
	Start()
	// Serve 运行服务器
	Serve()
	// Stop 关闭服务器
	Stop()
	// AddRouter 路由功能，给当前Server添加一个Router
	AddRouter(msgId uint32, router IRouter)
	// GetConnMgr 获取Server的连接管理器
	GetConnMgr() IConnManager
	// SetOnConnStart 注册OnConnStart 钩子函数的方法
	SetOnConnStart(func(IConnection))
	// SetOnConnStop 注册OnConnStop 钩子函数的方法
	SetOnConnStop(func(IConnection))
	// CallOnConnStart 调用OnConnStart 钩子函数的方法
	CallOnConnStart(IConnection)
	// CallOnConnStop 调用OnConnStop 钩子函数的方法
	CallOnConnStop(IConnection)
}
