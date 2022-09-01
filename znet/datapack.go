package znet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"zinx/src/zinx/utils"
	"zinx/src/zinx/ziface"
)

/*
	针对TCP粘包问题对 Message进行封装
*/
type Datapack struct {
}

//初始化一个封包 拆包的实例对象
func NewDataPack() ziface.IDataPack {
	return &Datapack{}
}

/*
	获取包头的长度
	DataLen uint32(4字节) + ID uint32(4字节)
*/
func (dp *Datapack) GetHeadLen() uint32 {
	return 8
}

//封包方法
func (dp *Datapack) Pack(msg ziface.IMessage) ([]byte, error) {
	//得到一个存放 Bytes字节的缓冲
	dataBuff := bytes.NewBuffer([]byte{})

	//将MsgLen写进dataBuff	小端:LittleEndian
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgLen()); err != nil {
		return nil, err
	}

	//将ID写进dataBuff
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgId()); err != nil {
		return nil, err
	}

	//将Data写进dataBuff
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgData()); err != nil {
		return nil, err
	}

	return dataBuff.Bytes(), nil
}

//拆包方法	(读取包里的 Head信息,之后再根据 Head信息里的 Data长度再次读取)
func (dp *Datapack) Unpack(binaryData []byte) (ziface.IMessage, error) {
	//创建一个从二进制输入数据中读取的ioReader
	dataBuff := bytes.NewReader(binaryData)

	//只解压 Head得到 MsgLen和 ID
	msg := &Message{}
	//读取 MsgLen
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.MsgLen); err != nil {
		return nil, err
	}
	//读取 ID
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.ID); err != nil {
		return nil, err
	}

	//判断 MsgLen是否超出允许长度
	if msg.MsgLen > utils.GlobalObject.MaxPackageSize && utils.GlobalObject.MaxPackageSize > 0 {
		return nil, errors.New("too large Msg Data... ")
	}

	return msg, nil
}
