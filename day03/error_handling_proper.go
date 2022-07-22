package main

import (
	"fmt"
	"log"
)

func main()  {
	divide()
	fmt.Println("End of main")
}

func divide() {
	defer catch()
	num := 10
	den := 0
	if den == 0 {
		panic("You cannot divide by zero")
	}
	ans := num / den
	fmt.Println(ans)
}

func catch() {
	if rec := recover(); rec != nil {
		log.Println(rec)
	}
}