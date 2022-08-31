package znet

/*
	将请求的消息封装到Message中
*/
type Message struct {
	Id     uint32 //消息ID
	MsgLen uint32 //消息长度
	Data   []byte //消息内容
}

//获取消息Id
func (m *Message) GetMsgId() uint32 {
	return m.Id

}

//获取消息长度
func (m *Message) GetMsgLen() uint32 {
	return m.MsgLen
}

//获取消息内容
func (m *Message) GetMsgData() []byte {
	return m.Data
}

//设置消息Id
func (m *Message) SetMsgId(Id uint32) {
	m.Id = Id
}

//设置消息长度
func (m *Message) SetMsgLen(MsgLen uint32) {
	m.MsgLen = MsgLen
}

//设置消息内容
func (m *Message) SetMsgData(Data []byte) {
	m.Data = Data
}
