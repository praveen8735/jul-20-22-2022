package main

import (
	"fmt"
	"math/rand"
	"time"
)

var globalVar int

func main() {
	x := 10
	// if x > 10
	// 	fmt.Println("gt 10")

	rand.Seed(time.Now().UnixNano())
	if x > 10 {

	} else if x > 10 {

	} else {

	}

	if x%2 == 0 { //Paranthesis is not required in if-else
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
	if y := 10; y%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
	day := "Tuesday"
	switch day {
	case "Tuesday":
		fmt.Println("Weekday")
	case "Friday":
		fmt.Println("TGIF")
	case "Saturday", "Sunday":
		fmt.Println("Weekend")
	default:
		fmt.Println("Boreeee")
	}

}
