package main

import "fmt"

func main() {
	const firstName string = "Syafiq Fajrian"
	const lastName = "Emha"

	// error
	// firstName = "Budi"

	const (
		age    int = 21
		gender     = "Male"
	)
	fmt.Println(age)
	fmt.Println(gender)
}
