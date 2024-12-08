package main

import "fmt"

func main() {
	fmt.Println("Syafiq")
	fmt.Println("Syafiq Fajrian")
	fmt.Println("Syafiq Fajrian Emha")

	fmt.Println(len("Syafiq"))
	fmt.Println("Syafiq Fajrian"[0]) // hasilnya 83 karena S dalam byte (uint) adalah 83, jadi harus conversion ke string agar hasilnya S
	fmt.Println("Syafiq Fajrian Emha"[1])
}
