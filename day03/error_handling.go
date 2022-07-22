package main

import (
	"fmt"
	"log"
)

func main()  {
	defer func() {
		if rec := recover(); rec != nil {
			log.Fatal("Error: ", rec)
		}
	} ()
	num := 10
	den := 0
	ans := num / den	//raises a panic
	fmt.Println(ans)
	fmt.Println("End of main")
}