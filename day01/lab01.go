package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(99) + 1
	guess := -1
	gameOver := false
	attempts := 0
	fmt.Println("Enter a number between 1 and 100")

	for !gameOver {
		fmt.Scanf("%d", &guess)
		attempts++
		if guess > target {
			fmt.Println("Aim lower")
		} else if guess < target {
			fmt.Println("Aim higher")
		} else {
			fmt.Printf("You've got it in %d attempts\n", attempts)
			gameOver = true
		}
	}

}
