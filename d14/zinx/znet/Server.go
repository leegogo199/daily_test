package znet

import (
	"dailytest/daily_test/d13/zinx/ziface"
	"fmt"
	"net"
)

//iServer 的接口实现，定义一个server的服务器结构体
type Server struct{
	//服务器的名称
	Name string
	//服务器绑定的ip版本
	IPVersion string
	//服务器绑定的IP地址
	IP string
	//服务器绑定的端口号
	Port int
}
//启动服务器
func (s *Server)Start(){
	fmt.Printf("start server listener at ip:%s,port:%d\n",s.IP,s.Port)
	//1 获取一个tcp 的addr
	addr,err:=net.ResolveTCPAddr(s.IPVersion,fmt.Sprintf("%s:%d",s.IP,s.Port))
	if err!=nil{
		fmt.Println("resolve tcp addr error:",err)
		return
	}

	//2 监听服务器的地址
	listenner,err:=net.ListenTCP(s.IPVersion,addr)
	if err!=nil{
		fmt.Println("listen",s.IPVersion,"err ",err)
	}
	fmt.Println("start zinx server succ",s.Name,"listenning...")



	//3 阻塞等待客户端连接，处理业务
	for{
		//如果有客户端将连接过来，阻塞会返回
		conn,err:=listenner.AcceptTCP()
		if err!=nil{
			fmt.Println("Accept err",err)
			continue
		}
		//客户端建立连接，做一些业务,做一个最基本的512字节长度的回写业务
		go func(){
				for {
					buf:=make([]byte,512)
					cnt,err:=conn.Read(buf)
					if err!=nil{
						fmt.Println("recv buf err",err)
						continue
								}
					//回写功能
					if _,err:=conn.Write(buf[:cnt]);err!=nil{
						fmt.Println("write back buf err ",err)
						continue
															}
					}
				}()
		}

}
//停止服务器
func (s *Server)Stop(){
//todo 释放占用的资源

}
//运行服务器
func (s *Server)Serve(){
	//启动server的服务功能
	s.Start()
	// TODO 做一些额外工作
	//
	//阻塞状态
	select{}
}

//初始化server
func NewServer(name string)ziface.IServer{
	s:=&Server{
		Name:name,
		IPVersion: "tcp4",
		IP:"0.0.0.0",
		Port:8999,
	}
	return s
}

