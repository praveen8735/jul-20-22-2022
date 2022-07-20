package main

import "fmt"

func main() {
	//map[keyType]valueType

	var romanNumbers map[int]string = map[int]string{1: "I", 2: "II"}
	romanNumbers[3] = "III"
	delete(romanNumbers, 2)
	fmt.Println(romanNumbers)
	for key, value := range romanNumbers {
		fmt.Println(key, value)
	}

	var capitals map[string]string = make(map[string]string, 2)
	fmt.Println(capitals)
	capitals["India"] = "ND"
	capitals["US"] = "DC"
	capitals["England"] = "London"
	capitals["France"] = "Paris"

	fmt.Println(capitals)

}
