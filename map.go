package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Syafiq",
		"address": "Purbalingga",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "Buku Golang"
	book["author"] = "Syafiq"
	book["wrong"] = "Ups"

	fmt.Println(book)

	delete(book, "wrong")

	fmt.Println(book)
}
