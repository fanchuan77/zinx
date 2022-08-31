package ziface

type IMessage interface {
	//获取消息Id
	GetMsgId() uint32

	//获取消息长度
	GetMsgLen() uint32

	//获取消息内容
	GetMsgData() []byte

	//设置消息Id
	SetMsgId(Id uint32)

	//设置消息长度
	SetMsgLen(MsgLen uint32)

	//设置消息内容
	SetMsgData(Data []byte)
}
