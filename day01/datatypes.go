package main

import "fmt"

var company string

//const PI = 3.14
//const TRUE = true
const (
	PI         = 3.14
	TRUE       = true
	FALSE bool = false
)

//state := "KA"

func main() {
	/*
		multi line comment
	*/
	const key = 100
	fmt.Println("Datatypes")
	//int, bool, string, float, character
	//var variableName type

	var num int = 10
	var name string = "Sam"
	var isAvailable bool = true

	var stockPrice float32
	fmt.Println(num, name, isAvailable, stockPrice)

	//Type inference
	var x = 12
	//x = "string"
	fmt.Printf("%T\n", x)

	//var i int = 100
	//Shortcut declarations
	i := 100
	city := "Bangalore"
	inStock := true

	fmt.Println(i, city, inStock)

}
