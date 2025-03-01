package main

import "fmt"

func main() {
	name := "Budi"

	switch name {
	case "Syafiq":
		fmt.Println("Hallo, Syafiq!")
	case "Budi":
		fmt.Println("Hallo, Budi!")
	default:
		fmt.Println("Hai, Boleh Kenalan?")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
