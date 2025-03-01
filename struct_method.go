package main

import "fmt"

type CustomerTo struct {
	Name, Address string
	Age           int
}

func (customer CustomerTo) sayHello(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func main() {
	var syafiq CustomerTo

	fmt.Println(syafiq)

	syafiq.Name = "Syafiq Fajrian Emha"
	syafiq.Address = "Purbalingga"
	syafiq.Age = 21

	fmt.Println(syafiq)
	fmt.Println(syafiq.Name)
	fmt.Println(syafiq.Address)
	fmt.Println(syafiq.Age)

	joko := CustomerTo{
		Name:    "Joko",
		Address: "Indonesia",
		Age:     35,
	}
	fmt.Println(joko)

	budi := CustomerTo{"Budi", "Indonesia", 30}
	fmt.Println(budi)

	budi.sayHello("Agus")
}
