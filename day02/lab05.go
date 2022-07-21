package main

import "fmt"

type MyNumbers []int

func main() {
	myCollection := MyNumbers{10, 200, 34, 78, 9}
	myCollection.sum()
	myCollection.max()
	myCollection.min()
}

func (lst MyNumbers) min() {
	smallest := lst[0]
	for _, value := range lst {
		if value < smallest {
			smallest = value
		}
	}
	fmt.Println("Min", smallest)
}

func (lst MyNumbers) max() {
	largest := lst[0]
	for _, value := range lst {
		if value > largest {
			largest = value
		}
	}
	fmt.Println("Max", largest)
}

func (lst MyNumbers) sum() {
	total := 0
	for _, value := range lst {
		total += value
	}
	fmt.Println("Sum", total)
}
