package main

import (
	"fmt"
	"math/rand"
)

type currency float32
type stock string

func main() {
	var dollar currency = 80.12
	var euro currency = 101.23
	//printValue(dollar)
	//printValue(euro)
	dollar.printValue()
	euro.printValue()
	var ibm stock = "IBM"
	var facebook stock = "FB"
	//getCurrentMarketPrice(ibm, "Nasdaq")
	//getCurrentMarketPrice(facebook, "S&P")
	ibm.getCurrentMarketPrice("NASDAQ")
	facebook.getCurrentMarketPrice("S&P 500")

}
func (st stock) getCurrentMarketPrice(market string) {
	fmt.Println("CMP of ", st, "in", market, ": ", rand.Intn(10000))
}

func (c currency) printValue() {
	fmt.Println(c)
}
