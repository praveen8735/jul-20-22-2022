* Create __lab_07.go__ file
* In this exercise you'll use __cities.csv__ file to perform some operations
* Study the __cities.csv__ file in the repository
* Define a __City struct__ accordingly
<br/>

* Define functions to load the cities.csv, convert every item to the City struct and maintain in a collection
* After loading the cities into a collection implement the following


### To Do
* Print all the cities grouped by countries
* Print the city name with largest and smallest population
* Print China's population

* Code to load the csv file is given below. You can use it and modify it appropriately 

``` go
import (
	"bufio"
	"os"
)
	fileName := "cities.csv"
	file, err := os.Open(fileName)
	scanner := bufio.NewScanner(file)
	if err != nil {
		fmt.Println(err)
	} else {
		defer file.Close()
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println(line)
		}
``` 

* Note: You'll learn _defer_ keyword later





