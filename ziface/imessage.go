package ziface

type IMessage interface {
	// GetMsgId 获取消息Id
	GetMsgId() uint32

	// GetMsgLen 获取消息长度
	GetMsgLen() uint32

	// GetMsgData 获取消息内容
	GetMsgData() []byte

	// SetMsgId 设置消息ID
	SetMsgId(ID uint32)

	// SetMsgLen 设置消息长度
	SetMsgLen(MsgLen uint32)

	// SetMsgData 设置消息内容
	SetMsgData(Data []byte)
}
