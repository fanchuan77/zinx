package ziface

/*
	封包、拆包的模块
	解决TCP的粘包问题
*/
type IDataPack interface {
	//获取包的头的长度
	GetHeadLen() uint32
	//封包方法
	Pack(msg IMessage) ([]byte, error)
	//拆包方法
	Unpack(binaryData []byte) (IMessage, error)
}
