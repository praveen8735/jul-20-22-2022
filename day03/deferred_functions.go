package main

import "fmt"

func main() {
	//deferred functions are pushed to a stack FIFO
	//defer doSomething()  // 3
	//defer doSomethingElse() //2
	//defer func() {
	//	fmt.Println("inside func") //1
	//}()
	//fmt.Println("end of main")
	talkToDB()
}

func talkToDB()  {
	fmt.Println("Open a connection to DB")
	connOpen := true
	defer func() {
		fmt.Println("Close the connection to DB")
		connOpen = false
	}()
	if connOpen {
		connOpen = false
	}
	fmt.Println("Execute queries")

}

func doSomethingElse() {
	fmt.Println("doing something else")
}

func doSomething() {
	fmt.Println("doing something")
}
