package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Syafiq"
	names[1] = "Fajrian"
	names[2] = "Emha"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		90,
		80,
		95,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(values))
	values[0] = 100
	fmt.Println(values)

	var i = [...]int{1, 2, 3, 4}

	fmt.Println(i)
}
