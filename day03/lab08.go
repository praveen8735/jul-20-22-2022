package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func main()  {
	wg.Add(len(symbols))
	for _, symbol := range symbols {
		go getCurrentMarketPriceFromServer(symbol)
	}
	wg.Wait()
	generateJson()
}

func generateJson() {
	contents, _ := json.Marshal(stockList)
	fmt.Println(string(contents))
}

func getCurrentMarketPriceFromServer(symbol string) {
	url := baseUrl + symbol
	response, _ := http.Get(url)
	output, _ := ioutil.ReadAll(response.Body)
	defer func() {
		response.Body.Close()
		wg.Done()
	}()
	stock := Stock{}
	json.Unmarshal(output, &stock)
	stockList = append(stockList, stock)
}
