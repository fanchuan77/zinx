package ziface

// IConnManager 连接管理模块抽象层
type IConnManager interface {
	// Add 增加连接
	Add(conn IConnection)

	// Remove 删除连接
	Remove(conn IConnection)

	// GetConnection 查询连接
	GetConnection(connID uint32) (IConnection, error)

	// Len 获得连接总数
	Len() uint32

	// Clear 清除所有连接并终止
	Clear()
}
