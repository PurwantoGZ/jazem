package example

import "fmt"

//Operator is function to learn about operator in golang
func Operator(test string) {
	var value = ((2+6)%3*4 - 2) / 3
	var isEquel = (value == 2)

	fmt.Printf("nilai %d (%t) \n", value, isEquel)
}
