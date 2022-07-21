package main

import "fmt"

var numbers = []int { 101, 22, 43, 14, 5, 906, 310, 561, 84, 23, 100 }
func main() {
	var isEven = func(num int) bool {
		return num % 2 == 0
	}
	evenNumbers := find(isEven)
	fmt.Println(evenNumbers)
}

func find(arg func(int) bool) []int {
	var resultArr = make([]int, 0)
	for _, value := range numbers {
		result := arg(value)
		if result {
			resultArr = append(resultArr, value)
		}
	}
	return resultArr
}