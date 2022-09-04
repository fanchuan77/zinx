package ziface

/*
	连接管理模块抽象层
*/
type IConnManager interface {
	//增加连接
	Add(conn IConnection)

	//删除连接
	Remove(conn IConnection)

	//查询连接
	GetConnection(connID uint32) (IConnection, error)

	//获得连接总数
	Len() uint32

	//清除所有连接并终止
	Clear()
}
