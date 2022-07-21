package main

import "fmt"

type EmptyFunction func()

func main() {
	var work = func() {
		fmt.Println("Working")
	}
	var sleep = func() {
		fmt.Println("Sleeping")
	}
	doSomething(work)
	doSomething(sleep)

}
func doSomething(arg func()) {
	arg()
}

func doSomethingElse(arg EmptyFunction) {
	arg()
}
