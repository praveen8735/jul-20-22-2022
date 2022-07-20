package main

import "fmt"

func main() {
	count := 0
	for count < 5 { //also used as a while loop
		fmt.Println(count)
		count++
	}
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

}
