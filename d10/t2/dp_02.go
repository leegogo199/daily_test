package Abstract_Factory

import "fmt"

type Lunch interface{
	Cook()
}
type Rise struct{

}
func (r *Rise) Cook{
	fmt.Println("it is  rise.")
}