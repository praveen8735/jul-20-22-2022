package main

import (
	"fmt"
	"strconv"
)

/*
interface PersonInterface {
	eat();
	sleep(int hours);
}
Class Person implements PersonInterface {
	String name;
	int age;
	eat() {}
	sleep(int hours) {}
}
*/

//struct is used to list the data attributes
//interface is used to list the behaviours

type Person struct {
	name string
	age int
}
type PersonInterface interface {
	eat()
	sleep(hours int)
}
func (p Person) eat() {
	fmt.Println(p.name + " is eating")
}
func (p Person) sleep(hours int) {
	fmt.Println(p.name + " is sleeping for " + strconv.Itoa(hours) + " hours")
}
