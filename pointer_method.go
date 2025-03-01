package main

import "fmt"

type Man struct {
	Name string
}

// method sangat diwajibkan menggunakan pointer
func (man *Man) Merried() {
	man.Name = "Mr. " + man.Name
}

func main() {
	syafiq := Man{"Syafiq"}
	syafiq.Merried()

	fmt.Println(syafiq.Name)
}
