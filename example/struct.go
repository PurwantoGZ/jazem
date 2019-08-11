package example

import "fmt"

type student struct {
	name  string
	grade int
}

// Getstruct is function for learn about struct in golang
func Getstruct(test string) {
	fmt.Printf(" belajar %s\n", test)
	var s1 student
	s1.name = "Purwanto"
	s1.grade = 4

	fmt.Println(s1.name)
	fmt.Println(s1.grade)
}
