package main

import "fmt"

func main() {
	buttonClicked := getButtonClickedFunc()
	numberOfTimes := buttonClicked()
	numberOfTimes = buttonClicked()
	numberOfTimes = buttonClicked()
	fmt.Println(numberOfTimes)

}

func getButtonClickedFunc() func() int{
	var count = 0
	return func() int{
	  count++
	  return count
   }
}
