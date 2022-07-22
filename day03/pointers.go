package main

import "fmt"
type Sample struct {
	num int
}
func main() {
	x := 10
	var ptr *int = &x
	fmt.Println(*ptr)
	*ptr = 111
	fmt.Println(x, *ptr)
	//ptr++ pointer arithmetic is not allowed
	var ptrToAString = new(string)
	fmt.Println(*ptrToAString)
	sample := Sample {num: 12}
	modifySample(&sample)
	fmt.Println(sample)

}
func modifySample(ptr *Sample) {
	ptr.num = 10000
}