package main

import "fmt"

type HashName interface {
	GetName() string
}

func sayHi(value HashName) {
	fmt.Println("Hi", value.GetName())
}

// implementasi interface
type Personal struct {
	Name string
}

func (person Personal) GetName() string {
	return person.Name
}

// implementasi interface
type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	person := Personal{"Syafiq"}
	sayHi(person)

	animal := Animal{Name: "Kucing"}
	sayHi(animal)
}
