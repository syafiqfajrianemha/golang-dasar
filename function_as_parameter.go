package main

import "fmt"

// func sayHelloWithFilter(name string, filter func(string) string) {

type Filter func(string) string // buat alias untuk function, biar jika parameternya panjang tidak perlu menulis panjang-panjang di sayHelloWithFilter

func sayHelloWithFilter(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

func filterName(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Syafiq", filterName)

	filter := filterName
	sayHelloWithFilter("Anjing", filter)
}
