package main

import "fmt"

type Book struct {
	title string
	price int
	read func() //NOT RECOMMENDED
}

func main() {
	b1 := Book{title: "ABC", price: 100, read: func() {
		fmt.Println("confused")
	}}
	increasePrice(b1) //pass by value and not pass by reference. Original structure doesn't get modified
	fmt.Println(b1)
	//buy(b1, "Amazon")
	b1.buy("Amazon")
	b1.read()
}

func (book Book) buy(store string)  {
	fmt.Printf("Buying %s from %s\n", book.title, store)
}

func increasePrice(b Book) {
	b.price = b.price + 100
}
