package znet

import (
	"sync"
	"zinx/src/zinx/ziface"
)

type ConnManager struct{
	//管理的连接集合
	connections map[uint32] ziface.IConnection
	//保护连接集合的读写锁
	connLock sync.RWMutex

}
func NewConnManager() *ConnManager {
	return &ConnManager{
		connections: make[map[uint32] ziface.IConnection],
		connLock: 
	}
}


//增加连接
func (cm *ConnManager) Add(conn ziface.IConnection){

}

//删除连接
func (cm *ConnManager) Remove(conn ziface.IConnection){

}

//查询连接
func (cm *ConnManager) GetConnection(connID uint32) (ziface.IConnection, error){

}
 
//获得连接总数
func (cm *ConnManager) Len() uint32{
	
}

//清除所有连接并终止
func (cm *ConnManager) Clear(){

}