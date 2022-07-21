package main

import "fmt"

func main() {
	var add = func(x int, y int) int {
		return x + y
	}
	var sub = func(x int, y int) int {
		return x - y
	}
	var multiply = func(x int, y int) int {
		return x * y
	}
	//var square = func(x int) int { return x * x }
	//doSomeMath(square)
	doSomeMath(add)
	doSomeMath(sub)
	doSomeMath(multiply)
	doSomeMath(func(x int, y int) int { return x / y })
}

type mathFunction func(int, int) int

func doSomeMath(operation mathFunction) {
	//func doSomeMath(operation func(int, int) int) {
	var num1 int = 10
	var num2 int = 20
	fmt.Println(operation(num1, num2))
}

/*
//OLD STYLE
func main() {
	fmt.Println(add(num1, num2))
	fmt.Println(sub(num1, num2))
	fmt.Println(multiply(num1, num2))
}

func add(x int, y int) int {
	return x + y
}

func sub(x int, y int) int {
	return x - y
}

func multiply(x int, y int) int {
	return x * y
}
*/
