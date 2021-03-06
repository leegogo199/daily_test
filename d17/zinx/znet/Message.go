package znet

type Message struct{
	//消息的ID
	Id uint32
	//消息的长度
	DataLen uint32
	//消息的内容
	Data []byte

}
//创建一个MEssage小溪包
func NewMsgPackage(id uint32,data []byte)*Message{
	return &Message{
		Id:id,
		DataLen: uint32(len(data)),
		Data:data,
	}
}

//获取消息的ID
func (m *Message)GetMsgId() uint32{
	return m.Id

}
//获取消息的长度
func (m *Message)GetMsgLen() uint32{
	return m.DataLen
}
//获取消息的内容
func (m *Message)GetData() []byte{
	return m.Data

}
//设置消息的ID
func (m *Message)SetMsgId(id uint32){
	m.Id=id

}
//设置消息的长度
func (m *Message)SetMsgLen(len uint32){
	m.DataLen=len

}
//设置消息的内容
func (m *Message)SetData(data []byte){
	m.Data=data
}