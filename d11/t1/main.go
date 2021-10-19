package main

import(
"net"
"fmt"
//"sync"
)

func main(){

    // wg:=sync.WaitGroup()
    var port string
    for i:=0;i<100;i++{
       port=fmt.Sprintf("127.0.0.1:%d",i)
       conn,err:=net.Dial("tcp",port)
       if err!=nil{
        fmt.Printf("%s 关闭了" ,port)
        continue
       }
       conn.Close()
    }

}

