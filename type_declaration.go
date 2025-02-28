package main

import "fmt"

func main() {
	// alias atatu membuat tipe data lain dari tipe data yang sudah ada
	type NoKtp string

	var ktpSyafiq NoKtp = "111111"

	var contoh string = "222222"
	var contohKtp NoKtp = NoKtp(contoh)

	fmt.Println(ktpSyafiq)
	fmt.Println(contohKtp)
}
