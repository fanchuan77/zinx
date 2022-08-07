package ziface

//定义一个服务器接口
type IServer interface {
	//启动服务器
	Start()
	//运行服务器
	Serve()
	//关闭服务器
	Stop()
}
