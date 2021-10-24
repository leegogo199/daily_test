package ziface
//定义一个解决TCP粘包问题的封包拆包的模块
//针对message进行TLV格式的封装
//针对message进行TLV的拆包
//先读取固定长度的head包括 内容长度，消息id，再读取消息内容。
//写的时候，先写长度，再写id，再写内容。
//
type IDataPack interface{
	//获取包的长度方法
	GetHeadLen() uint32
	//封包方法
	Pack(msg ziface.IMessage)([]byte,error)
	//拆包方法
	Unpack([]byte)(ziface.IMessage,error)


}