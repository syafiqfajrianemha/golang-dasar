package main

import "fmt"

type AddressTo struct {
	City, Province, Country string
}

func main() {
	var address1 AddressTo = AddressTo{"Purbalingga", "Central Java", "Indonesia"}
	var address2 *AddressTo = &address1 // pointer
	address2.City = "Purwokerto"
	fmt.Println(address1) // ikut berubah
	fmt.Println(address2) // berubah menjadi Purwokerto

	//address2 = &AddressTo{"Jakarta", "DKI Jakarta", "Indonesia"}

	// asterisk operator atau * (bintang)
	*address2 = AddressTo{"Jakarta", "DKI Jakarta", "Indonesia"}

	fmt.Println(address1) // address 1 berubah
	fmt.Println(address2)
}
