//抽象工厂模式
package abstractfactory

import "fmt"

type Lunch interface{
	Cook()
}
type Rise struct{

}
func (r *Rise) Cook(){
	fmt.Println("it is rise")
}
type Tomato struct{

}
func (t *Tomato) Cook(){
	fmt.Println("it is tomato")
}
type LunchFactory interface{
	CreateFood() Lunch
	CreateVeg() Lunch
}
type SimpleLunchFactory struct{
}
func (s *SimpleLunchFactory)CreateFood()Lunch{
	return &Rise{}
}
func (s *SimpleLunchFactory)CreateVeg()Lunch{
	return &Tomato{}
}
func NewSimpleLunchFactory() LunchFactory{
	return &SimpleLunchFactory{}
}
