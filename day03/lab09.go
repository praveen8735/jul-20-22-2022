package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

type Stock struct {
	Symbol string
	Cmp float32
}

var symbols = []string { "HP", "GOOG", "FB", "WIPRO", "INFY", "DELL", "HUL", "ITC", "RELIND", "NETFLIX", "AMAZON" }
var baseUrl = "http://healthycoder.in/stocks.php?symbol="
var stockList = []Stock{}
var wg = sync.WaitGroup{}
var stockOutputChannel = make(chan []byte, 1)

func main()  {
	wg.Add(len(symbols))
	for _, symbol := range symbols {
		go getCurrentMarketPriceFromServer(symbol)
	}
	go processJsonOutput()
	wg.Wait()
	printStockJson()
}

func getCurrentMarketPriceFromServer(symbol string) {
	url := baseUrl + symbol
	log.Println("Fetching data for " + symbol + " ...")
	response, _ := http.Get(url)
	output, _ := ioutil.ReadAll(response.Body)
	defer func() {
		response.Body.Close()
	}()
	stockOutputChannel <- output

}

func processJsonOutput() {
	for len(stockList) != len(symbols) {
		jsonOutput := <- stockOutputChannel
		stock := Stock{}
		json.Unmarshal(jsonOutput, &stock)
		log.Println("Processing " + string(jsonOutput))
		stockList = append(stockList, stock)
		wg.Done()
	}
}

func printStockJson() {
	contents, _ := json.Marshal(stockList)
	log.Println(string(contents))
}