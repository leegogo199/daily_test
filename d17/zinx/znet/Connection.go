package znet

import (
	"dailytest/daily_test/d17/zinx/ziface"
	"errors"
	"fmt"
	"io"
	"net"
)

// 链接模块
type Connection struct{
	//获取链接套接字
	Conn *net.TCPConn
	//链接的ID
	ConnID uint32
	//当前的链接状态
	isClosed bool
	//当前链接所绑定的处理方法API
	//handleAPI ziface.HandleFunc
	//告知当前链接已经退出的/停止 channel
	ExitChan chan bool
	// 该链接处理的方法Router
	Router ziface.IRouter

}
//初始化链接模块的方法
func NewConnection(conn *net.TCPConn,connID uint32,router ziface.IRouter)*Connection{
	c:=&Connection{
		Conn:conn,
		ConnID: connID,
		Router: router,
		isClosed: false,
		ExitChan: make(chan bool,1),
	}
	return c
}
//链接的读业务方法
func(c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running...")
	defer fmt.Println("connID=",c.ConnID," Reader is exit,remote addr is ",c.GetRemoteAddr().String())
	defer c.Stop()
	for {
		//读取客户端的数据到buf中，最大512字节


		//buf:=make([]byte,utils.GlobalObject.MaxPackageSize)
		//_,err:=c.Conn.Read(buf)
		//if err!=nil{
		//	fmt.Println("recv buf err",err)
		//	continue
		//}
		// 调用当前链接绑定的HandleAPI
		//if err:=c.handleAPI(c.Conn,buf,cnt);err!=nil{
		//	fmt.Println("ConnID",c.ConnID,"handle is error",err)
		//	break
		//}

		//创建一个拆包解包对象
		dp:=NewDataPack()

		//读取客户端的Msg Head 二进制流 8个字节
		headData:=make([]byte,dp.GetHeadLen())
		//拆包，得到msgID和 msgDatalen放在msg消息中
		if _,err:=io.ReadFull(c.GetTCPConnection(),headData);err!=nil{
			fmt.Println("read msg head error",err)
			return
		}
		msg,err:=dp.Unpack(headData)
		if err!=nil{
			fmt.Println("unpack error",err)
			break
		}
		//根据datalen，再次读去 data，放在msg.Data中。
		var data []byte
		if msg.GetMsgLen()>0{
			data=make([]byte,msg.GetMsgLen())
			if _,err:=io.ReadFull(c.GetTCPConnection(),data);err!=nil{
				fmt.Println("read msg data error",err)
				break
			}

		}
		msg.SetData(data)



		//
		//

		//得到当前conn数据的request请求数据
		req:=Request{
			conn:c,
			msg:msg,
		}
		//调用路由处理数据。

		//从路由中，找到注册绑定的Conn对应的router。
		//执行注册的路由方法。
		go func(request ziface.IRequest) {
			c.Router.BeforeHandle(request)
			c.Router.Handle(request)
			c.Router.AfterHandle(request)
		}(&req)


	}

}


//启动链接，让当前的链接准备开始工作
func (c *Connection ) Start(){
	fmt.Println("Conn Start()..ConnID=",c.ConnID)
	//启动从当前链接读数据的业务
	go c.StartReader()
	//TODO 启动从当前链接写数据的业务

	//

}
//停止链接，结束当前链接的工作
func (c *Connection) Stop(){
	fmt.Println("Conn stop()...ConnID= ",c.ConnID)
//如果当前链接已经关闭
if c.isClosed==true{
	return
}
c.isClosed=true
	//关闭SOCket链接
c.Conn.Close()
//关闭通道
close(c.ExitChan)
}

//获取当前链接的绑定socket conn
func (c *Connection) GetTCPConnection() *net.TCPConn{
	return c.Conn
}

//获取当前链接模块的链接ID
func (c *Connection) GetConnID() uint32{
	return c.ConnID
}
//获取远程客户端的TCP状态 IP port
func (c *Connection) GetRemoteAddr() net.Addr{
	return c.Conn.RemoteAddr()
}

//提供一个SendMsg方法，将我们要发送给客户端的数据，先进行封包
//再发送
func (c *Connection)SendMsg(msgId uint32,data []byte) error{
	if c.isClosed==true{
		return errors.New("Connection close when send msg")
	}
	//将data进行封包，msgdatalen和msgid和data
	dp:=NewDataPack()
	binaryMsg,err:=dp.Pack(NewMsgPackage(msgId,data))
	if err!=nil{
		fmt.Println("Pack error msg id=",msgId)
		return errors.New("Pack error msg")
	}
	//将数据发送给哭护短
	if _,err:=c.Conn.Write(binaryMsg);err!=nil{
		fmt.Println("Write msg id",msgId,"error",err)
		return errors.New("conn write error")
	}

	return nil
}