package example

import "fmt"

//Condition is function to learn about condition in golang
func Condition(test string) {
	fmt.Printf(" Learn about %s \n", test)

	var point = 8840.0
	if percent := point / 100; percent >= 0 {
		fmt.Printf("%.1f%s perfect !\n", percent, "%")
	}
}
