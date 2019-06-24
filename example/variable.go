package example

import "fmt"

//Variable is function to show variable
func Variable(test string) {

	fmt.Printf("Ini test %s \n", test)

	var firstName string = "Purwanto"

	var lastName string
	lastName = Guguk(test)

	fmt.Printf("Hallo %s %s ! \n", firstName, lastName)
}
