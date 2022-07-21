package main

import (
	"fmt"
)

func main() {
	var eat = func() {
		fmt.Println("Eating")
	}
	eat()
	fmt.Printf("%T\n", eat)

	var add = func(a int, b int) int {
		return a + b
	}
	fmt.Println(add(12, 13))

}
