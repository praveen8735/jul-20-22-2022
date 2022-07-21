package main

import "fmt"

type Person struct {
	name string
	age int
	car Car
	hobbies []string
	Brain
	//AnotherBrain
	//brain Brain
}
type AnotherBrain struct {
	iq int
}
type Brain struct {
	iq int
}
type Car struct {
	model string
	engine Engine
}
type Engine struct {
	power int
}

func main() {
	var p1 = Person{name: "Sam", age: 12, car: Car { model: "BMW"}, hobbies: []string {"Eat", "Sleep"}}
	p1.age = 13
	fmt.Println(p1)
	p1.iq = 120 //p1.brain.iq

	p2 := Person{name: "Ram"}
	p2.car = Car { model: "Tesla", engine: Engine{power: 1800}}
	fmt.Println(p2.car.engine.power)
}
