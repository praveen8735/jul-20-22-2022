package main

import "fmt"

type Car struct {
	model string
}
type Truck struct {
	power int
}
type Vehicle interface {
	printInfo()
}
func (c Car) printInfo() {	//class Car implements Vehicle { printInfo(...) }
	fmt.Println("Car", c.model)
}
func (t Truck) printInfo() { //class Truck implements Vehicle { printInfo(...) }
	fmt.Println("Truck", t.power)
}

func main() {
	bmw := Car {model: "BMW"}
	teslaTruck := Truck{power: 5000}
	toll(bmw)
	toll(teslaTruck)
	vehicles := []Vehicle { bmw, teslaTruck }
	fmt.Println(vehicles)
	//passThroughToll(bmw)
	//passThroughTollgate(teslaTruck)
}

func toll(v Vehicle) {
	v.printInfo()
}

func passThroughToll(car Car) {
	fmt.Println("Car", car.model)
}
func passThroughTollgate(truck Truck) {
	fmt.Println("Truck", truck.power)
}

