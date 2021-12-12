package main

import "fmt"
//编译器静态属性清楚
//执行阶段 反射、接口。类型断言都很麻烦。
//built-in type int  no method
//类型描述信息，全局唯一，构成类型系统。
//类型名称，类型大小，对齐边界，是否自定义
//存放于 runtime——type中，其他描述形式，，包路径，方法数目，方法元数据，数组的偏移值。
// type myt=int32 别名 关联到同一个类型元数据
//type myt int32 自定义类型创建新类型，拥有自己的类型元数据。
//
//
type T struct{
	name string
}
func (t T)F1(){
	fmt.Println(t.name)
}
func main(){
	t:=T{name:"eggo"}
	t.F1()
}
//interface{}
//
//
//
//
//
/*
空接口，可接受任何数据
_type *_type动态类型
data unsafe.Pointer动态值
 e
_type =nil
data =nil

f,_:=os.open()
e=f
e
_type ->*os.File 方法元数据数组
data=f
非空接口
type iface struct{
	tab *itab
	data unsafe.Pointer
}
type itab struct{
inter *interfacetype 类型元数据
_type *_type 动态类型元数据
hash uint32 该类型的hash值
_ [4]byte
fun [1]uintptr 要求实现的方法
}
类型元数据
type interfacetype struct{
type _type
pkgpath name
mhdr []imethod方法列表




}
var rw io.ReadWriter
//此时的rw
tab=nil
data=nil

f,:=os.Open("egg")
rw=f
此时的rw          接口和动态确定了，itab也就确定了
tab=itab {inter {ioead readwtietr的 元数据} _type 类型元数据，fun【0】方法元数据拷贝方法地址}
data=f
缓存itab
接口类型和动态类型为key
*itab为value
runtime.itabtabletype





 */
