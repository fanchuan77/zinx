package ziface

/*
	获取Connection连接对象 和 获取连接数据，封装到一个Request中
*/

type IRequest interface {
	// GetConnection 获取Connection
	GetConnection() IConnection

	// GetMsgData 获取消息数据
	GetMsgData() []byte

	// GetMsgId 获取消息ID
	GetMsgId() uint32
}
