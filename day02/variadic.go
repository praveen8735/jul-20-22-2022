package main

import "fmt"

func main() {
	fmt.Println(addOldStyle([]int{100, 200, 300}))
	fmt.Println(addOldStyle([]int{1, 2, 3, 4}))
	add(100, 200, 300)
	add(1, 2, 3, 4, 5)
	arr := []int{90, 91, 92, 93}
	add(arr...) //spread operator (add(90, 91, 92, 93))
}

func add(numbers ...int) (sum int) {
	for _, num := range numbers {
		sum += num
	}
	return
}

func addOldStyle(numbers []int) (sum int) {
	for _, num := range numbers {
		sum += num
	}
	return sum
}
