//工厂方法。
package factory
//点外卖生产外卖消费。
import "fmt"
type Restaurant interface{
	GetFood()
}
type KFC struct{

}
func (k *KFC) GetFood(){
	fmt.Println("KFC's food is okay,continue ...")

}
type MCD struct {

}
func (m *MCD)GetFood(){
	fmt.Println("MCD's food is okay,continue...")
}
func NewRestaurant(name string) Restaurant{
	switch name{
		case "k":
			return &KFC{}
		case "m":
			return &MCD{}

	}
	return nil
}


