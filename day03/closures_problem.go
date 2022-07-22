package main

import "fmt"

func main() {
	buttonClicked()
	buttonClicked()
	count = 51
	buttonClicked()
	fmt.Println(count)
}

var count = 0
func buttonClicked() {
	count++
}