package main

import "fmt"

type currency float32

func main() {

	var dollar currency = 80.23
	var euro currency = 101.12
	fmt.Println(dollar, euro)

	var name string = "Ram"
	fmt.Println(name)
}
