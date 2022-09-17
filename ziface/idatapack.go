package ziface

/*
	封包、拆包的模块
	解决TCP的粘包问题
*/
type IDataPack interface {
	// GetHeadLen 获取包的头的长度
	GetHeadLen() uint32
	// Pack 封包方法
	Pack(msg IMessage) ([]byte, error)
	// Unpack 拆包方法
	Unpack(binaryData []byte) (IMessage, error)
}
