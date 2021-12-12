package abstractfactory

import "testing"

func TestNewSimpleLunchFactory(t *testing.T) {
	factory:=NewSimpleLunchFactory()
	food:=factory.CreateFood()
	food.Cook()
	veg:=factory.CreateVeg()
	veg.Cook()
}
