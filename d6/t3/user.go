package main
import(
	"golang.org/x/net/dns/dnsmessage"
	"net"
	"strings"
)
type User struct{
	Name string
	Addr string
	C chan string
	conn net.Conn
	server *Server
}
//创建一个用户的API
func NewUser(conn net.Conn,server *Server) *User{
	userAddr:=conn.RemoteAddr().String()
	user:=&User{
		Name:userAddr,
		Addr:userAddr,
		C:make(chan string),
		conn:conn,
		server: server,
	}
	//启动监听当前user channel 消息的goroutine
	go user.ListenMessage()
	return user
}
// 监听当前user channel的方法，一旦有消息，就直接发送给对端客户端。
func (u *User)ListenMessage(){
	for{
		msg:=<-u.C
		u.conn.Write([]byte(msg+"\n"))
	}
}
//
func (u *User)Online(){
	// 广播当前用户上线消息
	u.server.maplock.Lock()
	u.server.OnlineMap[u.Name]=u
	u.server.maplock.Unlock()
	u.server.BroadCast(u,"已上线")

}
//下线
func (u *User)Offline(){
	u.server.maplock.Lock()
	delete(u.server.OnlineMap,u.Name)
	u.server.maplock.Unlock()
	u.server.BroadCast(u,"已下线")
}

//给当前User对应的客户端发送消息
func (u *User)SendMsg(msg string){
	u.conn.Write([]byte(msg))
}
//处理业务
func (u *User)DoMessage(msg string){
   if msg=="who"{
	// 查询当前在线用户；
		u.server.maplock.lock()
			for _,user:=range u.server.OnlineMap{
				onlineMsg:="["+user.Addr+"]"+user.Name+":"+"在线。。。\n"
				u.SendMsg(onlineMsg)
												}
		u.server.maplock.Unlock()				
	}else if len(msg)>7&& msg[:7]=="rename|"{
		//消息格式 rename张三
		newname:=strings.Split(msg,"|")[1]
		//判断newname是否存在
		_,ok:=u.server.OnlineMap[newname]
		if ok{
			u.SendMsg("当前用户名被使用、\n")

		}else{
			u.server.maplock.Lock()
			delete(u.server.OnlineMap,u.Name)
			u.server.OnlineMap[newname]=u
			u.server.maplock.Unlock()
			u.Name=newname
			u.SendMsg("您已经更新用户名:"+u.Name+"\n")

		}
   } else{
		u.server.BroadCast(u,msg)
}
}