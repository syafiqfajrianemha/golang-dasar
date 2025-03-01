package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Purbalingga", "Central Java", "Indonesia"}
	// address2 := address1 // copy value

	// address2.City = "Purwokerto"

	// fmt.Println(address1) // tidak berubah
	// fmt.Println(address2) // berubah menjadi Purwokerto

	// pointer (pass by reference)
	address2 := &address1
	address2.City = "Purwokerto"

	fmt.Println(address1)
	fmt.Println(address2)
}
