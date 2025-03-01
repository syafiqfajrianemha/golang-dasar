package main

import "fmt"

func main() {
	name := "Syafiq"

	if name == "Syafiq" {
		fmt.Println("Hallo, Syafiq!")
	} else if name == "Budi" {
		fmt.Println("Hallo, Budi!")
	} else if name == "Joko" {
		fmt.Println("Hallo, Joko!")
	} else {
		fmt.Println("Hai, Boleh Kenalan?")
	}

	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
