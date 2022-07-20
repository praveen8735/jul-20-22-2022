package main

import "fmt"

func main() {
	numbers := []int{100, 200, 300, 400, 500}
	firstThreeNumbers := numbers[0:3] //slice but shallow copy
	fmt.Println(firstThreeNumbers)
	lastThreeNumbers := numbers[len(numbers)-3 : len(numbers)]
	fmt.Println(lastThreeNumbers)
	firstThreeNumbers[0] = 982347
	fmt.Println(numbers, firstThreeNumbers)

}
