package znet

import (
	"bytes"
	"dailytest/daily_test/d18/zinx/utils"
	"dailytest/daily_test/d18/zinx/ziface"
	"encoding/binary"
	"errors"
)
// 封包，拆包模块，直接面向tcp链接中的数据流，处理tcp粘包问题。
type DataPack struct{

}
//拆包封包实例的一个初始化方法
func NewDataPack() *DataPack{
	return &DataPack{}
}
//

//获取包的长度方法
func(dp *DataPack)GetHeadLen() uint32{
	//Datalen uint32(4字节）+ID uint32(4字节)
	return 8
}
//封包方法
//datalen|msgID|data|
func(dp *DataPack)Pack(msg ziface.IMessage)([]byte,error){
	//创建一个存放bytes字节的缓存
	DataBuff:=bytes.NewBuffer([]byte{})
	//将datalen写进databuff中
	if err:=binary.Write(DataBuff,binary.LittleEndian,msg.GetMsgLen());err!=nil{
		return nil,err
	}
	//将MsgId写进databuff中
	if err:=binary.Write(DataBuff,binary.LittleEndian,msg.GetMsgId());err!=nil{
		return nil,err
	}
	//将data数据 写进databuff中
	if err:=binary.Write(DataBuff,binary.LittleEndian,msg.GetData());err!=nil{
		return nil,err
	}
	return DataBuff.Bytes(),nil
}
//拆包方法 (将包的head信息读出来)之后再根据head信息里的data长度，在进行一次读。
func (dp *DataPack)Unpack(binaryData []byte)(ziface.IMessage,error){
	//创建一个从输入二进制数据的ioReader
	dataBuff:=bytes.NewReader(binaryData)
	//只解压head信息，得到datalen，和MsgID
	msg:=&Message{}

		//读dataLen
	if err:=binary.Read(dataBuff,binary.LittleEndian,&msg.DataLen);err!=nil{
		return nil,err
	}
		//读MsgId
	if err:=binary.Read(dataBuff,binary.LittleEndian,&msg.Id);err!=nil{
		return nil,err
	}
	//判断datalen 是否已经超出了我们允许的最大包长度。
	if (utils.GlobalObject.MaxPackageSize>0&&msg.DataLen>utils.GlobalObject.MaxPackageSize){
		return nil,errors.New("too large msg data recv!")
	}
	return msg,nil

}