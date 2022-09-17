package znet

import (
	"errors"
	"fmt"
	"lib/zinx/ziface"
	"sync"
)

type ConnManager struct {
	//管理的连接集合
	connections map[uint32]ziface.IConnection
	//保护连接集合的读写锁
	connLock sync.RWMutex
}

func NewConnManager() *ConnManager {
	return &ConnManager{
		connections: make(map[uint32]ziface.IConnection),
	}
}

// Add 增加连接
func (cm *ConnManager) Add(conn ziface.IConnection) {
	//保护共享资源,加写锁
	cm.connLock.Lock()
	defer cm.connLock.Unlock()

	//将conn加入map
	cm.connections[conn.GetConnID()] = conn

	fmt.Println("Connection connID =", conn.GetConnID(), "Add succ!! connLen =", cm.Len())
}

// Remove 删除连接
func (cm *ConnManager) Remove(conn ziface.IConnection) {
	//保护共享资源,加写锁
	cm.connLock.Lock()
	defer cm.connLock.Unlock()

	//删除连接
	delete(cm.connections, conn.GetConnID())

	fmt.Println("Connection connID =", conn.GetConnID(), "Remove succ!! connLen =", cm.Len())
}

// GetConnection 查询连接
func (cm *ConnManager) GetConnection(connID uint32) (ziface.IConnection, error) {
	//保护共享资源,加读锁
	cm.connLock.RLock()
	defer cm.connLock.RUnlock()

	if conn, ok := cm.connections[connID]; ok {
		//找到了
		return conn, nil
	} else {
		return nil, errors.New("Connection Not FOUND")
	}
}

// Len 获得连接总数
func (cm *ConnManager) Len() uint32 {
	return uint32(len(cm.connections))
}

// Clear 清除所有连接并终止
func (cm *ConnManager) Clear() {
	//保护共享资源,加写锁
	cm.connLock.Lock()
	defer cm.connLock.Unlock()

	for connID, conn := range cm.connections {
		//停止
		err := conn.Stop()
		if err != nil {
			return
		}
		//删除
		delete(cm.connections, connID)
	}
	fmt.Println("clear all Connection succ!! connLen =", cm.Len())
}
