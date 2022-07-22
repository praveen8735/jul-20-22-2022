package main

import (
	"fmt"
)

func main() {
	//channel is way to pass information between go routines
	//called inter thread communication in other languages
	introToChannel()
	fmt.Println("end of main")
}
func introToChannel() {
	chn := make(chan int, 1) //create a channel that can have 1 number
	chn <- 10
	x :=  <- chn //blocking call; waits for the channel to have data
	fmt.Println(x)
}
