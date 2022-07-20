package main

import (
	"fmt"
	"strconv"
	"strings"
)

const max = 10001

func main() {
	part1()
	part2()
	part2Democratic()
	part3()
}

func part3() {
	input := "+5 -1 +9 +5 -67 +5 -3 +2 -4 +6 +8 -13 +2 -4 +6 +18 -3 +2 -4 +6 +88 +15 -1 +9 +5 -67 +45 -3 +2 -4 +36 +8 -13 +2 -4 +6 +18 -3 +2 -74 +11 +109"
	values := strings.Split(strings.TrimSpace(input), " ")
	sum := 0
	for _, value := range values {
		num, _ := strconv.Atoi(value)
		sum += num
	}
	fmt.Println("Sum", sum)
}

func part2Democratic() {
	exclusionList := []int{10, 19, 21, 39, 309, 431, 643, 942, 1209, 7981, 8888, 9910}
	sum := 0
	for i := 1; i <= max; i++ {
		foundInExclusionList := false
		for _, value := range exclusionList {
			if i == value {
				foundInExclusionList = true
				break
			}
		}
		if !foundInExclusionList {
			sum += i
		}
	}
	average := sum / (max - len(exclusionList))
	fmt.Println("Average ", average)
}

func part2() {
	exclusionList := []int{10, 19, 21, 39, 309, 431, 643, 942, 1209, 7981, 8888, 9910}
	var sum = (max * (max + 1)) / 2 //n(n+1)/2
	for _, value := range exclusionList {
		sum -= value
	}
	average := sum / (max - len(exclusionList))
	fmt.Println("Average ", average)
}

func part1() { //Calculate the Sum of ODD numbers between 1 to 10001: 25010001
	sum := 0
	for i := 1; i <= max; i += 2 {
		sum += i
	}
	fmt.Println("Sum of odd numbers between 1 to 10001", sum)
}
