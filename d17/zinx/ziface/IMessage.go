package ziface
//定义一个消息的结构，属性，消息的ID，消息的内容，消息的长度


//将消息封装机制集成到zinx框架中，修改链接读取数据的机制，将之前的单纯的
//读取byte改成拆包形式的读取。
//按照TLV读取。

//将message 添加到request属性中。
//给链接提供一个发包机制，将发送的信息进行打包，再发送
//
//


//将请求的消息封装到一个Message中，定义抽象的接口。
type IMessage interface{
	//获取消息的ID
	GetMsgId() uint32
	//获取消息的长度
	GetMsgLen() uint32
	//获取消息的内容
	GetData() []byte
	//设置消息的ID
	SetMsgId(uint32)
	//设置消息的长度
	SetMsgLen(uint32)
	//设置消息的内容
	SetData([]byte)
}
//
//
