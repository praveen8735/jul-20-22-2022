* You have a collection Stock Symbols
* HP, GOOG, FB, WIPRO, INFY, DELL, HUL, ITC, RELIND, NETFLIX, AMAZON
* Maintain a collection of these symbols
* You'll connect to a service and pass the symbol
* The web service gives you the current stock price that can be displayed in the console.
* I have created a service and hosted it in my blog __http://healthycoder.in__

* Go to __http://healthycoder.in/stocks.php?symbol=GOOG__
* You'll get the following JSON response

```json
{"symbol": "GOOG", "cmp": 73.19}
```
* cmp is __current market price__
* The service has delay added. So each time you access the service by passing a symbol, you'll get the response with a delay of 2 seconds or 3 seconds or 4 or more.

* __You'll design the application in such a way that, every symbol executes in a separate go routine. The main waits till all the stock prices are fetched.__

* Deserialize JSON response and display it on the console

#### Note

* Use the following code to talk to a server

``` go
response, _ := http.Get("url in string")
output, _ := ioutil.ReadAll(resp.Body)
defer func() {
		response.Body.Close()
}()
```