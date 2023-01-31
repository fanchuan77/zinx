package znet

/*
将请求的消息封装到 Message中
*/
type Message struct {
	ID     uint32 //消息ID
	MsgLen uint32 //消息内容长度
	Data   []byte //消息内容
}

// NewMsgPackage 创建 Message包对象
func NewMsgPackage(msgId uint32, data []byte) *Message {
	return &Message{
		ID:     msgId,
		MsgLen: uint32(len(data)),
		Data:   data,
	}
}

// GetMsgId 获取消息Id
func (m *Message) GetMsgId() uint32 {
	return m.ID

}

// GetMsgLen 获取消息长度
func (m *Message) GetMsgLen() uint32 {
	return m.MsgLen
}

// GetMsgData 获取消息内容
func (m *Message) GetMsgData() []byte {
	return m.Data
}

// SetMsgId 设置消息Id
func (m *Message) SetMsgId(ID uint32) {
	m.ID = ID
}

// SetMsgLen 设置消息长度
func (m *Message) SetMsgLen(MsgLen uint32) {
	m.MsgLen = MsgLen
}

// SetMsgData 设置消息内容
func (m *Message) SetMsgData(Data []byte) {
	m.Data = Data
}
