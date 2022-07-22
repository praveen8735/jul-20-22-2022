package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var randomNumberChannel = make(chan int, 5)

func main() {
	rand.Seed(time.Now().UnixNano())
	go printRandomNumber()
	go generateRandomNumber()
}

func generateRandomNumber() {
	x := rand.Intn(1000)
	fmt.Println("Generated random number: " + strconv.Itoa(x))
	randomNumberChannel <- x
}

func printRandomNumber() {
	x := <-randomNumberChannel //Waits till it finds a number; Blocking call
	fmt.Println("Printing random number: " + strconv.Itoa(x))
}
