package ziface

//定义一个服务器接口
type IServer interface {
	//启动服务器
	Start()
	//运行服务器
	Serve()
	//关闭服务器
	Stop()
	//路由功能，给当前Server添加一个Router
	AddRouter(msgId uint32, router IRouter)
	//获取Server的连接管理器
	GetConnMgr() IConnManager
}
