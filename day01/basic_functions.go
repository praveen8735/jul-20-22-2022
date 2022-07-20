package main

import "fmt"

func main() {
	fmt.Println(add(10, 20))
	//a, b := 1, 2
	a, b := tuples()
	fmt.Println(a, b)

	x, _ := tuples()
	fmt.Println(x)

	sq, cu := compute(12)
	fmt.Println(sq, cu)
}
func compute(num int) (square int, cube int) {
	square = num * num
	cube = num * num * num
	return
}

func tuples() (int, int) {
	return 1, 2
}

func doSomething() {
	fmt.Println("Doing something")
}

func add(a int, b int) int {
	return a + b
}

func subtract(a int, b int) (result int) {
	result = a - b
	return
}
