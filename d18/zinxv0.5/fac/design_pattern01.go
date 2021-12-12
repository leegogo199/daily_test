//工厂方法模式
package fac

import(
	"fmt"
	)
type Fandian interface{
	Getfood()
}
type Donlaishun struct{

}
func (d *Donlaishun) Getfood(){
	fmt.Println("donglaishun shengchanwanbi jiuxu")
}
type Dezhuang struct{

}
func (d *Dezhuang)Getfood(){
	fmt.Println("dezhuangshengchanwanbi")
}
func NewFandian(name string) Fandian {
	switch name{
	case "d":
		return &Donlaishun{}
	case "de":
		return  &Dezhuang{}
	}
	return nil
}

func main(){

}

