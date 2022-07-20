package main

import "fmt"

func main() {
	//var size2 int = 10
	// numbers := [10]int{}
	// fmt.Println(numbers)

	// println("Enter the size of your array")
	// var size int
	// fmt.Scanf("%d", &size)
	// myNumbers := make([]int, size)
	// fmt.Println(myNumbers)
	copyingArr()
}
func copyingArr() {
	var numbersColl = []int{100, 200, 300}
	var anotherNumbersColl = numbersColl //Shallow copy
	fmt.Println(numbersColl, anotherNumbersColl)
	anotherNumbersColl[0] = 1000 //Modifies numbersColl also
	fmt.Println(numbersColl, anotherNumbersColl)

	var yetAnotherNumbersColl = make([]int, len(numbersColl))
	copy(yetAnotherNumbersColl, numbersColl) //deep copy
	yetAnotherNumbersColl[2] = 999999
	fmt.Println(numbersColl, yetAnotherNumbersColl)
}
