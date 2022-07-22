package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	wg.Add(3)
	go watchHotstar()
	go eatPopcorn()
	go func() {
		fmt.Println("Do something anonymously")
		wg.Done()
	}()
	wg.Wait() //Blocks main

	fmt.Println("End of main")
}

func eatPopcorn() {
	fmt.Println("Eating")
	time.Sleep(time.Second * 10)
	wg.Done()
}
func watchHotstar() {
	fmt.Println("Watching movie")
	//time.Sleep(time.Second * 2)
	wg.Done()
}
