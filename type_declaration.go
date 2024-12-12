package main

import "fmt"

func main() {
	// alias
	type NoKTP string
	var ktpSyafiq NoKTP = "11111111111"

	var contoh string = "22222222222"
	var contohKtp NoKTP = NoKTP(contoh) // conversion

	fmt.Println(ktpSyafiq)
	fmt.Println(contohKtp)
}
