package main

import "fmt"

var numbers = []int { 101, 22, 43, 14, 5, 906, 310, 561, 84, 23, 100 }
func main() {
	//Part I
	var isEven = func(num int) bool {
		return num % 2 == 0
	}
	evenNumbers := find(isEven)
	oddNumbers := find(func(num int) bool {
		return num % 2 != 0
	})
	divBy5Numbers := find(func(num int) bool {
		return num % 5 == 0
	})
	fmt.Println(evenNumbers, oddNumbers, divBy5Numbers)

	//Part II
	var square = func(num int) int {
		return num * num
	}
	var cube = func(num int) int {
		return num * num * num
	}
	var sumOfSquares = compute(square)
	fmt.Println("Sum of squares", sumOfSquares)

	var sumOfCubes = compute(cube)
	fmt.Println("Sum of cubes", sumOfCubes)
}

func compute(arg func(int) int) (sum int)  {
	for _, value := range numbers {
		result := arg(value)
		sum += result
	}
	return
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