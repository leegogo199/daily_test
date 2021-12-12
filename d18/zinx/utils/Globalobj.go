package utils

import (
	"dailytest/daily_test/d18/zinx/ziface"
	"encoding/json"
	"io/ioutil"
)

//存储一切有关zinx的全局参数，
//供其他模块调用。
type GlobalObj struct {
	// server
	//当前zinx全局的server对象
	TcpServer ziface.IServer
	//当前服务器主机监听的IP
	Host string
	//当前服务器主机监听的端口号
	TcpPort int
	//当前服务器的名称
	Name string
	//
	// ZInx
	//当前zinx的版本号
	Version string
	//当前服务器主机允许的最大连接数
	MaxConn int
	//当前zinx框架数据包的最大值
	MaxPackageSize uint32
}
// 定义一个全局的对外GlobalObj
var GlobalObject *GlobalObj
//从zinx.json去加载用户自定义的参数。
func (g *GlobalObj) Reload() {
	data,err:=ioutil.ReadFile("conf/zinx.json")
	if err!=nil{
		panic(err)
	}
	//将json文件数据解析到struct中
	err=json.Unmarshal(data,&GlobalObject)
	if err!=nil{
		panic(err)
	}


}
//提供一个init方法，初始化当前的GLobal Object
func init(){
	//如果配置文件没有加载，默认的值
	GlobalObject=&GlobalObj{
		Name:"ZinxServerApp",
		Version: "V0.4",
		TcpPort:8999,
		Host:"0.0.0.0",
		MaxConn:1000,
		MaxPackageSize:4096,
	}
	//应该尝试从conf/zinx.json 去加载一些用户自定义的参数
	//GlobalObject.Reload()

}