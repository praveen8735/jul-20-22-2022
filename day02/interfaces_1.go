package main

import "fmt"

type Circle struct {
	radius int
}
type Square struct {
	side int
}
type Rectangle struct {
	len int
	breadth int
}

type Shape interface {
	area() int
}

func (s Square) area() int{
	return s.side * s.side
}

func (c Circle) area() int{
	return c.radius * c.radius * 3 //BAD
}
func (r Rectangle) area() int {
	return r.len * r.breadth
}

func main() {
	c1 := Circle{radius: 12}
	sq1 := Square{side: 12}
	rect := Rectangle{len: 12, breadth: 13}

	var shapesList []Shape = []Shape{c1, sq1, rect}
	fmt.Println(shapesList)

	//var shapesList []Circle = []Circle { c1, sq1}
	//var shapesList2 []Square = []Square { c1, sq1}
}
