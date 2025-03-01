package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Syafiq"
	middleName = "Fajrian"
	lastName = "Emha"

	return firstName, middleName, lastName
}

func main() {
	a, b, c := getCompleteName()
	fmt.Println(a, b, c)
}
