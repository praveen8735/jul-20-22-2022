package main

import (
	"fmt"
	"time"
)

//By default every go application has one go routine running
//That is the main function
func main() {
	go eat()
	go watchNetflix()
	fmt.Println("End of main")
	time.Sleep(time.Second * 2)
}

func eat() {
	time.Sleep(time.Second * 1)
	fmt.Println("Eating")
}

func watchNetflix() {
	fmt.Println("Watching a movie on Netflix")
}