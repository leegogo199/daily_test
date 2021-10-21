package znet


//iServer 的接口实现，定义一个server的服务器结构体
type Server struct{
	//服务器的名称
	Name string
	//服务器绑定的ip版本
	IPVsersion string
	//服务器绑定的IP地址
	IP string
	//服务器绑定的端口号
	Port int
}
//启动服务器
func (s *Server)Start(){

}
//停止服务器
func (s *Server)Stop(){

}
//运行服务器
func (s *Server)Serve(){

}
