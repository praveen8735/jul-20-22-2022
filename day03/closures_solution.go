package main

import "fmt"

func main() {
	x := 1 //var x = 1
	x++

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
