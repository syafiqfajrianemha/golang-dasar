package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var syafiq Customer

	fmt.Println(syafiq)

	syafiq.Name = "Syafiq Fajrian Emha"
	syafiq.Address = "Purbalingga"
	syafiq.Age = 21

	fmt.Println(syafiq)
	fmt.Println(syafiq.Name)
	fmt.Println(syafiq.Address)
	fmt.Println(syafiq.Age)

	joko := Customer{
		Name:    "Joko",
		Address: "Indonesia",
		Age:     35,
	}
	fmt.Println(joko)

	budi := Customer{"Budi", "Indonesia", 30}
	fmt.Println(budi)
}
