package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
type City struct {
	name string
	country string
	population int
}
var allCities []City = make([]City, 0)
var countryCitiesMap map[string][]City = make(map[string][]City)

func main() {
	readFile("./cities.csv")
	printTheCityNameWithLargestAndSmallestPopulation()
	printPopulationOf("China")
}

func printPopulationOf(cityName string) {
	cities := countryCitiesMap[cityName]
	total := 0
	for _, city := range cities {
		total += city.population
	}
	fmt.Println("Population of ", cityName, " ", total)
}

func printTheCityNameWithLargestAndSmallestPopulation() {
	smallestPopulation := allCities[0].population
	smallestCity := allCities[0].name
	largestPopulation := allCities[0].population
	largestCity := allCities[0].name

	for _, city := range allCities {
		population := city.population
		cityName := city.name
		if population > largestPopulation {
			largestPopulation = population
			largestCity = cityName
		}
		if population < smallestPopulation {
			smallestPopulation = population
			smallestCity = cityName
		}
	}
	fmt.Println("Smallest city", smallestCity, smallestPopulation)
	fmt.Println("Largest city", largestCity, largestPopulation)
}



func constructCity(line string) {
	values := strings.Split(line, ",")
	cityName := values[0]
	countryName := values[1]
	population, _ := strconv.Atoi(values[2])
	newCity := City{
		name: cityName,
		country: countryName,
		population: population,
	}
	allCities = append(allCities, newCity)
	existingCitiesArr := countryCitiesMap[countryName]  //Fetch the cities in the country
	existingCitiesArr = append(existingCitiesArr, newCity) //append the new city to the array
	countryCitiesMap[countryName] = existingCitiesArr //reassign cities to the map
}

func readFile(fileName string) {
	file, err := os.Open(fileName)
	scanner := bufio.NewScanner(file)
	if err != nil {
		fmt.Println(err)
	} else {
		defer file.Close()
		for scanner.Scan() {
			line := scanner.Text()
			if(!strings.Contains(line, "City")) {
				constructCity(line)
			}

		}
	}
}