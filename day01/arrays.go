package main

import "fmt"

func main() {
	//[size]type
	var numbers [5]int
	fmt.Println(numbers)

	var names = [3]string{"Sam", "Ram", "Mary"}
	fmt.Println(names)

	var cities = []string{"Cochin"}
	fmt.Println(cities)

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for key, value := range names {
		fmt.Println(key, value)
	}

	for _, value := range names {
		fmt.Println(value)
	}

	for index, _ := range names {
		fmt.Println(index)
	}

}
