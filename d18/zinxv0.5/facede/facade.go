package Facede

import "fmt"

type CarModel struct{

}
func NewCarModel()*CarModel{
	return &CarModel{}
}

func (c* CarModel)SetCarModel(){
	fmt.Println("Car model--Select Model")
}
type CarEngine struct{

}

func NewCarEngine()*CarEngine{
	return &CarEngine{}
}
func (c *CarEngine)SetCarEngine(){
	fmt.Println("Car Engine set body")
}
type CarBody struct{

}

func NewCarBody() *CarBody{
	return &CarBody{}
}
func (c *CarBody)SetCarBody(){
	fmt.Println("Car body set body")
}
type CarFacade struct{
	model CarModel
	engine CarEngine
	body CarBody
}
func NewCarFacade() *CarFacade{
	return &CarFacade{
		model:CarModel{},
		engine: CarEngine{},
		body:CarBody{},
	}
}
func (c *CarFacade)CreateCompleteCar(){
	c.model.SetCarModel()
	c.engine.SetCarEngine()
	c.body.SetCarBody()
}