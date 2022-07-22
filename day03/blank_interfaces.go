package main

import "fmt"

type Empty interface {}
type Book struct {}
func main() {
	doSomething("hello")
	doSomething(12)
	doSomething(true)
	b1 := Book{}
	doSomething(b1)
	genericCollection := []Empty { 1, true, "hello", Book{}, []int{}}
	fmt.Println(genericCollection)
	printDetails(1, true, "hello", Book{}, []int{})
	printDetails("hello", "bye")
}

func printDetails(any ... Empty) {

}

func doSomething(any Empty) {
	fmt.Println(any)
}