package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var name string = "hello"
	var letter rune = 'a'             //rune is similar to character, though it's int32
	fmt.Println(name, string(letter)) //string(variable)
	fmt.Println(string(name[0]), string(name[1]))
	fmt.Println(len(name))
	fmt.Println(strings.ToUpper(name))
	fmt.Println(strings.ReplaceAll(name, "l", "z"))
	i := "12"
	fmt.Println(strconv.Atoi(i))

	paragraph := `This
	is a
	paragraph`
	fmt.Println(paragraph)

	city := "Chennai"
	for i := 0; i < len(city); i++ {
		fmt.Println(string(city[i]))
	}
}
