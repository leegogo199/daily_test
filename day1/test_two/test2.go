package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)
func main(){
	c,err:=redis.Dial("tcp","localhost:6379")
	if err!=nil{
		fmt.Println("conn redis failed.",err)
	}
	defer c.Close()
	_,err=c.Do("expire","first",10)
	if err!=nil{
		fmt.Println(err)
		return
	}

}
