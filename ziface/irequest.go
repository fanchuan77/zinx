package ziface

/*
	获取Connetion连接对象 和 获取连接数据，封装到一个Request中
*/

type IRequest interface {
	//获取Connection
	GetConnection() IConnection

	//获取消息数据
	GetMsgData() []byte

	//获取消息ID
	GetMsgId() uint32
}
