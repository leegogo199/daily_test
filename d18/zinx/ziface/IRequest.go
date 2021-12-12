package ziface



//request 请求封装。
//实际上是把客户端请求的连接信息，和请求的数据包装到了一个Request中
//将连接和数据绑定在一起。

//
type IRequest interface{
	//得到当前链接
	GetConnection() IConnection
	//得到请求的消息数据
	GetData() []byte
	//得到请求的消息ID
	GetMsgID() uint32
	//
	//
	//
	//
}